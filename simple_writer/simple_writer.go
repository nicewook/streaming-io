package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {

	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}

	var writer bytes.Buffer

	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))
		if err != nil {
			log.Fatalf("could not Write: %v", err)
		}
		if n != len(p) {
			log.Fatalf("fail to write data: %v", err)
		}
	}
	fmt.Println(writer.String())
}
