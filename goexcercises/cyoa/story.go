package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

type Story map[string]Chapter

func JsonStory(r io.Reader) (Story, error) {
	var story Story
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type handler struct {
	s Story
}

var tpl = template.Must(template.New("chapter").Parse(defaultHandlerTpl))

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		http.Error(w, "Could not execute template code", http.StatusInternalServerError)
		return
	}
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

var defaultHandlerTpl = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Choose your own adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
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
