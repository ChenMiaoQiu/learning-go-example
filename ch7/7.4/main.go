package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func main() {
	_, err := html.Parse(NewReader("<dir>hello world</dir>"))

	if err != nil {
		fmt.Println(err)
		return
	}
}

type StringRead string

func (s *StringRead) Read(p []byte) (int, error) {
	copy(p, *s)
	return len(*s), io.EOF
}

func NewReader(s string) io.Reader {
	str := StringRead(s)
	return &str
}
