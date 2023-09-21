package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type countWord int
type countLine int

func (cnt *countWord) Write(s []byte) (n int, err error) {
	str := bufio.NewScanner(bytes.NewReader(s))
	str.Split(bufio.ScanWords)

	for str.Scan() {
		*cnt++
	}
	return int(*cnt), nil
}

func (cnt *countLine) Write(s []byte) (n int, err error) {
	str := bufio.NewScanner(bytes.NewReader(s))
	str.Split(bufio.ScanLines)

	for str.Scan() {
		*cnt++
	}
	return int(*cnt), nil
}

func main() {
	var a countWord
	var b countLine

	fmt.Println(a.Write([]byte(`hello world
	you
	thanks`)))
	fmt.Println(b.Write([]byte(`hello world
	you
	thanks`)))
}
