// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/dA7TBdoL5S

// go build -race

// Sample program to show how to use the atomic package functions
// Store and Load to provide safe access to numeric types.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// shutdown is a flag to alert running goroutines to shutdown.
var shutdown int64

// main is the entry point for the application.
func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)

	// Create two goroutines.

	go func() {
		doWork("A")
		wg.Done()
	}()

	go func() {
		doWork("B")
		wg.Done()
	}()

	// Give the goroutines time to run so we can see
	// the shutdown flag work.
	time.Sleep(time.Second)

	// Safely flag it is time to shutdown.
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	// Wait for the goroutines to finish.
	wg.Wait()
}

// doWork simulates a goroutine performing work and
// checking the Shutdown flag to terminate early.
func doWork(name string) {
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// Do we need to shutdown.
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
