// https://blog.stackademic.com/go-concurrency-visually-explained-semaphore-3ffe23f11388

// Purpose:

// A semaphore is a signaling mechanism that controls access to a shared resource. It maintains a count and allows a certain number of threads to access the shared resource simultaneously.
// Example: Print Job Queue

// Consider a scenario where there are multiple printers, and a limited number of print jobs can be processed concurrently. A semaphore is used to control access to the printers. Each printer corresponds to a semaphore, and the semaphore count represents the available printers. Threads (print jobs) acquire a semaphore before using a printer and release it afterward.

package main

import (
	"context"
	"log"
	// "log"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
)

// var log, _ = zap.NewProduction()

var sugar = zap.NewExample().Sugar()

func main() {
	// log := logger
	pool := semaphore.NewWeighted(2)
	go swim("Candier", pool)
	go swim("Swimmer", pool)
	go swim("Partier", pool)
	go swim("Stringer", pool)

	time.Sleep(5 * time.Second) // For brevity, better use sync.WaitGroup
	sugar.Info("Main: Done, shutting down")
}

func swim(name string, pool *semaphore.Weighted) {
	sugar.Info("I want to swim", "name", name)

	// In real applications, pass in your context such as HTTP request context
	ctx := context.Background()

	// We can also Acquire/Release more than 1
	// when the workloads consume different amount of resources
	if err := pool.Acquire(ctx, 1); err != nil {
		sugar.Info("Ops, something went wrong! I cannot acquire a lane", "name", name)
		return
	}

	log.Printf("%v: I got a lane, I'm swimming", name)
	time.Sleep(time.Second)
	sugar.Info("I'm done. Releasing my lane", "name", name)
	pool.Release(1)
}
