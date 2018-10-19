package instagram

import "time"

type Search struct {
	Click string    `json:"search_click"`
	Time  time.Time `json:"time"`
	Type  string    `json:"type"`
}
