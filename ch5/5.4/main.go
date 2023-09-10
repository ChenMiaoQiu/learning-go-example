// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visitScript(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(n *html.Node) string {
	str := ""
	if n == nil {
		return str
	}

	if n.Type == html.TextNode {
		str += n.Data
	}

	if n.Data != "style" && n.Data != "script" {
		str += visit(n.FirstChild)
	}

	str += visit(n.NextSibling)

	return str
}

func visitScript(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "img") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}

	links = visitScript(links, n.FirstChild)
	links = visitScript(links, n.NextSibling)

	return links
}
