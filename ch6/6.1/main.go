package main

import (
	"fmt"

	"github.com/ChenMiaoQiu/learning-go-example/ch6/6.1/bit"
)

func main() {
	var x, y bit.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(x.Len())
	x.Remove(9)
	fmt.Println(x.String())

	y = *x.Copy()
	x.Clear()
	fmt.Println(x.String())
	fmt.Println(y.String())
}
