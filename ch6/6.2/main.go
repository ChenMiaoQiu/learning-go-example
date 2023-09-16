package main

import (
	"fmt"

	"github.com/ChenMiaoQiu/learning-go-example/ch6/6.2/bit"
)

func main() {
	var x bit.IntSet
	x.AddAll(1, 2, 3, 4, 5)
	fmt.Println(x.String())
}
