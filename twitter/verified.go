package twitter

import (
	"encoding/json"
)

type VerifiedAccount struct {
	ID       int64 `json:"accountId,string"`
	Verified bool  `json:"verified"`
}

type verifiedObj struct {
	Obj verified `json:"verified"`
}

type verified VerifiedAccount

func (v *VerifiedAccount) UnmarshalJSON(b []byte) error {
	var wrapper verifiedObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*v = VerifiedAccount(wrapper.Obj)
	return nil
}
