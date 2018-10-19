package instagram

import (
	"encoding/json"
	"errors"
	"time"
)

type Likes struct {
	Media   []Like `json:"media_likes"`
	Comment []Like `json:"comment_likes"`
}

type Like struct {
	Author    string
	Timestamp time.Time
}

func (l Like) UnmarshalJSON(b []byte) error {
	var a []string
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}
	if len(a) != 2 {
		return errors.New("expected like to be a slice of two strings")
	}

	l.Timestamp, err = time.Parse(Stamp, a[0])
	if err != nil {
		return err
	}
	l.Author = a[1]
	return nil
}
