package flipp

type Keyword string

type CampaignRequestBodyUser struct {

	// Key to keep track of user for retargeting purposes.
	// Required: true
	Key *string `json:"key"`
}

// Location location
//
// swagger:model Location
type Location struct {

	// accuracy radius
	AccuracyRadius int64 `json:"accuracy_radius,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// ip
	IP string `json:"ip,omitempty"`

	// metro code
	MetroCode string `json:"metro_code,omitempty"`

	// postal code
	PostalCode string `json:"postal_code,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// region name
	RegionName string `json:"region_name,omitempty"`
}

// Properties properties
//
// swagger:model Properties
type Properties struct {

	// content code
	ContentCode *string `json:"contentCode,omitempty"`

	// loc
	Loc Location `json:"loc"`
}

type PrebidRequest struct {

	// creative type
	// Required: true
	// Pattern: ^(DTX|NativeX)$
	CreativeType *string `json:"creativeType"`

	// height
	// Required: true
	Height *int64 `json:"height"`

	// publisher name identifier
	// Required: true
	PublisherNameIdentifier *string `json:"publisherNameIdentifier"`

	// request Id
	// Required: true
	RequestID *string `json:"requestId"`

	// width
	// Required: true
	Width *int64 `json:"width"`
}

type Placement struct {

	// ad types
	// Required: true
	// Min Items: 1
	AdTypes []int64 `json:"adTypes"`

	// count
	// Required: true
	Count *int64 `json:"count"`

	// div name
	DivName string `json:"divName,omitempty"`

	// network Id
	NetworkID int64 `json:"networkId,omitempty"`

	// prebid
	Prebid *PrebidRequest `json:"prebid,omitempty"`

	// properties
	Properties *Properties `json:"properties,omitempty"`

	// site Id
	// Required: true
	SiteID *int64 `json:"siteId"`

	// zone ids
	ZoneIds []int64 `json:"zoneIds"`
}

type CampaignRequestBody struct {

	// Accuracy Radius data from Maxmind for more precise targeting
	AccuracyRadius int64 `json:"accuracy_radius,omitempty"`

	// User IP that will be sent as part of the kevel request
	IP string `json:"ip,omitempty"`

	// An array of Keywords
	Keywords []Keyword `json:"keywords"`

	// An array of Placements
	// Required: true
	// Min Items: 1
	Placements []*Placement `json:"placements"`

	// Filters flyers based on preferred language (ISO 639-1 2char code).,
	// Pattern: ^(en|es|fr)$
	PreferredLanguage *string `json:"preferred_language,omitempty"`

	// The URL for the requesting page
	URL string `json:"url,omitempty"`

	// user
	// Required: true
	User *CampaignRequestBodyUser `json:"user"`
}

type CampaignResponseBody struct {

	// candidate retrieval
	CandidateRetrieval interface{} `json:"candidateRetrieval,omitempty"`

	// decisions
	// Required: true
	Decisions *Decisions `json:"decisions"`

	// location
	Location Location `json:"location"`
}

type Decisions struct {

	// inline
	Inline Inline `json:"inline,omitempty"`
}

type Inline []*InlineModel

type Contents []*Content

type Content struct {

	// body
	Body string `json:"body,omitempty"`

	// custom template
	CustomTemplate string `json:"customTemplate,omitempty"`

	// data
	Data *Data2 `json:"data,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

type Data2 struct {

	// JSON as a string of custom data returned from Kevel
	CustomData interface{} `json:"customData,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// width
	Width int64 `json:"width,omitempty"`
}

type InlineModel struct {

	// ad Id
	AdID int64 `json:"adId,omitempty"`

	// advertiser Id
	AdvertiserID int64 `json:"advertiserId,omitempty"`

	// campaign Id
	CampaignID int64 `json:"campaignId,omitempty"`

	// click Url
	ClickURL string `json:"clickUrl,omitempty"`

	// contents
	Contents Contents `json:"contents,omitempty"`

	// creative Id
	CreativeID int64 `json:"creativeId,omitempty"`

	// flight Id
	FlightID int64 `json:"flightId,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// impression Url
	ImpressionURL string `json:"impressionUrl,omitempty"`

	// prebid
	Prebid *PrebidResponse `json:"prebid,omitempty"`

	// priority Id
	PriorityID int64 `json:"priorityId,omitempty"`

	// storefront
	Storefront Storefront `json:"storefront,omitempty"`

	// width
	Width int64 `json:"width,omitempty"`
}

type PrebidResponse struct {

	// cpm
	// Required: true
	Cpm *float64 `json:"cpm"`

	// creative
	// Required: true
	Creative *string `json:"creative"`

	// creative type
	// Required: true
	CreativeType *string `json:"creativeType"`

	// request Id
	// Required: true
	RequestID *string `json:"requestId"`
}

type CampaignConfig struct {

	// fallback image Url
	FallbackImageURL string `json:"fallbackImageUrl,omitempty"`

	// fallback link Url
	FallbackLinkURL string `json:"fallbackLinkUrl,omitempty"`

	// tags
	Tags CampaignConfigTags `json:"tags,omitempty"`
}

type CampaignConfigTags struct {

	// advertiser
	Advertiser CampaignConfigTagsAdvertiser `json:"advertiser"`
}

type CampaignConfigTagsAdvertiser struct {

	// engagement
	Engagement string `json:"engagement,omitempty"`

	// impression
	Impression string `json:"impression,omitempty"`

	// open
	Open string `json:"open,omitempty"`
}

type Storefront struct {

	// campaign config
	CampaignConfig CampaignConfig `json:"campaignConfig"`

	// Flyer ID
	FlyerID int64 `json:"flyer_id,omitempty"`

	// Flyer Run ID
	FlyerRunID int64 `json:"flyer_run_id,omitempty"`

	// Flyer Type ID
	FlyerTypeID int64 `json:"flyer_type_id,omitempty"`

	// is fallback
	IsFallback *bool `json:"is_fallback,omitempty"`

	// Merchant Name
	Merchant string `json:"merchant,omitempty"`

	// Merchant ID
	MerchantID int64 `json:"merchant_id,omitempty"`

	// Flyer Name
	Name string `json:"name,omitempty"`

	// Merchant Logo for Storefont
	StorefrontLogoURL string `json:"storefront_logo_url,omitempty"`

	// storefront payload url
	StorefrontPayloadURL string `json:"storefront_payload_url,omitempty"`

	// Flyer Valid From
	ValidFrom string `json:"valid_from,omitempty"`

	// Flyer Valid To
	ValidTo string `json:"valid_to,omitempty"`
}
