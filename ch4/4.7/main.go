package main

import "fmt"

func reverse(s []byte) {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {
	a := "123456789"
	s := []byte(a)

	reverse(s)

	fmt.Println(string(s))
}
