package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
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
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

type HandlerOption func(h *handler)

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	handler := handler{
		s: s,
		t: tpl,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&handler)
		}

	}
	return handler
}

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithCustomPath(pathFunc func(request *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = pathFunc
	}
}

var tpl = template.Must(template.New("chapter").Parse(defaultHandlerTpl))

func defaultPathFunc(r *http.Request) string {
	log.Printf("The path received is %s", r.URL.Path)
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	return path
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var path string
	var storyTemplate *template.Template
	if h.pathFn == nil {
		path = defaultPathFunc(r)
	} else {
		path = h.pathFn(r)
	}
	if chapter, ok := h.s[path]; ok {
		log.Printf("found chapter inside dict for path %s", path)
		if h.t == nil {
			storyTemplate = tpl
		} else {
			storyTemplate = h.t
		}
		err := storyTemplate.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
}

var defaultHandlerTpl = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Choose Your Own Adventure</title>
    <style>
        * { box-sizing: border-box; margin: 0; padding: 0; }

        body {
            background-color: #1a1a2e;
            color: #e0d7c6;
            font-family: 'Georgia', serif;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 2rem;
        }

        .chapter {
            background-color: #16213e;
            border: 1px solid #2e3a5a;
            border-radius: 8px;
            max-width: 720px;
            width: 100%;
            padding: 2.5rem 3rem;
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.5);
        }

        h1 {
            font-size: 1.8rem;
            color: #c9a84c;
            border-bottom: 1px solid #2e3a5a;
            padding-bottom: 0.75rem;
            margin-bottom: 1.5rem;
            letter-spacing: 0.03em;
        }

        p {
            line-height: 1.8;
            margin-bottom: 1rem;
            font-size: 1.05rem;
            color: #ccc5b5;
        }

        .options {
            list-style: none;
            margin-top: 2rem;
            display: flex;
            flex-direction: column;
            gap: 0.75rem;
        }

        .options li a {
            display: block;
            padding: 0.75rem 1.25rem;
            background-color: #0f3460;
            color: #e0d7c6;
            text-decoration: none;
            border-radius: 5px;
            border-left: 3px solid #c9a84c;
            transition: background-color 0.2s, border-color 0.2s;
            font-size: 0.98rem;
        }

        .options li a:hover {
            background-color: #1a4a80;
            border-left-color: #e8c46a;
            color: #fff;
        }

        .options li a::before {
            content: "→ ";
            color: #c9a84c;
        }
    </style>
</head>
<body>
    <div class="chapter">
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}
        {{if .Options}}
        <ul class="options">
            {{range .Options}}
            <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
            {{end}}
        </ul>
        {{end}}
    </div>
</body>
</html>
`
