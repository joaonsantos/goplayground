package main

import (
	"iter"
)

type Ints struct{}

func (n *Ints) All() iter.Seq[int] {
	return func(yield func(int) bool) {
		i := 0
		for {
			if !yield(i) {
				return
			}
			i++
		}
	}
}

func limit[V any](max int, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		count := 0
		for v := range seq {
			if count > max {
				return
			}
			if !yield(v) {
				return
			}
			count++
		}
	}
}

func main() {
	seq := new(Ints).All()
	max := 10_000_000
	limitSeq := limit(max, seq)
	for range limitSeq {
		//fmt.Printf("value %d\n", n)
	}
}
