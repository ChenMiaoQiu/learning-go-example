package main

import "fmt"

func expand(s string, f func(string) string) string {
	return f(s)
}

func main() {
	fmt.Println(expand("666", func(a string) string {
		return a + "7777"
	}))
}
