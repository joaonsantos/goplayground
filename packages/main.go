package main

import "fmt"
import "github.com/joaonsantos/goscripts-learning/packages/pkg/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)

	min := math.Min(xs)
	fmt.Println(min)

	max := math.Max(xs)
	fmt.Println(max)
}
