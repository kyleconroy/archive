package twitter

import (
	"encoding/json"
	"time"
)

type EmailAddressChange struct {
	AccountID int64
	From      string
	To        string
	At        time.Time
}

type emailChange struct {
	AccountID int64     `json:"accountID,string"`
	Change    emailDiff `json:"emailChange"`
}

type emailDiff struct {
	From string    `json:"changedFrom"`
	To   string    `json:"changedTo"`
	At   time.Time `json:"changedAt"`
}

type emailChangeObj struct {
	Obj emailChange `json:"emailAddressChange"`
}

func (ec *EmailAddressChange) UnmarshalJSON(b []byte) error {
	var obj emailChangeObj
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}
	ec.AccountID = obj.Obj.AccountID
	ec.From = obj.Obj.Change.From
	ec.To = obj.Obj.Change.To
	ec.At = obj.Obj.Change.At
	return nil
}
