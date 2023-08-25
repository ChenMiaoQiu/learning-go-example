package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var res bytes.Buffer

	if strings.HasPrefix(s, "-") {
		res.WriteString("-")
		s = s[1:]
	}

	var last string
	if strings.Contains(s, ".") {
		last = s[strings.IndexAny(s, "."):]
		s = s[:strings.IndexAny(s, ".")]
	}

	cnt := len(s) % 3
	for i := 0; i < cnt; i++ {
		res.WriteByte(s[i])
	}

	for i, j := cnt, 0; i < len(s); i, j = i+1, (j+1)%3 {
		if cnt != 0 && j == 0 {
			res.WriteString(",")
		}
		res.WriteByte(s[i])
	}

	res.WriteString(last)

	return res.String()
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(comma(s))
}
