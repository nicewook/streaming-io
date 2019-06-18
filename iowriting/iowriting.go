package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fn := "./magic_msg.txt"
	f, err := os.Create(fn)
	if err != nil {
		log.Fatalf("fail to create: %v", err)
	}

	defer f.Close()
	if _, err := io.WriteString(f, "Go is fun!\n"); err != nil {
		log.Fatalf("fail to WriteString to file: %v", err)
	}

	f, err = os.Open(fn)
	if _, err = io.Copy(os.Stderr, f); err != nil {
		log.Fatalf("fail to WriteString to STDERR")
	}
}
