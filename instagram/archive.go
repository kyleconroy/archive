package instagram

import (
	"encoding/json"

	"github.com/kyleconroy/archive"
)

const Stamp = "2006-01-02T15:04:05"

type Archive struct {
	Comments      Comments
	Connections   Connections
	Contacts      []string
	Conversations []Thread
	Likes         Likes
	Media         Media
	Profile       Profile
	Searches      []Search
	Settings      Settings

	archive.Dir
}

func ParsePath(path string) (*Archive, error) {
	a := &Archive{Dir: archive.Dir(path)}

	files := map[string]interface{}{
		"comments.json":    &a.Comments,
		"connections.json": &a.Connections,
		"contacts.json":    &a.Contacts,
		"likes.json":       &a.Likes,
		"media.json":       &a.Media,
		"messages.json":    &a.Conversations,
		"profile.json":     &a.Profile,
		"searches.json":    &a.Searches,
		"settings.json":    &a.Settings,
	}

	for filename, p := range files {
		commentFile, err := a.Open(filename)
		if err != nil {
			return nil, err
		}
		if err := json.NewDecoder(commentFile).Decode(p); err != nil {
			return nil, err
		}
	}

	return a, nil
}
