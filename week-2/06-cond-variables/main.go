package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var mu sync.Mutex
	var cond = sync.NewCond(&mu)

	const targetCount = 100

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				mu.Lock()
				if count >= targetCount {
					mu.Unlock()
					break
				}
				count++
				fmt.Printf("Incremented count to %d\n", count)
				if count == targetCount {
					cond.Broadcast() // Signal all waiting goroutines
				}
				mu.Unlock()
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		for count < targetCount {
			cond.Wait() // Wait until count reaches targetCount
		}
		mu.Unlock()
		fmt.Println("Reached target count!")
	}()

	wg.Wait()
	fmt.Printf("Final count: %d\n", count)
}
