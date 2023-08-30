package main

import (
	"fmt"
	"unicode"
)

func uniqueSpace(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i])) {
			j := i
			for j < len(s) && unicode.IsSpace(rune(s[j])) {
				j++
			}
			s = append(s[:i], s[j-1:]...)
			i = j - 1
		}
	}
	return s
}

func main() {
	s := "afdigneigu  \n\t   fjidfiefh  jsdifj frewf      "
	str := []byte(s)

	str = uniqueSpace(str)

	fmt.Println(string(str))
}
