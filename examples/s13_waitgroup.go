//go:build ignore

// Section 13, Topic 98: sync.WaitGroup
//
// WaitGroup waits for a collection of goroutines to finish.
//   wg.Add(n)   — increment counter by n
//   wg.Done()   — decrement counter (typically deferred)
//   wg.Wait()   — block until counter reaches 0
//
// GOTCHA: Call Add() BEFORE starting the goroutine, not inside it.
// GOTCHA: Passing WaitGroup by value copies it — always use pointer.
//
// Run: go run examples/s13_waitgroup.go

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== sync.WaitGroup ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic WaitGroup usage
	// ─────────────────────────────────────────────
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment BEFORE go
		go func(id int) {
			defer wg.Done() // decrement when done
			time.Sleep(time.Duration(id*10) * time.Millisecond)
			fmt.Printf("  Worker %d done\n", id)
		}(i)
	}

	wg.Wait() // blocks until all workers finish
	fmt.Println("All workers complete!")

	// ─────────────────────────────────────────────
	// 2. Passing WaitGroup to function
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Passing WaitGroup --")
	var wg2 sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg2.Add(1)
		go worker(i, &wg2) // MUST pass pointer!
	}
	wg2.Wait()
	fmt.Println("All workers done")

	// ─────────────────────────────────────────────
	// 3. GOTCHA: Add before go
	// ─────────────────────────────────────────────
	// BAD:
	//   go func() {
	//       wg.Add(1)  // Race! Wait() might return before Add()
	//       defer wg.Done()
	//       ...
	//   }()
	//   wg.Wait()
	//
	// GOOD:
	//   wg.Add(1)     // Add in the launching goroutine
	//   go func() {
	//       defer wg.Done()
	//       ...
	//   }()
	//   wg.Wait()

	// ─────────────────────────────────────────────
	// 4. Batch Add
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Batch Add --")
	n := 3
	var wg3 sync.WaitGroup
	wg3.Add(n) // add all at once
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg3.Done()
			fmt.Printf("  Task %d\n", id)
		}(i)
	}
	wg3.Wait()
	fmt.Println("Batch complete")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("  Worker %d running\n", id)
}
