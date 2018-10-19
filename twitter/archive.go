package twitter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/kyleconroy/archive"
)

type Archive struct {
	Accounts           []Account           `json:"account"`
	AccountCreationIPs []AccountCreationIP `json:"account_creation_ip"`
	AdEngagements      []Ad                `json:"ad_engagements"`
	AdImpressions      []Ad                `json:"ad_impressions"`
	BlockedAccounts    []BlockedAccount    `json:"block"`
	ScreenNameChanges  []ScreenNameChange  `json:"screen_name_change"`
	Tweets             []Tweet             `json:"tweet"`

	archive.Dir
}

type archiveEntry struct {
	filename string
	part     string
}

func ParsePath(path string) (*Archive, error) {
	a := &Archive{Dir: archive.Dir(path)}

	paths := []archiveEntry{
		{"account-creation-ip.js", "account_creation_ip"},
		{"account.js", "account"},
		{"ad-engagements.js", "ad_engagements"},
		{"ad-impressions.js", "ad_impressions"},
		// {"ageinfo.js", "ageinfo", "ageinfos"},
		{"block.js", "block"},
		// {"connected-application.js", "connected_application", "connected_applications"},
		// {"contact.js", "contact", "contacts"},
		// {"direct-message-headers.js", "direct_message_headers", ""},
		// {"direct-message.js", "direct_message", "direct_messages"},
		// {"email-address-change.js", "email_address_change", "email_address_changes"},
		// {"follower.js", "follower", "followers"},
		// {"following.js", "following", ""},
		{"screen-name-change.js", "screen_name_change"},
		{"tweet.js", "tweet"},
	}

	for _, e := range paths {
		p, err := a.Open(e.filename)
		if err != nil {
			return nil, err
		}

		input, err := ioutil.ReadAll(p)
		if err != nil {
			return nil, err
		}

		// TODO: Make this a regular expression
		output := bytes.Replace(input,
			[]byte("window.YTD."+e.part+".part0 = "),
			[]byte("{\""+e.part+"\": "),
			-1)
		output = append(output, "}"...)

		if err := json.Unmarshal(output, &a); err != nil {
			return nil, err
		}
	}
	return a, nil
}
