package instagram

import (
	"encoding/json"
	"time"
)

type Media struct {
	Stories []MediaItem `json:"stories"`
	Direct  []MediaItem `json:"direct"`
	Videos  []MediaItem `json:"videos"`
	Photos  []MediaItem `json:"photos"`
}

type MediaItem struct {
	Caption string    `json:"caption"`
	TakenAt time.Time `json:"taken_at"`
	Path    string    `json:"path"`
}

func (p MediaItem) UnmarshalJSON(b []byte) error {
	var m map[string]string
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	p.TakenAt, err = time.Parse(Stamp, m["taken_at"])
	if err != nil {
		// Try two formats
		p.TakenAt, err = time.Parse(time.RFC3339, m["taken_at"])
		if err != nil {
			return err
		}
	}
	p.Caption = m["caption"]
	p.Path = m["path"]
	return nil
}
