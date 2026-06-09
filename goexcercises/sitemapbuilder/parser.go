package sitemapbuilder

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Link string
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return text
}

func Parse(r io.Reader) ([]Link, error) {
	var links []Link
	node, err := html.Parse(r)
	if err != nil {
		return links, err
	}
	for _, node := range linkNodes(node) {
		var href string
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				href = attr.Val
				break
			}
		}
		{
			links = append(links, Link{
				Href: href,
				Link: strings.TrimSpace(extractText(node)),
			})
		}

	}
	return links, nil
}

func linkNodes(node *html.Node) []*html.Node {
	var aNodes []*html.Node
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		childLinkNodes := linkNodes(c)
		aNodes = append(aNodes, childLinkNodes...)
	}
	return aNodes
}
