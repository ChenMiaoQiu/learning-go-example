package main

import "fmt"

func reverse(s *[5]int) {
	l, r := 0, 4

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	reverse(&a)

	fmt.Println(a)
}
