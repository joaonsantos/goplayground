package main

import (
	"fmt"
	"os"
	"strconv"
)

func makeOddGenenerator() func() uint{
  i := uint(1)

  return func() uint {
    defer func(){i += 2}()
    return i
  }
}

func main() {
  args := os.Args
  if len(args) < 2 {
    fmt.Println("usage: go run main.go <count>\n" +
    "where count is the number of odd numbers to be generated")
    return
  }

  nextOdd := makeOddGenenerator()
  oddCount, err := strconv.Atoi(args[1])
  if err != nil {
    panic(err)
  }

  i := 0
  for i < oddCount {
    fmt.Println(nextOdd())
    i++
  }

}
