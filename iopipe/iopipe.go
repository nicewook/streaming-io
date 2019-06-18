package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	pReader, pWriter := io.Pipe()

	go func() {
		defer pWriter.Close()
		io.Copy(pWriter, proverbs)
	}()

	io.Copy(os.Stdout, pReader)
	pReader.Close()
}
