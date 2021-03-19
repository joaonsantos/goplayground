package main

import (
	"fmt"
	"sort"
)

func main(){
  a := []int{6, 3, 2, 5, 4, 1}
  fmt.Println("before order:", a)

  // one-liner to reserve memory and initialize a slice copy
  o := append([]int{}, a...)

  // convert []int to IntSlice as it implements the sort.Sort interface and sort
  sort.Sort(sort.IntSlice(o))
  fmt.Println("after order:", o)

  // additionally convert IntSlice to sort.Reverse as it redefines Less, finally sort
  sort.Sort(sort.Reverse(sort.IntSlice(o)))
  fmt.Println("after order reverse:", o)
}

