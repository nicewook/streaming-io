package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fn := "./planets.txt"

	f, err := os.Open(fn)
	if err != nil {
		log.Fatalf("fail to create: %v", err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("fail to read: %v", err)
		}
		fmt.Print(line)
	}

}
