package cyoa

import (
	"encoding/json"
	"io"
	"log"
)

// Store the story in a map
type Story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title"`
	Paragraps []string `json:"story"`
	Options   []Option `json:"options"`
}

type Option struct {
	Text    string
	Chapter string
}

// Decode the file and return a Story object
func JsonStory(r io.Reader) (Story, error) {
	var story Story
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&story); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return story, nil

}
