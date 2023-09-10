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
	fmt.Println(visit(doc))
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
