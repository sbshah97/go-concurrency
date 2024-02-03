package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var count int32
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	const numGoroutines = 10
	const targetCount = 100

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Goroutine %d: Context canceled\n", id)
					return
				default:
					atomic.AddInt32(&count, 1)
					if count == targetCount {
						return // Exit goroutine when target count is reached
					}
				}
			}
		}(i)
	}

	// Simulating some work
	// Cancel the context after a while to stop goroutines
	go func() {
		// Simulating some work
		// Sleep for a while and then cancel the context
		// to signal goroutines to exit gracefully
		time.Sleep(5 * time.Second)
		defer cancel()
		// Simulating work
	}()

	wg.Wait()
	fmt.Printf("Count after reaching target: %d\n", count)
}
