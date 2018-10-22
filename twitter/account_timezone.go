package twitter

import "encoding/json"

// TODO: Parse the Timezone string into a time.Location
type AccountTimezone struct {
	AccountID int64  `json:"accountId,string"`
	Timezone  string `json:"timeZone"`
}

type accountTimezoneObj struct {
	Obj accountTimezone `json:"accountTimezone"`
}

type accountTimezone AccountTimezone

func (at *AccountTimezone) UnmarshalJSON(b []byte) error {
	var wrapper accountTimezoneObj
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*at = AccountTimezone(wrapper.Obj)
	return nil
}
