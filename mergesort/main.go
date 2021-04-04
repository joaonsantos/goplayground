package main

import (
	"fmt"

	"github.com/joaonsantos/goscripts-learning/mergesort/sort"
)

func main() {
	s := []int{6, 3, 2, 5, 4, 1, 10, 9, 8, 7}
	fmt.Println("before mergesort:", s)

	ss := sort.Mergesort(s)
	fmt.Println("after mergesort:", ss)
}
