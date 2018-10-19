package twitter

import (
	"encoding/json"
)

type BlockedAccount struct {
	ID int64 `json:"accountId,string"`
}

type blockingObj struct {
	Obj BlockedAccount `json:"blocking"`
}

func (ba *BlockedAccount) UnmarshalJSON(b []byte) error {
	var wrapper blockingObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*ba = wrapper.Obj
	return nil
}
