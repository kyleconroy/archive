package twitter

import "encoding/json"

type Follower struct {
	ID int64 `json:"accountId,string"`
}

type followerObj struct {
	Obj follower `json:"follower"`
}

type follower Follower

func (f *Follower) UnmarshalJSON(b []byte) error {
	var wrapper followerObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*f = Follower(wrapper.Obj)
	return nil
}
