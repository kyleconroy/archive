package twitter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/kyleconroy/archive"
)

type Archive struct {
	AccountCreationIPs  []AccountCreationIP  `json:"account_creation_ip"`
	AccountTimezones    []AccountTimezone    `json:"account_timezone"`
	Accounts            []Account            `json:"account"`
	AgeInfos            []AgeInfo            `json:"ageinfo"`
	BlockedAccounts     []BlockedAccount     `json:"block"`
	EmailAddressChanges []EmailAddressChange `json:"email_address_change"`
	Followers           []Follower           `json:"follower"`
	Friends             []Friend             `json:"following"`
	IPAudits            []IPAudit            `json:"ip_audit"`
	Likes               []Like               `json:"like"`
	ScreenNameChanges   []ScreenNameChange   `json:"screen_name_change"`
	Tweets              []Tweet              `json:"tweet"`
	CreatedLists        []List               `json:"lists_created"`
	SubscribedLists     []List               `json:"lists_subscribed"`
	JoinedLists         []List               `json:"lists_member"`

	// Empty
	AccountSuspensions              []AccountSuspension    `json:"account_suspension"`
	AdEngagements                   []Ad                   `json:"ad_engagements"`
	AdImpressions                   []Ad                   `json:"ad_impressions"`
	AttributedMobileAdConversions   []AdConversion         `json:"ad_mobile_conversions_attributed"`
	AttributedOnlineAdConversions   []AdConversion         `json:"ad_online_conversions_attributed"`
	ConnectedApplications           []ConnectedApplication `json:"connected_application"`
	Contacts                        []Contact              `json:"contact"`
	FacebookConnections             []FacebookConnection   `json:"facebook_connection"`
	UnattributedMobileAdConversions []AdConversion         `json:"ad_mobile_conversions_unattributed"`
	UnattributedOnlineAdConversions []AdConversion         `json:"ad_online_conversions_unattributed"`

	archive.Dir
}

type archiveEntry struct {
	filename string
	part     string
}

func ParsePath(path string) (*Archive, error) {
	a := &Archive{Dir: archive.Dir(path)}

	paths := []archiveEntry{
		// {"contact.js", "contact", "contacts"},
		// {"direct-message-headers.js", "direct_message_headers", ""},
		// {"direct-message.js", "direct_message", "direct_messages"},
		// {"email-address-change.js", "email_address_change", "email_address_changes"},
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
		{"connected-application.js", "connected_application"},
		{"contact.js", "contact"},
		{"block.js", "block"},
		{"email-address-change.js", "email_address_change"},
		{"facebook-connection.js", "facebook_connection"},
		{"follower.js", "follower"},
		{"following.js", "following"},
		{"ip-audit.js", "ip_audit"},
		{"like.js", "like"},
		{"lists-created.js", "lists_created"},
		{"lists-member.js", "lists_member"},
		{"lists-subscribed.js", "lists_subscribed"},
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
