package main

import (
	"fmt"
	"io"
	"os"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var str []byte
	fmt.Scan(&str)

	cnt, err := w.Write(str)
	if err != nil {
		return w, nil
	}
	Cnt := int64(cnt)
	return w, &Cnt
}

func main() {
	input, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	_, num := CountingWriter(input)

	fmt.Println(*num)
}
