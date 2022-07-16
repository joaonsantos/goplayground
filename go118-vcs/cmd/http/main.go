package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	bi, _ := debug.ReadBuildInfo()
	fmt.Println(bi)
}
