package main

import "fmt"

func fibonacci(iteration byte) {
  var first, second byte = 0, 1
  var iter byte = 1


  for iter <= iteration {
    if iter == 1 {
      fmt.Println(0)
    } else if iter == 2 {
      fmt.Println(1)
    } else {
      result := first + second
      fmt.Println(result)

      // update first and second element
      first, second = second, result
    }
      iter++
  }


}

func main() {
  fmt.Printf(
    "This program computes the first 10 " +
    "iterations of the fibonacci sequence " +
    "in an iterative manner.\n")

    fibonacci(10)

}

