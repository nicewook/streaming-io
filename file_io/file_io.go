package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	fn := "./proverbs.txt"
	f, err := os.Create(fn)
	if err != nil {
		log.Fatalf("fail to create file: %v", err)
	}
	defer f.Close()

	for _, p := range proverbs {
		n, err := f.Write([]byte(p))
		if err != nil {
			log.Fatalf("fail to write: %v", err)
		}
		if n != len(p) {
			log.Fatalf("fail to write all: %v", err)
		}
	}

	b := make([]byte, 1024)
	f, err = os.Open(fn)
	_, err = f.Read(b)

	if err != nil && err != io.EOF {
		log.Fatalf("fail to read: %v", err)
	}

	fmt.Printf("%v", string(b))
}
