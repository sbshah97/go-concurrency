package main

import (
	"fmt"
	"sync"
	"time"
)

var count int

func increment(wg *sync.WaitGroup) {
	count++ // No lock acquired
	time.Sleep(50 * time.Millisecond)
	wg.Done()
}

func readCount() int {
    return count
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// Read while writing
	fmt.Println("Read count while writes are happening", readCount())
	fmt.Println("Read count while writes are happening", readCount())
	fmt.Println("Read count while writes are happening", readCount())

	wg.Wait()
	// Waits for all goroutines to finish

	fmt.Println("Final count:", count) // Output might be less than 1000 due to race conditions
}
