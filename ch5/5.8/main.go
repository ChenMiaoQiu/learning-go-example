// go run ..\fetch\fetch.go https://www.gopl.io/ | go run main.go
// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var singleElement = map[string]bool{
	"link": true,
	"img":  true,
	"meta": true,
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	forEachNode(doc, startElement, endElement)
	c := ElementByID(doc, "img")
	startElement(c)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if !pre(n) {
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil && !singleElement[n.Data] {
		if !post(n) {
			return
		}
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	if doc.Type == html.ElementNode && doc.Data == id {
		return doc
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		child := ElementByID(c, id)
		if child != nil {
			return child
		}
	}

	return nil
}

var depth int

func startElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)

		for _, attr := range n.Attr {
			fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
		}

		if singleElement[n.Data] {
			fmt.Printf("/>\n")
		} else {
			fmt.Printf(">\n")
			depth++
		}
	}

	return true
}

func endElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return true
}
