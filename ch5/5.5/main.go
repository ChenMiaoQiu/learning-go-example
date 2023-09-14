package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	fmt.Println(resp)
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	words, images = 0, 0

	if n.Type == html.TextNode {
		str := strings.Split(n.Data, " ")
		words += len(str)
	} else if n.Data == "img" {
		images = 1
	}

	if n.Data != "style" && n.Data != "script" {
		x, y := countWordsAndImages(n.FirstChild)
		words += x
		images += y
	}

	x, y := countWordsAndImages(n.FirstChild)
	words += x
	images += y

	return
}

func main() {
	var url string
	fmt.Scan(&url)

	fmt.Println(CountWordsAndImages(url))
}
