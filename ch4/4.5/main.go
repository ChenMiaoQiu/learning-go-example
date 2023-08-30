package main

import "fmt"

func main() {
	var a []string
	var n int

	fmt.Scan(&n)

	for n > 0 {
		var s string
		fmt.Scan(&s)
		a = append(a, s)
		n--
	}

	for i := 0; i < len(a); i++ {
		j := i + 1
		for j < len(a) && a[j] == a[i] {
			j++
		}
		a = append(a[0:i+1], a[j:]...)
		i = j - 1
	}

	fmt.Println(a)
}
