package main

import (
	"fmt"
	"sync"
	"time"
)

// RealLife examples for Barrier
// 1. Pedestrian Crosswalk: Imagine a busy intersection with a crosswalk. A barrier (traffic light) ensures all vehicles stop before pedestrians can cross. Once all vehicles have stopped, the barrier (light turns green) releases, allowing pedestrians to cross safely.

// 2. Airport Baggage Claim: Arriving passengers wait at a barrier (belt) until their luggage is unloaded and sorted. Once everyone's luggage is on the belt, the barrier (signal or announcement) releases, allowing passengers to claim their bags.

// 3. School Field Trip Permission Slips: A teacher collects permission slips from all students before a field trip. The barrier (collected slips) ensures no one goes on the trip without consent. Once all slips are collected, the barrier is lifted, and the trip proceeds.

type Barrier struct {
	n           int // Number of participants
	m           sync.Mutex
	count       int // Count of arrived participants
	triggerOnce sync.Once
}

func NewBarrier(n int) *Barrier {
	if n <= 0 {
		panic("Barrier: number of participants must be positive")
	}
	return &Barrier{
		n:           n,
		m:           sync.Mutex{},
		count:       0,
		triggerOnce: sync.Once{},
	}
}

func (b *Barrier) Wait() {
	b.m.Lock()
	defer b.m.Unlock()

	b.count++

	if b.count == b.n {
		b.triggerOnce.Do(func() {
			// Release all waiting goroutines
			for i := 0; i < b.n; i++ {
				fmt.Println("releasing routines")
				b.count--
			}
		})
	} else {
		// Wait for the barrier to be released
		for b.count != b.n {
			b.m.Unlock()
			b.m.Lock()
		}
	}
}

func worker(id int, barrier *Barrier) {
	fmt.Printf("Worker %d: Start\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d: Reached the barrier\n", id)
	barrier.Wait()
	fmt.Printf("Worker %d: Passed the barrier\n", id)
}

func main() {
	const numWorkers = 5

	// Create a Barrier with a count of the number of workers
	barrier := NewBarrier(numWorkers)

	// Launch multiple goroutines (workers)
	for i := 0; i < numWorkers-1; i++ {
		go worker(i, barrier)
	}

	// Wait for all workers to reach the barrier
	time.Sleep(time.Second)
	fmt.Println("Main: Waiting for all workers to reach the barrier")
	barrier.Wait()
	fmt.Println("Main: All workers passed the barrier")

	// The program will exit when all goroutines have finished
}
