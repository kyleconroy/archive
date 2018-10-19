package instagram

import (
	"encoding/json"
	"time"
)

type Profile struct {
	DateJoined     time.Time `json:"date_joined"`
	Email          string    `json:"email"`
	Gender         string    `json"gender"`
	PrivateAccount bool      `json:"private_account"`
	Name           string    `json:"name"`
	PhoneNumber    string    `json:"phone_number"`
	ProfilePicURL  string    `json:"profile_pic_url"`
	Username       string    `json:"username"`
}

type profile struct {
	DateJoined     string `json:"date_joined"`
	Email          string `json:"email"`
	Gender         string `json"gender"`
	PrivateAccount bool   `json:"private_account"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	ProfilePicURL  string `json:"profile_pic_url"`
	Username       string `json:"username"`
}

func (p Profile) UnmarshalJSON(b []byte) error {
	var pp profile
	err := json.Unmarshal(b, &pp)
	if err != nil {
		return err
	}
	p.DateJoined, err = time.Parse(Stamp, pp.DateJoined)
	if err != nil {
		return err
	}

	p.Email = pp.Email
	p.Gender = pp.Gender
	p.PrivateAccount = pp.PrivateAccount
	p.Name = pp.Name
	p.PhoneNumber = pp.PhoneNumber
	p.ProfilePicURL = pp.ProfilePicURL
	p.Username = pp.Username
	return nil
}
