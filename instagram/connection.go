package instagram

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type Connection struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

type Connections struct {
	BlockedUsers       []Connection `json:"blocked_users"`
	FollowRequestsSent []Connection `json:"follow_requests_sent"`
	Followers          []Connection `json:"followers"`
	Following          []Connection `json:"following"`
	FollowingHashtags  []Connection `json:"following_hashtags"`
}

type connections struct {
	BlockedUsers       map[string]string `json:"blocked_users"`
	FollowRequestsSent map[string]string `json:"follow_requests_sent"`
	Followers          map[string]string `json:"followers"`
	Following          map[string]string `json:"following"`
	FollowingHashtags  map[string]string `json:"following_hashtags"`
}

func parseConnections(m map[string]string) ([]Connection, error) {
	c := make([]Connection, len(m))
	i := 0
	for k, v := range m {
		ts, err := time.Parse(Stamp, v)
		if err != nil {
			return c, err
		}
		c[i] = Connection{Name: k, Timestamp: ts}
		i += 1
	}
	sort.Slice(c, func(i, j int) bool { return c[j].Timestamp.Before(c[i].Timestamp) })
	return c, nil
}

func (c *Connections) UnmarshalJSON(b []byte) error {
	var cc connections
	err := json.Unmarshal(b, &cc)
	if err != nil {
		return err
	}

	c.BlockedUsers, err = parseConnections(cc.BlockedUsers)
	if err != nil {
		return fmt.Errorf("parsing blocked users failed :%s", err)
	}

	c.FollowRequestsSent, err = parseConnections(cc.FollowRequestsSent)
	if err != nil {
		return fmt.Errorf("parsing follow requests failed :%s", err)
	}

	c.Followers, err = parseConnections(cc.Followers)
	if err != nil {
		return fmt.Errorf("parsing followers failed :%s", err)
	}

	c.Following, err = parseConnections(cc.Following)
	if err != nil {
		return fmt.Errorf("parsing following failed :%s", err)
	}

	c.FollowingHashtags, err = parseConnections(cc.FollowingHashtags)
	if err != nil {
		return fmt.Errorf("parsing following hashtags failed :%s", err)
	}

	return nil
}
