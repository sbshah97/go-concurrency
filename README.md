# Go Concurrency Deep Dive with Salman Shah
Understand Go's concurrency with this comprehensive series by Salman Shah, available on his YouTube channel.

This project delves into the intricacies of Go's concurrency model, demonstrating various synchronization techniques to ensure data consistency in multi-threaded environments. It's designed to solidify your understanding of concepts like goroutines, mutexes, RWMutexes, and WaitGroups.

# Project Structure

## Week 1
- **01-no-locks:** Explores the consequences of concurrent access without synchronization.
    - **no-wg:** Showcases race conditions without WaitGroups.
    - **wg:** Introduces WaitGroups for goroutine coordination, but race conditions persist.
- **02-mutex:** Demonstrates resolving race conditions using mutexes for exclusive access.
- **03-rwmutex:** Explores RWMutex for optimized performance in read-heavy scenarios.

## Week 2
- 

**Prerequisites**

- Install Go: [https://go.dev/doc/install](https://go.dev/doc/install)

**Executing Code**

1. Open a terminal in the project directory.
2. Use the provided Makefile to run specific examples:

   ```bash
   make run-no-lock-no-wg   # Run without locks or WaitGroups
   make run-no-lock-wg     # Run with WaitGroups but no locks
   make run-mutex          # Run with mutex for synchronization
   make run-rwmutex        # Run with RWMutex for optimized read access
   ```

**Observations**

- Analyze the output of each example to understand the effects of different synchronization approaches.
- Observe race conditions, inconsistent results, and how synchronization resolves them.
- Compare the performance impact of mutexes and RWMutex in different scenarios.


## Dive into the World of Multithreaded Go
As Go applications grow in complexity, understanding concurrency becomes crucial for building efficient and performant systems. However, grasping these concepts can be challenging, even for experienced developers. To bridge this gap, this insightful series on Go Concurrency provides a clear and engaging learning experience.

## What You'll Learn
This series delves into a treasure trove of concurrency topics, equipping you with practical knowledge and intuition:

* Wait Groups and how they are useful
* Mutexes and how they're useful
* RwMutexes and how they're useful
* .... soon adding more

## Get Started, Level Up your Skills
To embark on your learning journey, simply head to Salman Shah's YouTube channel and explore the "Go Concurrency Deep Dive" series. Each video comes with clear explanations, practical examples, and insightful discussions, making it an ideal learning resource for developers of all skill levels.

By diving into this series, you'll gain the confidence and expertise to tackle even the most intricate concurrency challenges in your Go projects. Let Salman Shah be your guide, and elevate your skills to the next level!

**Next Steps**

- Experiment with code modifications to reinforce your understanding.
- Explore additional concurrency features like channels and atomic operations.
- Dive deeper into Go's memory model for a comprehensive grasp of concurrent programming.

Don't forget to check out:
- Salman Shah's [YouTube channel](https://www.youtube.com/channel/UCZF43Gp9WiEBmZD1xgrzdoA)

## TODO
- [ ] The "Go Concurrency Deep Dive" series playlist: (link to playlist)

## Roadmap
1. Goroutines and Channels:

- [ ] Goroutines: Introduce lightweight threads as the foundation of Go's concurrency.
- [ ] Launching goroutines: go keyword.
- [ ] Synchronization and communication: Channels for data exchange and coordination.
- [ ] Channel types: Unbuffered vs. buffered channels.
- [ ] Channel operations: Send, receive, close.
- [ ] Range over channels: Iterating until closed.
- [ ] Select statement: Managing multiple channels.

2. Synchronization Primitives:

- [x] WaitGroups: Synchronizing goroutine completion.
    - [x] Add(), Done(), Wait() methods.
- [x] Mutex: Exclusive access to shared resources.
    - [x] Lock(), Unlock() methods.
- [x] RWMutex: Reader-writer locks for efficient read-heavy scenarios.
    - [x] RLock(), RUnlock(), Lock(), Unlock() methods.
- [ ] Atomic operations: Safe, lock-free operations on shared variables.
- [ ] Context: Canceling groups of goroutines gracefully.
- [ ] Cond: Condition variables for signaling and waiting between goroutines.

3. Patterns and Practices:

- [ ] Fan-out, fan-in: Parallelizing tasks and gathering results.
- [ ] Pipelines: Building data processing pipelines with channels.
- [ ] Worker pools: Managing a pool of goroutines for efficient task handling.
- [ ] Rate limiting: Controlling the flow of data or events.
- [ ] Error handling in concurrent code: Strategies and patterns.
4. Advanced Concurrency Topics:

- [ ] Memory models: Understanding Go's memory consistency guarantees.
- [ ] Lock-free programming: Techniques for concurrency without explicit locks.
- [ ] Concurrent data structures: Designing safe data structures for concurrent access.
- [ ] Performance optimization: Tuning concurrent code for efficiency.
