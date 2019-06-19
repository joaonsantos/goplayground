package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func compute_recursive(iteration uint64) uint64 {
  if iteration == 1 || iteration == 0 {
    return iteration
  } else {
    return compute_recursive(iteration - 1) +
    compute_recursive(iteration - 2)
  }
}

func fibonacci_recursive(iteration uint64) {
    var iter uint64 = 0
    for iter < 10 {
      fmt.Println(compute_recursive(iter))
      iter++
    }
}

func fibonacci_iterative(iteration uint64) {
  var first, second uint64 = 0, 1
  var iter uint64 = 1


  for iter <= iteration {
    if iter == 1 {
      fmt.Println(0)
    } else if iter == 2 {
      fmt.Println(1)
    } else {
      result := first + second
      fmt.Println(result)

      // Update first and second element
      first, second = second, result
    }
      iter++
  }
}

func fibonacci(iteration uint64, mode string) {
  if mode == "iterative" {
    fibonacci_iterative(iteration)
  } else {
    fibonacci_recursive(iteration)
  }
}

func main() {
  // Create new parser object
	parser := argparse.NewParser(
    "Fibonnaci Sucession Computation",
    "This program computes the given number of " +
    "iterations of the fibonacci sequence")

    // Create a flag to control number of iterations
    iterations := parser.Int(
      "i",
      "iterations",
      &argparse.Options{
        Required: true,
        Help: "Number of iterations to compute",
        Default: 10 })

    // Flag to choose computation method
    mode := parser.String(
      "m",
      "mode",
      &argparse.Options{
        Required: true,
        Help: "Choose weather to compute recursively or iteratively." +
        "Eg. \"recursive\" or \"iterative\"",
        Default: "iterative" })

    // Parse input
    err := parser.Parse(os.Args)
    if err != nil {
      // In case of error print error and print usage
      // This can also be done by passing -h or --help flags
      fmt.Print(parser.Usage(err))
    }

    // Run the computation, print to stdout
    fibonacci(uint64(*iterations), *mode)
}

