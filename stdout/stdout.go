package main

import (
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

	for _, p := range proverbs {
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			log.Fatalf("fail to write to Stdout")
		}
		if n != len(p) {
			log.Fatalf("fail to write data to Stdout")
		}
	}
}
