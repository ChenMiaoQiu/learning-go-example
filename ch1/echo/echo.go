package main

import (
	"fmt"
	"os"
	"strings"
)

// go run .\echo.go 666
func main() {
	echo1()
	echo2()
	echo3()
}

func echo1() {
	s, sep := "", ""

	//Args[0]: 使用地址, 后面则为使用时输入的字符串，以空格隔开
	for _, arg := range os.Args[1:] {
		s += sep + arg
		// 实现间隔, 如go run .\echo1.go 666 777 则输出666 777
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo3() {
	fmt.Println(os.Args[1:])
}
