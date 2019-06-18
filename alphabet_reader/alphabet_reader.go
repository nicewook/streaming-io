package main

import (
	"fmt"
	"io"
	"log"
)

type abcReader struct {
	s string // source string
	i int    // index
}

func newABCReader(s string) *abcReader {
	return &abcReader{s: s} // i == 0 as default
}

func abcFilter(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}
func (a *abcReader) Read(p []byte) (int, error) {
	if a.i >= len(a.s) {
		return 0, io.EOF // read all up
	}

	x := len(a.s) - a.i
	n, bound := 0, 0
	if x >= len(p) {
		bound = len(p)
	} else {
		bound = x
	}
	buf := make([]byte, bound)
	for n < bound {
		if char := abcFilter(a.s[a.i]); char != 0 {
			buf[n] = char
			n++
		}
		a.i++
		if a.i >= len(a.s) {
			n--
			break
		}
	}
	copy(p, buf)
	return n, nil
}
func main() {
	reader := newABCReader("Hello! It's 9am, where is the sun?")
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
