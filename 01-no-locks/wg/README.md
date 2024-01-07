## Race Conditions using WaitGroups:

**Prerequisites:**

- Ensure you have Golang installed. Follow the official installation instructions at [https://go.dev/doc/install](https://go.dev/doc/install): [https://go.dev/doc/install](https://go.dev/doc/install)

**Executing:**

1. Open a terminal in the directory containing the `main.go` file.
2. Run the following command:

   ```bash
   go run main.go
   ```

**Output:**

You'll observe multiple lines of output, including:

- **Read count while writes are happening:** These lines reflect inconsistent count values due to concurrent modifications.
- **Final count:** This final count might still be less than 1000, even though WaitGroups are used.

**Key Points:**

- **WaitGroups for Goroutine Synchronization:** The code demonstrates how WaitGroups help synchronize goroutine completion, ensuring the main function waits for all goroutines to finish before proceeding.
- **Race Conditions Still Present:** Despite WaitGroups, race conditions persist because multiple goroutines access a shared variable (`count`) without proper synchronization.
- **Synchronization Still Necessary:** This highlights that WaitGroups alone don't prevent race conditions. Mutexes or other synchronization mechanisms are essential for safe concurrent data access.

**Next Steps:**

- Introduce a mutex to protect the shared variable `count` and observe the corrected output.
- Experiment with different synchronization techniques to gain a deeper understanding of their roles in concurrent programming.
- Continue exploring Go's concurrency model and best practices for building reliable concurrent applications.
