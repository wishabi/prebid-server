package flipp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/gofrs/uuid"
	"github.com/prebid/openrtb/v17/openrtb2"
	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/config"
	"github.com/prebid/prebid-server/errortypes"
	"github.com/prebid/prebid-server/openrtb_ext"
)

var BANNER_TYPE = "banner"
var INLINE_DIV_NAME = "inline"
var AD_TYPES = []int64{4309, 641}
var DTX_TYPES = []int64{5061}
var COUNT = int64(1)
var FLIPP_BIDDER = "flipp"
var DEFAULT_CURRENCY = "USD"

type adapter struct {
	endpoint string
}

// Builder builds a new instance of the Foo adapter for the given bidder with the given config.
func Builder(bidderName openrtb_ext.BidderName, config config.Adapter, server config.Server) (adapters.Bidder, error) {
	bidder := &adapter{
		endpoint: config.Endpoint,
	}
	return bidder, nil
}

func (a *adapter) MakeRequests(request *openrtb2.BidRequest, reqInfo *adapters.ExtraRequestInfo) ([]*adapters.RequestData, []error) {
	adapterRequests := make([]*adapters.RequestData, 0, len(request.Imp))

	for _, imp := range request.Imp {
		var flippExtParams openrtb_ext.ImpExtFlipp
		params, _, _, err := jsonparser.Get(imp.Ext, "prebid", "bidder", FLIPP_BIDDER)
		if err != nil {
			return nil, []error{err}
		}
		err = json.Unmarshal(params, &flippExtParams)
		if err != nil {
			return nil, []error{err}
		}
		publisherUrl, err := url.Parse(request.Site.Page)
		if err != nil {
			return nil, []error{err}
		}
		prebidRequest := PrebidRequest{
			CreativeType:            &flippExtParams.CreativeType,
			PublisherNameIdentifier: &flippExtParams.PublisherNameIdentifier,
			RequestID:               &request.ID,
			Height:                  &imp.Banner.Format[0].H,
			Width:                   &imp.Banner.Format[0].W,
		}
		contentCode := publisherUrl.Query().Get("flipp-content-code")
		placement := Placement{
			DivName: INLINE_DIV_NAME,
			SiteID:  &flippExtParams.SiteID,
			AdTypes: getAdTypes(flippExtParams.CreativeType),
			ZoneIds: flippExtParams.ZoneIds,
			Count:   &COUNT,
			Prebid:  &prebidRequest,
			Properties: &Properties{
				ContentCode: &contentCode,
			},
		}

		var userKey string
		if request.User != nil && request.User.ID != "" {
			userKey = request.User.ID
		} else if flippExtParams.UserKey != "" {
			userKey = flippExtParams.UserKey
		} else {
			uid, err := uuid.NewV4()
			if err != nil {
				return nil, []error{err}
			}
			userKey = uid.String()
		}

		keywordsArray := strings.Split(request.Site.Keywords, ",")
		var keywordsSlice []Keyword

		for _, k := range keywordsArray {
			keywordsSlice = append(keywordsSlice, Keyword(k))
		}

		campaignRequestBody := CampaignRequestBody{
			Placements: []*Placement{&placement},
			URL:        request.Site.Page,
			Keywords:   keywordsSlice,
			IP:         request.Device.IP,
			User: &CampaignRequestBodyUser{
				Key: &userKey,
			},
		}

		adapterReq, err := a.makeRequest(request, campaignRequestBody)
		if err != nil {
			return nil, []error{err}
		}

		if adapterReq != nil {
			adapterRequests = append(adapterRequests, adapterReq)
		}
	}
	return adapterRequests, nil
}

func (a *adapter) makeRequest(request *openrtb2.BidRequest, campaignRequestBody CampaignRequestBody) (*adapters.RequestData, error) {
	campaignRequestBodyJSON, err := json.Marshal(campaignRequestBody)
	if err != nil {
		return nil, err
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	headers.Add("User-Agent", request.Device.UA)
	headers.Add("X-Forwarded-For", request.Device.IP)
	return &adapters.RequestData{
		Method:  "POST",
		Uri:     a.endpoint,
		Body:    campaignRequestBodyJSON,
		Headers: headers,
	}, err
}

func (a *adapter) MakeBids(request *openrtb2.BidRequest, requestData *adapters.RequestData, responseData *adapters.ResponseData) (*adapters.BidderResponse, []error) {
	if responseData.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if responseData.StatusCode == http.StatusBadRequest {
		err := &errortypes.BadInput{
			Message: "Unexpected status code: 400. Bad request from publisher. Run with request.debug = 1 for more info.",
		}
		return nil, []error{err}
	}

	if responseData.StatusCode != http.StatusOK {
		err := &errortypes.BadServerResponse{
			Message: fmt.Sprintf("Unexpected status code: %d. Run with request.debug = 1 for more info.", responseData.StatusCode),
		}
		return nil, []error{err}
	}

	var campaignResponseBody CampaignResponseBody
	if err := json.Unmarshal(responseData.Body, &campaignResponseBody); err != nil {
		return nil, []error{err}
	}

	bidResponse := adapters.NewBidderResponseWithBidsCapacity(len(request.Imp))
	bidResponse.Currency = DEFAULT_CURRENCY
	for _, decision := range campaignResponseBody.Decisions.Inline {
		b := &adapters.TypedBid{
			Bid: &openrtb2.Bid{
				CrID:  fmt.Sprint(decision.CreativeID),
				Price: *decision.Prebid.Cpm,
				AdM:   *decision.Prebid.Creative,
				ID:    fmt.Sprint(decision.AdID),
				W:     decision.Contents[0].Data.Width,
				H:     decision.Contents[0].Data.Height,
			},
			BidType: openrtb_ext.BidType(BANNER_TYPE),
		}
		bidResponse.Bids = append(bidResponse.Bids, b)
	}
	return bidResponse, nil
}

func getAdTypes(creativeType string) []int64 {
	if creativeType == "DTX" {
		return DTX_TYPES
	}
	return AD_TYPES
}
