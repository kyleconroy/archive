package twitter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/kyleconroy/archive"
)

type Archive struct {
	Accounts                        []Account           `json:"account"`
	AccountCreationIPs              []AccountCreationIP `json:"account_creation_ip"`
	AccountSuspensions              []AccountSuspension `json:"account_suspension"`
	AccountTimezones                []AccountTimezone   `json:"account_timezone"`
	AdEngagements                   []Ad                `json:"ad_engagements"`
	AdImpressions                   []Ad                `json:"ad_impressions"`
	AgeInfos                        []AgeInfo           `json:"ageinfo"`
	BlockedAccounts                 []BlockedAccount    `json:"block"`
	ScreenNameChanges               []ScreenNameChange  `json:"screen_name_change"`
	Tweets                          []Tweet             `json:"tweet"`
	Followers                       []Follower          `json:"follower"`
	Friends                         []Friend            `json:"following"`
	Likes                           []Like              `json:"like"`
	AttributedMobileAdConversions   []AdConversion      `json:"ad_mobile_conversions_attributed"`
	UnattributedMobileAdConversions []AdConversion      `json:"ad_mobile_conversions_unattributed"`
	AttributedOnlineAdConversions   []AdConversion      `json:"ad_online_conversions_attributed"`
	UnattributedOnlineAdConversions []AdConversion      `json:"ad_online_conversions_unattributed"`

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
		{"account-suspension.js", "account_suspension"},
		{"account-timezone.js", "account_timezone"},
		{"account.js", "account"},
		{"ad-engagements.js", "ad_engagements"},
		{"ad-impressions.js", "ad_impressions"},
		{"ad-mobile-conversions-attributed.js", "ad_mobile_conversions_attributed"},
		{"ad-mobile-conversions-unattributed.js", "ad_mobile_conversions_unattributed"},
		{"ad-online-conversions-attributed.js", "ad_online_conversions_attributed"},
		{"ad-online-conversions-unattributed.js", "ad_online_conversions_unattributed"},
		{"ageinfo.js", "ageinfo"},
		{"block.js", "block"},
		{"like.js", "like"},
		// {"connected-application.js", "connected_application", "connected_applications"},
		// {"contact.js", "contact", "contacts"},
		// {"direct-message-headers.js", "direct_message_headers", ""},
		// {"direct-message.js", "direct_message", "direct_messages"},
		// {"email-address-change.js", "email_address_change", "email_address_changes"},
		{"follower.js", "follower"},
		{"following.js", "following"},
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
