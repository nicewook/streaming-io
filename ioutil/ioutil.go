package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	fn := "./planets.txt"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("fail to ReadFile: %v", err)
	}
	fmt.Println(string(b))

}
