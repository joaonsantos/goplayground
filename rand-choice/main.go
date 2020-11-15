package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runGenerator() {
  rand.Seed(time.Now().Unix())

  values := []string {
    "is rainy.",
    "is sunny.",
    "is foggy.",
    "is clear.",
  }

  for {
    fmt.Println("Today the weather", values[rand.Intn(len(values))])
    time.Sleep(1000 * time.Millisecond)
  }

}

func main() {
  runGenerator()
}
