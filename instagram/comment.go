package instagram

import (
	"encoding/json"
	"errors"
	"time"
)

type Comments struct {
	Live  []Comment `json:"live_comments"`
	Media []Comment `json:"media_comments"`
}

type Comment struct {
	Author    string
	Message   string
	Timestamp time.Time
}

func (c Comment) UnmarshalJSON(b []byte) error {
	var a []string
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}
	if len(a) != 3 {
		return errors.New("expected comment to be a slice of three strings")
	}
	c.Timestamp, err = time.Parse(Stamp, a[0])
	if err != nil {
		return err
	}
	c.Message = a[1]
	c.Author = a[2]
	return nil
}
