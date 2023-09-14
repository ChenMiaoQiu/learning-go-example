package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ChenMiaoQiu/learning-go-example/ch5/links"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	create, _ := os.Create("save.txt")

	breadthFirst(crawl, os.Args[1:], create)
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(string, io.Writer) []string, worklist []string, out io.Writer) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, out)...)
			}
		}
	}
}

func crawl(url string, out io.Writer) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	for _, s := range list {
		if strings.HasPrefix(s, url) {
			fmt.Fprintln(out, s)
		}
	}
	return list
}
