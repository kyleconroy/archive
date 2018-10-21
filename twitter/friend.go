package twitter

import (
	"encoding/json"
)

type Friend struct {
	ID int64 `json:"accountId,string"`
}

type friendObj struct {
	Obj friend `json:"following"`
}

type friend Friend

func (f *Friend) UnmarshalJSON(b []byte) error {
	var wrapper friendObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*f = Friend(wrapper.Obj)
	return nil
}
