package twitter

import "encoding/json"

type Profile struct {
	AvatarMediaURL string      `json:"avatarMediaUrl"`
	Description    Description `json:"description"`
	HeaderMediaURL string      `json:"headerMediaUrl"`
}

type Description struct {
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Website  string `json:"website"`
}

type profileObj struct {
	Obj profile `json:"profile"`
}

type profile Profile

func (p *Profile) UnmarshalJSON(b []byte) error {
	var wrapper profileObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*p = Profile(wrapper.Obj)
	return nil
}
