package twitter

import (
	"encoding/json"
)

type MutedAccount struct {
	ID int64 `json:"accountId,string"`
}

type mutingObj struct {
	Obj mutedAccount `json:"muting"`
}

type mutedAccount MutedAccount

func (m *MutedAccount) UnmarshalJSON(b []byte) error {
	var wrapper mutingObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*m = MutedAccount(wrapper.Obj)
	return nil
}
