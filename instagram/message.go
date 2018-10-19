package instagram

import "time"

type Thread struct {
	Conversations []Message `json:"conversation"`
	Participants  []string  `json:"participants"`
}

type Message struct {
	CreatedAt         time.Time     `json:"created_at"`
	MediaOwner        string        `json:"media_owner"`
	MediaShareCaption string        `json:"media_share_caption"`
	Sender            string        `json:"sender"`
	StoryShare        string        `json:"story_share"`
	Text              string        `json:"text"`
	Heart             string        `json:"heart"`
	Likes             []MessageLike `json:"likes"`
}

type MessageLike struct {
	Username string    `json:"username"`
	Date     time.Time `json:"date"`
}
