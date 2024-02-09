// https://www.sobyte.net/post/2022-07/go-sync-cond/
// https://kaviraj.me/understanding-condition-variable-in-go/
package main

import (
	"fmt"
	"sync"
	"time"
)

// below program demonstrates the need for the condition variable. In case of one thread and a lock each thread have to loop and and check the variable and unloack the mutex again and again.
// This will consume the CPU cycles and is not efficient.
type Record struct {
	sync.Mutex
	data string
}

func main1() {
	var wg sync.WaitGroup

	rec := &Record{}
	wg.Add(1)
	go func(rec *Record) {
		defer wg.Done()
		for {
			rec.Lock()
			if rec.data != "" {
				time.Sleep(10 * time.Second)
				fmt.Println("Data: ", rec.data)
				rec.Unlock()
				return
			}
			fmt.Println("checking data in thread")
			rec.Unlock()
		}
	}(rec)

	fmt.Println("locking the record")

	rec.Lock()
	fmt.Println("main waits for 10 seconds")
	time.Sleep(10 * time.Second)
	rec.data = "gopher"
	rec.Unlock()
	fmt.Println("main unlocks record")

	wg.Wait() // wait till all goutine completes
}

// The most important thing to note is that the goroutines are waiting for the sharedRsc to change.
// .wait internally releases the lock and waits for a signal to be sent by the main goroutine.
// Thats the reason other threads are also able to aquite the lock after the wait from the first thread.
var sharedRsc = false

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait()
		}
		fmt.Println("goroutine1", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			c.Wait()
		}
		fmt.Println("goroutine2", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine3 wait")
			c.Wait()
		}
		fmt.Println("goroutine3", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine4 wait")
			c.Wait()
		}
		fmt.Println("goroutine4", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	// this one writes changes to sharedRsc
	time.Sleep(2 * time.Second)
	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	c.Broadcast()
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()
	wg.Wait()
}
