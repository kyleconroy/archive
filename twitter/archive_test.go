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
		"block":               {golden.BlockedAccounts, a.BlockedAccounts},
		"screen-name-change":  {golden.ScreenNameChanges, a.ScreenNameChanges},
		"tweet":               {golden.Tweets, a.Tweets},
		"friend":              {golden.Friends, a.Friends},
		"follower":            {golden.Followers, a.Followers},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(test.a, test.b, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("%s differ: (-want +got)\n%s", name, diff)
			}
		})
	}
}
