package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var res bytes.Buffer

	cnt := len(s) % 3
	for i := 0; i < cnt; i++ {
		res.WriteByte(s[i])
	}

	for i, j := cnt, 0; i < len(s); i, j = i+1, (j+1)%3 {
		if res.Len() != 0 && j == 0 {
			res.WriteString(",")
		}
		res.WriteByte(s[i])
	}

	return res.String()
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(comma(s))
}
