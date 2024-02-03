package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var rwMutex sync.RWMutex // RWMutex for read/write access

func increment(wg *sync.WaitGroup) {
	rwMutex.Lock() // Acquire the lock
	defer rwMutex.Unlock()                    // Release the lock
	
	count++
	time.Sleep(time.Millisecond * 50) // Simulate some work
	
	// Complete wait group
	wg.Done()
}

func readCount() int {
    rwMutex.RLock() // Acquire a read lock
	fmt.Println("Read count in the loop", count)
	time.Sleep(time.Millisecond * 100) // Simulate some work

    defer rwMutex.RUnlock()
    return count
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	go readCount()
	go readCount()
	go readCount()
	time.Sleep(time.Millisecond * 1000)
	go readCount()
	go readCount()
	// // Read while writing
	// fmt.Println("Read count while writes are happening", readCount())
	// fmt.Println("Read count while writes are happening", readCount())
	// fmt.Println("Read count while writes are happening", readCount())

	
	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final count:", count)
}
