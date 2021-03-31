package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func main() {

  type task struct {
    id int
    name string
  }

  // Create task queue.
  tasks := make(chan task)
  workers := 2

  // Create tasks.
  go func() {
    for i := 0; i < 10; i++ {
      fmt.Println("iteration:", i)
      tasks <- task{i, fmt.Sprintf("task %d", i)}
    }
    close(tasks)
  }()

  // Create a sync group with workers and schedule results close when all tasks are finished
  var wg sync.WaitGroup
  wg.Add(workers)
  results := make(chan string)
  go func() {
    wg.Wait()
    fmt.Println("tasks finished")
    close(results)
  }()

  // Consume tasks in using workers
  for i := 0; i < workers; i++ {
    go func() {
      defer wg.Done()

      for t := range tasks {
        seconds := rand.Intn(10)
        fmt.Println("seconds", seconds)
        time.Sleep(time.Second * time.Duration(seconds))
        results <- t.name
      }
    }()
  }

  for r := range results {
    fmt.Println("received:", r)
  }
  fmt.Println("consumed all results")

}
