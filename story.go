package cyoa

import (
	"encoding/json"
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
