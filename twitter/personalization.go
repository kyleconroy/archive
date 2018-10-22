package twitter

import "encoding/json"

type Personalization struct {
	Demographics    Demographics      `json:"demographics"`
	Interests       Interests         `json:"interests"`
	LocationHistory []LocationHistory `json:"locationHistory"`
}

type Demographics struct {
	GenderInfo GenderInfo `json:"genderInfo"`
	Languages  []Language `json:"languages"`
}

type GenderInfo struct {
	Gender string `json:"gender"`
}

type LocationHistory struct {
}

type Language struct {
	Language string `json:"language"`
	Disabled bool   `json:"isDisabled"`
}

type Interests struct {
	Interests              []string `json:"interests"`
	PartnerInterests       []string `json:"partnerInterests"`
	AudienceAndAdvertisers Audience `json:"audienceAndAdvertisers"`
	Shows                  []string `json:"shows"`
}

type Audience struct {
	Size        int      `json:"numAudiences,string"`
	Advertisers []string `json:"advertisers"`
}

type personalizationWrapper struct {
	Obj personalization `json:"p13nData"`
}

type personalization Personalization

func (p *Personalization) UnmarshalJSON(b []byte) error {
	var wrapper personalizationWrapper
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*p = Personalization(wrapper.Obj)
	return nil
}
