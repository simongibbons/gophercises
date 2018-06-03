package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) (l []Link, err error) {
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	for _, node := range linkNodes(root) {
		l = append(l, linkFromNode(node))
	}

	return l, nil
}

func linkFromNode(n *html.Node) (l Link) {
	return Link{
		Href: getHref(n.Attr),
		Text: strings.Join(getText(n), ""),
	}
}

func getHref(attrs []html.Attribute) string {
	for _, attr := range attrs {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func getText(node *html.Node) (rv []string) {
	if node.Type == html.TextNode {
		return []string{node.Data}
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		rv = append(rv, getText(n)...)
	}
	return rv
}

func linkNodes(root *html.Node) (nodes []*html.Node) {
	if root.Type == html.ElementNode && root.DataAtom == atom.A {
		return []*html.Node{root}
	}

	for n := root.FirstChild; n != nil; n = n.NextSibling {
		nodes = append(nodes, linkNodes(n)...)
	}

	return nodes
}