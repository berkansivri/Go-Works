package main

import (
	"io"
	"fmt"
	"os"
)

type fileWriter struct{}

func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fw := fileWriter{}
	io.Copy(fw, file)
}