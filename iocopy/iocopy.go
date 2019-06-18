package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	f, e := os.Create("./proverbs.txt")
	if e != nil {
		log.Fatalf("fail to create: %v", e)
	}
	defer f.Close()

	if _, e := io.Copy(f, proverbs); e != nil {
		log.Fatalf("fail to copy: %v", e)
	}
	fmt.Println("file created")

	f, e = os.Open("./proverbs.txt")
	if e != nil {
		log.Fatalf("fail to open: %v", e)
	}

	if _, e := io.Copy(os.Stdout, f); e != nil {
		log.Fatalf("fail to io.Copy: %v", e)
	}
}
