package main

import (
	"fmt"

	"github.com/ChenMiaoQiu/learning-go-example/ch6/6.3/bit"
)

func main() {
	var x, y bit.IntSet
	x.AddAll(1, 2, 3, 4, 5)
	y.AddAll(6, 7, 8, 9, 10, 1, 2, 3)
	fmt.Println(x.String())
	fmt.Println(y.String())
	x.IntersectWith(&y)
	fmt.Println(x.String())

	x.Clear()
	y.Clear()
	x.AddAll(1, 2, 3, 4, 5)
	y.AddAll(6, 7, 8, 9, 10, 1, 2, 3)
	x.DifferenceWith(&y)
	fmt.Println(x.String())

	x.Clear()
	y.Clear()
	x.AddAll(1, 2, 3, 4, 5)
	y.AddAll(6, 7, 8, 9, 10, 1, 2, 3)
	z := x.SymmetricDifference(&y)
	fmt.Println(z.String())

}
