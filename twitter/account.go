package twitter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Account struct {
	CreatedAt   time.Time
	CreatedVia  string `json:"createdVia"`
	DisplayName string `json:"accountDisplayName"`
	Email       string `json:"email"`
	ID          int64  `json:"accountId,string"`
	TimeZone    string `json:"timeZone"`
	UpdatedAt   time.Time
	Username    string `json:"username"`
}

type accountAlias Account

type accountJSON struct {
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	accountAlias
}

type accountObj struct {
	Account accountJSON `json:"account"`
}

// Twitter will return timestamps in two formats, so attempt to parse both
func parseTimestamp(ts string) (time.Time, error) {
	t, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		tt, err := time.Parse(time.RFC3339, ts)
		if err != nil {
			return time.Time{}, err
		}
		return tt, nil
	}
	return time.Unix(t, 0), nil
}

func (a *Account) UnmarshalJSON(b []byte) error {
	var ao accountObj
	err := json.Unmarshal(b, &ao)
	if err != nil {
		return err
	}

	aj := ao.Account
	*a = Account(aj.accountAlias)

	if aj.CreatedAt != "" {
		a.CreatedAt, err = parseTimestamp(aj.CreatedAt)
		if err != nil {
			return fmt.Errorf("createdAt: %s", err)
		}
	}

	if aj.UpdatedAt != "" {
		a.UpdatedAt, err = parseTimestamp(aj.UpdatedAt)
		if err != nil {
			return fmt.Errorf("updatedAt: %s", err)
		}
	}

	return nil
}
