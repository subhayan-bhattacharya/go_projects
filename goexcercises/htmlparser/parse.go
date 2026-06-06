package htmlparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents an href tag inside an HTML document
type Link struct {
	Href string
	Text string
}

// Parse would take an HTML document in and would parse me the links from it
func Parse(r io.Reader) ([]Link, error) {
	var links []Link
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	for _, linkNode := range linkNodes(doc) {
		links = append(links, buildLinks(linkNode))
	}
	return links, nil
}

func buildLinks(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}
func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	var nodes []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childLinkNodes := linkNodes(c)
		nodes = append(nodes, childLinkNodes...)
	}
	return nodes
}
