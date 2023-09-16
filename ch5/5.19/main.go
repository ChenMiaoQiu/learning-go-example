package main

import "fmt"

func test() (res int) {
	defer func() {
		if p := recover(); p != nil {
			res = 1
		}
	}()
	panic("666")
}

func main() {
	fmt.Println(test())
}
