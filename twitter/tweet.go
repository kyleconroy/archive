package twitter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Tweet struct {
	CreatedAt           time.Time
	DisplayTestRange    []int
	Entities            Entity `json:"entities"`
	ExtendedEntities    Entity `json:"extended_entities"`
	FavoriteCount       int    `json:"favorite_count,string"`
	Favorited           bool   `json:"favorited"`
	FullText            string `json:"full_text"`
	ID                  int64  `json:"id_str,string"`
	InReplyToScreenName string `json:"in_reply_to_screen_name"`
	InReplyToStatus     int64  `json:"in_reply_to_status_id_str,string"`
	InReplyToUser       int64  `json:"in_reply_to_user_id_str,string"`
	Language            string `json:"lang"`
	PossiblySensitive   bool   `json:"possibly_sensitive"`
	RetweetCount        int    `json:"retweet_count,string"`
	Retweeted           bool   `json:"retweeted"`
	Source              string `json:"source"`
	Truncated           bool   `json:"truncated"`
}

type tweetAlias Tweet

type tweetJSON struct {
	CreatedAt        string   `json:"created_at"`
	DisplayTestRange []string `json:"display_text_range"`

	tweetAlias
}

type Entity struct {
	// UserMentions []UserMention `json:"user_mentions"`
	// "urls" : [ ],
	// "symbols" : [ ],
	Hashtags []Hashtag `json:"hashtags"`
	Media    []Media   `json:"media"`
}

type Media struct {
	DisplayURL    string `json:"display_url"`
	ExpandedURL   string `json:"expanded_url"`
	ID            int64  `json:"id_str,string"`
	Indices       []int
	MediaURL      string `json:"media_url"`
	MediaURLHTTPS string `json:"media_url_https"`
	Sizes         Sizes  `json:"sizes`
	Type          string `json:"type"`
	URL           string `json:"url"`
}

type mediaAlias Media

type mediaJSON struct {
	Indices []string `json:"indices"`

	mediaAlias
}

func (m *Media) UnmarshalJSON(b []byte) error {
	var mj mediaJSON
	err := json.Unmarshal(b, &mj)
	if err != nil {
		return err
	}
	*m = Media(mj.mediaAlias)
	m.Indices, err = parseIntSlice(mj.Indices)
	if err != nil {
		return fmt.Errorf("indices: %s", err)
	}
	return nil
}

type Sizes struct {
	Large  Size `json:"large"`
	Medium Size `json:"medium"`
	Small  Size `json:"small"`
	Thumb  Size `json:"thumb"`
}

type Size struct {
	Height int    `json:"h,string"`
	Resize string `json:"resize"`
	Width  int    `json:"w,string"`
}

type Hashtag struct {
	Indices []int
	Text    string `json:"text"`
}

type hashtagAlias Hashtag

type hashtagJSON struct {
	Indices []string `json:"indices"`

	hashtagAlias
}

func (m *Hashtag) UnmarshalJSON(b []byte) error {
	var hj hashtagJSON
	err := json.Unmarshal(b, &hj)
	if err != nil {
		return err
	}
	*m = Hashtag(hj.hashtagAlias)
	m.Indices, err = parseIntSlice(hj.Indices)
	if err != nil {
		return fmt.Errorf("indices: %s", err)
	}
	return nil
}

func parseInt64(v string) (int64, error) {
	if v == "" {
		return 0, nil
	}
	return strconv.ParseInt(v, 10, 64)
}

func parseIntSlice(vs []string) ([]int, error) {
	var err error
	o := make([]int, len(vs))
	for i, v := range vs {
		o[i], err = strconv.Atoi(v)
		if err != nil {
			return o, fmt.Errorf("error parsing int in array: %d: %s", i, err)
		}
	}
	return o, nil
}

func (t *Tweet) UnmarshalJSON(b []byte) error {
	var tj tweetJSON
	err := json.Unmarshal(b, &tj)
	if err != nil {
		return err
	}
	*t = Tweet(tj.tweetAlias)
	t.CreatedAt, err = time.Parse(time.RubyDate, tj.CreatedAt)
	if err != nil {
		return fmt.Errorf("created_at: %s", err)
	}
	t.DisplayTestRange, err = parseIntSlice(tj.DisplayTestRange)
	if err != nil {
		return fmt.Errorf("display_text_range: %s", err)
	}
	return nil
}
