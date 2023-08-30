// cat .\test.txt | go run .\main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	hash := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		word := in.Text()
		hash[word]++
	}

	for word, cnt := range hash {
		fmt.Println(word, cnt)
	}
}
