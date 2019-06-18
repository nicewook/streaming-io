package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {

	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		if err != nil {
			// log.Fatalf == log.Prinf + os.Exit(1)
			log.Fatalf("could not read: %v", err)
		}
		fmt.Println(string(p[:n]))
	}
}
