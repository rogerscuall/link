package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link reporesent a link (<a href="...">...</a>) in an HTML document.
type Link struct {
	Href string
	Text string
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			ret.Href = a.Val
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
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

var r io.Reader

// Parse takes a reader and returns a slice of links.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var links []Link
	nodes := linkNodes(doc)
	for _, n := range nodes {
		links = append(links, buildLink(n))
	}
	//dfs(doc, "")
	return links, nil

}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func dfs(n *html.Node, padding string) {
	fmt.Println(padding, n.Type ,n.Data)
	for c:=n.FirstChild; c!=nil; c=c.NextSibling {
		dfs(c, padding+"  ")
	}
}

