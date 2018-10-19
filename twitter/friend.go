package twitter

import (
	"encoding/json"
	"fmt"
)

type Friend struct {
	ID int64 `json:"accountId,string"`
}

type friendObj struct {
	Obj Follower `json:"following"`
}

func (f *Friend) UnmarshalJSON(b []byte) error {
	fmt.Println(string(b))
	var wrapper friendObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	fmt.Printf("%+v: %s\n", wrapper, err)
	f.ID = wrapper.Obj.ID
	return nil
}
