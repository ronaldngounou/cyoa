package cyoa

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

// Store the story in a map
type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

// Decode the file and return a Story object
func JsonStory(r io.Reader) (Story, error) {
	var story Story
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil

}

// easier to package it inside our source code
var defaultHandlerTemplate = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Choose Your Own adventure.</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		<p>{{range .Paragraphs}}</p>
			<p>{{.}}</p>
		{{end}}

		<ul>
			{{range .Options}}
				<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
			{{end}}
		</ul>
		
	</body>
	</html>
`

type handler struct {
	s Story
	t *template.Template
}

// Functional Option design pattern
type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

// It is better to return an interface rather than returning a handler type
// because when we generate the godoc -http:8080, it makes it easier to see what
// methods are under the interface.
// If we were to return the handler type, the methods under the interface
// won't get exported.
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

// ServehTTP can be called by a handler object
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving HTTP request...")
	path := strings.TrimSpace(r.URL.Path)

	// edge case: no path at all or root path
	if path == "" || path == "/" {
		path = "/intro"
	}
	// "/intro" => "intro"
	path = path[1:]
	fmt.Printf("--- the path is: %s \n", path)
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			// Don't expose development errors to the end users because it could contain something
			// we don't want them to see like our Database password.
			http.Error(w, "Something went wrong..", http.StatusInternalServerError)
		}
		// Found the chapter
		fmt.Println("found the chapter")
		fmt.Printf("the chapter is: %s", chapter)
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)

}
