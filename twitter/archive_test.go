package twitter

import (
	"encoding/gob"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestParsePath(t *testing.T) {
	a, err := ParsePath("testdata/archive")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open("testdata/golden.bin")
	if err != nil {
		t.Fatal(err)
	}

	var golden Archive
	dec := gob.NewDecoder(f)
	if err := dec.Decode(&golden); err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		a interface{}
		b interface{}
	}{
		"account":             {golden.Accounts, a.Accounts},
		"account-creation-ip": {golden.AccountCreationIPs, a.AccountCreationIPs},
		"account-timezone":    {golden.AccountTimezones, a.AccountTimezones},
		"ageinfo":             {golden.AgeInfos, a.AgeInfos},
		"block":               {golden.BlockedAccounts, a.BlockedAccounts},
		"email-change":        {golden.EmailAddressChanges, a.EmailAddressChanges},
		"follower":            {golden.Followers, a.Followers},
		"friend":              {golden.Friends, a.Friends},
		"ip-audit":            {golden.IPAudits, a.IPAudits},
		"like":                {golden.Likes, a.Likes},
		"list-created":        {golden.CreatedLists, a.CreatedLists},
		"list-members":        {golden.JoinedLists, a.JoinedLists},
		"list-subscribe":      {golden.SubscribedLists, a.SubscribedLists},
		"mute":                {golden.MutedAccounts, a.MutedAccounts},
		"personalization":     {golden.Personalizations, a.Personalizations},
		"profile":             {golden.Profiles, a.Profiles},
		"screen-name-change":  {golden.ScreenNameChanges, a.ScreenNameChanges},
		"tweet":               {golden.Tweets, a.Tweets},
		"verified":            {golden.Verified, a.Verified},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(test.a, test.b, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("%s differ: (-want +got)\n%s", name, diff)
			}
		})
	}
}
