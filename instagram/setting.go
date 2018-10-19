package instagram

type Settings struct {
	AllowCommentsFrom string   `json:"allow_comments_from"`
	BlockedCommenters []string `json:"blocked_commenters"`
	FilteredKeywords  []string `json:"filtered_keywords"`
}
