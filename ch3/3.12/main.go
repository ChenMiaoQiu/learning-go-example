package main

import (
	"fmt"
)

func isSame(a, b string) bool {
	hash := make(map[rune]int)

	for _, val := range a {
		hash[val]++
	}

	for _, val := range b {
		if hash[val] <= 0 {
			return false
		}
		hash[val]--
	}

	for _, val := range hash {
		if val > 0 {
			return false
		}
	}

	return true
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	if isSame(a, b) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
