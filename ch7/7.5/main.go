package main

import (
	"bytes"
	"fmt"
	"io"
)

type limitedReader struct {
	n int64
	r io.Reader
}

func (l *limitedReader) Read(p []byte) (int, error) {
	if l.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) >= l.n {
		p = p[:l.n]
	}

	n, err := l.r.Read(p)
	l.n -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitedReader{n, r}
}

func main() {
	data := []byte("hello, world")
	read := LimitReader(bytes.NewReader(data), 5)
	buffer := make([]byte, 10)
	n, err := read.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buffer[:n]))
}
