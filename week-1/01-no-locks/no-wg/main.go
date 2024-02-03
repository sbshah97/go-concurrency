package main

import (
	"fmt"
	"time"
)

var count int

func increment() {
	count++ // No lock acquired
	time.Sleep(50 * time.Millisecond)

}

func readCount() int {
    return count
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	// Read while writing
	fmt.Println("Read count while writes are happening", readCount())
	fmt.Println("Read count while writes are happening", readCount())
	fmt.Println("Read count while writes are happening", readCount())

	// Does not wait for all goroutines to finish
	// Output is less than 1000 because of race + threads still running

	fmt.Println("Final count:", count) // Output might be less than 1000 due to race conditions
}
