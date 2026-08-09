package htmlparser2

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func extractLinkText(link *html.Node) string {
	var text strings.Builder
	for child := range link.Descendants() {
		if child.Type == html.TextNode {
			text.WriteString(child.Data)
		}
	}
	return strings.Join(strings.Fields(text.String()), " ")
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	//fmt.Println("----")
	//prettyPrint(doc)
	//fmt.Println("----")
	links := linkNodes(doc)
	var extractedLinks []Link
	var hrefVal string
	for _, link := range links {
		linkText := extractLinkText(link)
		for _, attr := range link.Attr {
			if attr.Key == "href" {
				hrefVal = attr.Val
			}
		}
		extractedLinks = append(extractedLinks, Link{
			Href: hrefVal,
			Text: linkText,
		})
	}

	return extractedLinks, nil
}

func linkNodes(node *html.Node) []*html.Node {
	var linkNodes []*html.Node
	for child := range node.Descendants() {
		if child.Type == html.ElementNode && child.Data == "a" {
			linkNodes = append(linkNodes, child)
		}
	}
	return linkNodes
}

func prettyPrint(node *html.Node) {
	for n := range node.Descendants() {
		if n.Type == html.ElementNode {
			depth := 0
			for parent := n.Parent; parent != nil; parent = parent.Parent {
				if parent.Type == html.ElementNode {
					depth++
				}
			}
			msg := "<" + n.Data + ">"
			padding := strings.Repeat("  ", depth)
			fmt.Println(padding, msg)
		}
	}
}
