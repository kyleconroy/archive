package twitter

import (
	"encoding/json"
	"time"
)

type ScreenNameChange struct {
	AccountID int64
	From      string
	To        string
	At        time.Time
}

type screenNameChange struct {
	AccountID int64          `json:"accountID,string"`
	Change    screenNameDiff `json:"screenNameChange"`
}

type screenNameDiff struct {
	From string    `json:"changedFrom"`
	To   string    `json:"changedTo"`
	At   time.Time `json:"changedAt"`
}

type screenNameChangeObj struct {
	Obj screenNameChange `json:"screenNameChange"`
}

func (snc *ScreenNameChange) UnmarshalJSON(b []byte) error {
	var obj screenNameChangeObj
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}
	snc.AccountID = obj.Obj.AccountID
	snc.From = obj.Obj.Change.From
	snc.To = obj.Obj.Change.To
	snc.At = obj.Obj.Change.At
	return nil
}
