package main

import (
	"fmt"
	"time"
  "sync"
)

// worker1 is a faster worker
func worker1(wg *sync.WaitGroup, out1 <-chan int) {
  defer wg.Done()
  for d := range out1 {
    time.Sleep(time.Second * 1)
    fmt.Println("process 1", d)
  }
}

// worker2 is a slower worker
func worker2(wg *sync.WaitGroup, out2 <-chan int) {
  defer wg.Done()
  for d := range out2 {
    time.Sleep(time.Second * 2)
    fmt.Println("process 2", d)
  }
}

// Fanout receives a channel and distributes to available channels
// since each channel is associated with a worker, effectively this selects the available worker
func Fanout(in chan int, out1, out2 chan int) {
  defer func(){
    close(out1)
    close(out2)
    fmt.Println("all input processed")
  }()

  for data := range in {
    select {
      case out1 <- data:
      case out2 <- data:
    }
  }
}

func main() {
  const tasks = 10
  const workers = 2

  in := make(chan int)
  out1 := make(chan int, 2)
  out2 := make(chan int, 2)

  var wg sync.WaitGroup
  wg.Add(workers)

  go worker1(&wg, out1)
  go worker2(&wg, out2)

  go func() {
    defer close(in)
    for i := 0; i < tasks; i++ {
      in <- i
    }
  }()

  Fanout(in, out1, out2)
  wg.Wait()
  fmt.Println("all tasks processed")
}
