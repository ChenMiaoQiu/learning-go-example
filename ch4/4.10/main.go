// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ChenMiaoQiu/learning-go-example/ch4/4.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var month, year, overYear []github.Issue
	for _, item := range result.Items {
		len := time.Since(item.CreatedAt).Hours()

		day := len / 24

		if day <= 30 {
			month = append(month, *item)
		} else if day <= 365 {
			year = append(year, *item)
		} else {
			overYear = append(overYear, *item)
		}
	}

	fmt.Println("Less than a month")
	for _, item := range month {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Less than a year")
	for _, item := range year {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Over a year")
	for _, item := range overYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
