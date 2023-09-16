package main

import (
	"fmt"

	"github.com/ChenMiaoQiu/learning-go-example/ch6/6.4/bit"
)

func main() {
	var x bit.IntSet
	x.AddAll(1, 2, 3, 4, 5, 6, 7, 8)

	for _, val := range x.Elems() {
		fmt.Println(val)
	}
}
