package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x string
	flag := 0

	for p, val := range os.Args[1:] {
		if p == 1 {
			x = val
		} else {
			flag, _ = strconv.Atoi(val)
		}
	}

	if flag == 0 {
		res := sha256.Sum256([]byte(x))
		fmt.Printf("%x", res)
	} else if flag == 1 {
		res := sha512.Sum384([]byte(x))
		fmt.Printf("%x", res)
	} else {
		res := sha512.Sum512([]byte(x))
		fmt.Printf("%x", res)
	}
}
