package main

import "fmt"

func fibonacci(iteration byte) byte {

  if iteration == 1 || iteration == 0 {
    return iteration
  } else {
    return fibonacci(iteration - 1) +
    fibonacci(iteration - 2)
  }
}


func main() {
  fmt.Printf(
    "This program computes the first 10 " +
    "iterations of the fibonacci sequence " +
    "in a recursive manner.\n")

    var iter byte = 0
    for iter < 10 {
      fmt.Println(fibonacci(iter))
      iter++
    }
}
