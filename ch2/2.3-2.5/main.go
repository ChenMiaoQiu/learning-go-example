package main

import (
	"fmt"
	"time"

	"github.com/ChenMiaoQiu/learning-go-example/ch2/2.3-2.5/popcount1"
	"github.com/ChenMiaoQiu/learning-go-example/ch2/2.3-2.5/popcount2"
	"github.com/ChenMiaoQiu/learning-go-example/ch2/2.3-2.5/popcount3"
	"github.com/ChenMiaoQiu/learning-go-example/ch2/2.3-2.5/popcount4"
)

func main() {
	var x uint64
	fmt.Scan(&x)
	start := time.Now()
	fmt.Println(popcount1.PopCount(x))
	fmt.Printf("第一个方法用时为:")
	fmt.Println(time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(popcount2.PopCount(x))
	fmt.Printf("第二个方法用时为:")
	fmt.Println(time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(popcount3.PopCount(x))
	fmt.Printf("第三个方法用时为:")
	fmt.Println(time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(popcount4.PopCount(x))
	fmt.Printf("第四个方法用时为:")
	fmt.Println(time.Since(start).Seconds())
}
