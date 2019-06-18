package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type abcReader struct {
	reader io.Reader
}

func newABCReader(reader io.Reader) *abcReader {
	return &abcReader{reader: reader}
}

func abcFilter(r byte) byte {
	if ('a' <= r && 'z' >= r) || ('A' <= r && 'Z' >= r) {
		return r
	}
	return 0
}

func (a *abcReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := abcFilter(p[i]); char != 0 {
			buf[i] = char
		}
	}
	copy(p, buf)
	return n, nil
}

func main() {

	f, err := os.Open("strings.go")
	if err != nil {
		log.Fatalf("could not open: %v", err)
	}
	defer f.Close()

	reader := newABCReader(f)
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read: %v", err)
		}
		fmt.Print(string(p[:n]))
	}
}
