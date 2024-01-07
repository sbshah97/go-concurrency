# Understanding Mutexes

- Exclusive access: Only one goroutine can hold the lock at a time.
- Used for both reads and writes.
- Simpler to use, but can limit concurrency if reads are frequent.

## Race Condition Prevention with Mutex

**Prerequisites:**

- Ensure you have Golang installed. Follow the official installation instructions at [https://go.dev/doc/install](https://go.dev/doc/install): [https://go.dev/doc/install](https://go.dev/doc/install).

**Executing:**

1. Open a terminal in the directory containing the `main.go` file.
2. Run the following command:

   ```bash
   go run main.go
   ```

**Output:**

You'll observe multiple lines of output, including:

- **Read count while writes are happening:** These lines now accurately reflect the count at different points in time due to synchronization.
- **Final count:** The final count will consistently be 1000, demonstrating successful race condition prevention.

**Key Points:**

- **Mutex for Data Synchronization:** The code demonstrates how a mutex (Mutual Exclusion) ensures consistent access to the shared variable `count` in a concurrent environment.
- **Lock Acquisition and Release:** Goroutines acquire the mutex lock before modifying `count` and release it afterward, preventing race conditions.
- **WaitGroup for Goroutine Coordination:** The `sync.WaitGroup` ensures the main function waits for all goroutines to complete before printing the final count.

**Next Steps:**

- Experiment with different synchronization techniques like RWMutex for read-mostly scenarios.
- Explore other concurrency primitives like channels for communication and coordination.
- Learn more about Go's memory model and best practices for building robust concurrent applications.
