package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	c1 := sha256.Sum256([]byte(a))
	c2 := sha256.Sum256([]byte(b))

	fmt.Printf("out1: %x\n", c1)
	fmt.Printf("out2: %x\n", c2)

	cnt := 0
	for i := 0; i < 32; i++ {
		a, b := c1[i], c2[i]

		for j := 0; j < 8; j++ {
			if (a >> j & 1) != (b >> j & 1) {
				cnt++
			}
		}
	}

	fmt.Println("diff:", cnt)
}
