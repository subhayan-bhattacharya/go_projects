package htmlparser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents an href tag inside an HTML document
type Link struct {
	Href string
	Text string
}

// Parse would take an HTML document in and would parse me the links from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := range n.ChildNodes() {
		dfs(c, padding+"  ")
	}
}
