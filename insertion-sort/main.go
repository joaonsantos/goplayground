package main

import (
  "fmt"
)

func insertionSort(unsorted []int) []int{
  sorted := make([]int, len(unsorted))
  copy(sorted, unsorted)

  for i := range sorted {
    for j := 0; j < i; j++ {
      if sorted[j] > sorted[i] {
        sorted[i], sorted[j]  = sorted[j], sorted[i]
      }
    }
  }

  return sorted
}

func main() {
  unsorted := []int{10,3,5,8,9,7,6,4,2,1}
  fmt.Println("Sorted ->", unsorted)

  sorted := insertionSort(unsorted)
  fmt.Println("Sorted ->", sorted)
}
