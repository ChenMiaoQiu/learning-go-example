package main

import (
	"fmt"
	"sort"
	"strings"
)

func IsPalIsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	str := "abccba"

	s := strings.Split(str, "")
	fmt.Println(s, ":", IsPalIsPalindrome(sort.StringSlice(s)))
}
