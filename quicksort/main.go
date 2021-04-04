package main

import (
  "fmt"
)

// partition finds correct pivot position in slice s with bounds l(ow) and h(igh)
// pivot is always the last element of the slice
//
// the idea is to use a "wall" which contains all elements smaller than the pivot
// an item is swapped when it is smaller than the pivot because we want it to be behind the wall
// when an element is swapped, advance the wall to fit more space
func partition(s []int, l, h int) int {
  pv := s[h]
  wall := l

  for i := l + 1; i < h; i++ {
    if s[i] <= pv {
      s[wall], s[i] = s[i], s[wall]
      wall++
    }
  }

  // useful if the array turns out to be already sorted
  // the pivot should be of a higher value than the wall value before swapping
  if s[h] < s[wall] {
    s[wall], s[h] = s[h], s[wall]
  }

  return wall
}

// function quicksort sorts slice s with bounds l(ow) and h(igh)
func quicksort(s []int, l, h int){
  if l < h {
    p := partition(s, l, h)
    quicksort(s, l, p)
    quicksort(s, p + 1 , h)
  }
}

func main(){
  a := []int{6, 3, 2, 5, 4, 1, 10, 9, 8, 7}
  fmt.Println("before quicksort:", a)

  o := append([]int{}, a...)

  quicksort(o, 0 , len(o) - 1)
  fmt.Println("after quicksort:", o)
}
