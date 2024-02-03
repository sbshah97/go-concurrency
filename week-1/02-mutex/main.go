package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	mutex.Lock() // Acquire the lock
	defer mutex.Unlock()                    // Release the lock
	
	count++
	time.Sleep(time.Millisecond * 50) // Simulate some work
	
	// Complete wait group
	wg.Done()
}

func readCount() int {
    mutex.Lock() // Acquire a read lock
    defer mutex.Unlock()
    return count
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	fmt.Println("Read count while writes are happening", readCount())
	fmt.Println("Read count while writes are happening", readCount())
	fmt.Println("Read count while writes are happening", readCount())
	wg.Wait()
	// Wait for all goroutines to finish
	// (this is not required for correctness, but it's good practice)
	// You can use a WaitGroup or similar mechanism for this

	fmt.Println("Final count:", count)
}
