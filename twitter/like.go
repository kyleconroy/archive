package twitter

import (
	"encoding/json"
)

type Like struct {
	TweetID int64 `json:"tweetId,string"`
}

type likeObj struct {
	Obj like `json:"like"`
}

type like Like

func (l *Like) UnmarshalJSON(b []byte) error {
	var wrapper likeObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*l = Like(wrapper.Obj)
	return nil
}
