package main

import "fmt"

func max(val ...int) (int, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("no num")
	}

	res := val[0]
	for _, val := range val {
		res = maxNum(res, val)
	}

	return res, nil
}

func min(val ...int) (int, error) {
	if len(val) == 0 {
		return 0, fmt.Errorf("no num")
	}

	res := val[0]
	for _, val := range val {
		res = minNum(res, val)
	}

	return res, nil
}

func minNum(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxNum(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	fmt.Println(max(1, 2, 3, 5))
	fmt.Println(max())
	fmt.Println(min(1, 2, 3, 5))
	fmt.Println(min())
}
