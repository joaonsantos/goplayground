package main

import (
	"fmt"
	"strconv"
)


func main() {
  c := make(chan string)

  for i := 0; i < 10; i++ {
    go func(i int, c chan string){
      goroutineName := "goroutine: " + strconv.Itoa(i)
      c <- goroutineName
    }(i, c)
  }

  for i := 0; i < 10; i++ {
    fmt.Println(<-c)
  }

  fmt.Println("done")


}
