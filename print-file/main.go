package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]

	fmt.Println("Going to print out the contents of", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
