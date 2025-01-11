package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	in := make(chan int, 1)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for n := range in {
				fmt.Println(n * 100)
				time.Sleep(time.Second * 10)
			}

			wg.Done()
		}()
	}

	for i := 10; i < 20; i++ {
		in <- i
	}

	// in <- -1
	close(in)
	wg.Wait()
}
