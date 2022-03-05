package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		panic("This program must receive exactly one argument with the path of the file to read")
	}

	f, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, f)

}
