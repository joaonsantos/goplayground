package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)

	for input.Scan() {
		counts[input.Text()] += 1
	}

	if err := input.Err(); err != nil {
		panic("error while scanning input")
	}

	for l, c := range counts {
		if c > 1 {
			fmt.Printf("%v\t%v\n", l, c)
		}
	}
}
