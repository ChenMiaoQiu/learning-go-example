package main

import "fmt"

func reverse(arr []int, k int) []int {
	len := len(arr)
	res := make([]int, len)
	p := 0
	k %= len

	for i := k; i < len; i++ {
		res[p] = arr[i]
		p++
	}

	for i := 0; i < k; i++ {
		res[p] = arr[i]
		p++
	}

	return res
}

func main() {
	var k int
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Scan(&k)

	arr = reverse(arr, k)

	fmt.Println(arr)
}
