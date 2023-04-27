package openrtb_ext

type ImpExtFlippOptions struct {
	StartCompact *bool `json:"startCompact,omitempty"`
}

type ImpExtFlipp struct {
	PublisherNameIdentifier string              `json:"publisherNameIdentifier"`
	CreativeType            string              `json:"creativeType"`
	SiteID                  int64               `json:"siteId"`
	ZoneIds                 []int64             `json:"zoneIds,omitempty"`
	UserKey                 string              `json:"userKey,omitempty"`
	Options                 *ImpExtFlippOptions `json:"options,omitempty"`
}
