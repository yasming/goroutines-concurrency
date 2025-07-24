# Goroutines Concurrency Demo

A simple Go project demonstrating safe concurrent programming using goroutines, mutexes, and WaitGroups.

## Overview

This project showcases fundamental concurrency concepts in Go by implementing a thread-safe counter that is incremented by multiple goroutines simultaneously. It demonstrates how to prevent race conditions using mutexes and how to coordinate goroutines using WaitGroups.

## What it does

The program:
- Creates 10 concurrent goroutines
- Each goroutine increments a shared global counter 1,000 times
- Uses a mutex to ensure thread-safe access to the counter
- Uses a WaitGroup to wait for all goroutines to complete
- Prints the final counter value (should always be 10,000)

## Key Concepts Demonstrated

- **Goroutines**: Lightweight threads managed by the Go runtime
- **Mutex**: Mutual exclusion lock to prevent race conditions
- **WaitGroup**: Synchronization primitive to wait for multiple goroutines to finish
- **Race Condition Prevention**: Safe concurrent access to shared resources

## Project Structure

```
.
├── README.md
├── go.mod
├── main.go
└── Makefile
```

## Prerequisites

- Go 1.23.5 or later

## Running the Project

### Basic execution:
```bash
go run main.go
```

### Race condition detection:
```bash
make check-concurrency
```
or
```bash
go run -race main.go
```

The `-race` flag enables Go's race detector, which helps identify potential race conditions in concurrent code.

## Expected Output

```
Final counter: 10000
```

## Code Explanation

### Global Variables
- `counter int`: Shared variable incremented by all goroutines
- `mutex sync.Mutex`: Protects the counter from concurrent access

### Main Function
1. Creates a WaitGroup to track goroutine completion
2. Launches 10 goroutines, each calling the `incremental` function
3. Waits for all goroutines to finish using `wg.Wait()`
4. Prints the final counter value

### Incremental Function
1. Increments the counter 1,000 times
2. Uses mutex lock/unlock to ensure thread-safe access
3. Calls `wg.Done()` when finished to signal completion

## Learning Points

This example illustrates:
- How to safely share data between goroutines
- The importance of mutexes in preventing race conditions
- How to coordinate multiple goroutines using WaitGroups
- Best practices for concurrent programming in Go

## Try This

1. **Remove the mutex**: Comment out the `mutex.Lock()` and `mutex.Unlock()` lines and run with `-race` to see race condition warnings
2. **Increase goroutines**: Change the loop to create more goroutines and observe the behavior
3. **Increase iterations**: Modify the inner loop to increment more times per goroutine

## Race Condition Warning

Without proper synchronization (mutex), multiple goroutines accessing the shared `counter` variable simultaneously would create a race condition, leading to unpredictable and incorrect results.
