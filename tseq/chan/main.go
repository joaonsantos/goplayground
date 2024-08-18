package main

func generator(qsig chan struct{}) <-chan int {
	out := make(chan int, 1)
	go func() {
		defer close(out)

		i := 0
		for {
			select {
			case <-qsig:
				return
			case out <- i:
				i++
			}
		}
	}()
	return out
}

func limit(max int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			if i > max {
				return
			}
			out <- i
		}
	}()
	return out
}

func main() {
	qsig := make(chan struct{})
	gen := generator(qsig)

	max := 10_000_000
	limitCh := limit(max, gen)
	for range limitCh {
		// fmt.Printf("value %d\n", n)
	}
	close(qsig)
}
