package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32 // Shared variable

	var wg sync.WaitGroup
	const numGoroutines = 100

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			atomic.AddInt32(&count, 1) // Atomic increment
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Count after atomic operations: %d\n", atomic.LoadInt32(&count))
}
