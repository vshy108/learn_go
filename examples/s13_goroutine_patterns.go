//go:build ignore

// Section 13, Topic 101: Goroutine Patterns
//
// Common patterns for structuring concurrent Go code.
//
// Run: go run examples/s13_goroutine_patterns.go

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Goroutine Patterns ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Fan-out: one goroutine launches many
	// ─────────────────────────────────────────────
	fmt.Println("-- Fan-out --")
	var wg sync.WaitGroup
	jobs := []string{"a", "b", "c", "d", "e"}
	for _, job := range jobs {
		wg.Add(1)
		go func(j string) {
			defer wg.Done()
			fmt.Printf("  Processing: %s\n", j)
		}(job)
	}
	wg.Wait()

	// ─────────────────────────────────────────────
	// 2. Worker pool
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Worker pool --")
	tasks := make(chan int, 10)
	results := make(chan string, 10)

	// Start 3 workers:
	for w := 1; w <= 3; w++ {
		go worker(w, tasks, results)
	}

	// Send 6 tasks:
	for t := 1; t <= 6; t++ {
		tasks <- t
	}
	close(tasks)

	// Collect results:
	for i := 0; i < 6; i++ {
		fmt.Printf("  %s\n", <-results)
	}

	// ─────────────────────────────────────────────
	// 3. Fire and forget
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Fire and forget --")
	go func() {
		// Log something asynchronously
		fmt.Println("  Background: logging event")
	}()
	time.Sleep(10 * time.Millisecond)

	// ─────────────────────────────────────────────
	// 4. Concurrent map processing
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Concurrent processing --")
	items := []int{1, 2, 3, 4, 5}
	resultsCh := make(chan int, len(items))

	for _, item := range items {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultsCh <- n * n // square each number
		}(item)
	}

	wg.Wait()
	close(resultsCh)

	var squares []int
	for r := range resultsCh {
		squares = append(squares, r)
	}
	fmt.Printf("  Squares: %v\n", squares)

	// ─────────────────────────────────────────────
	// 5. Goroutine lifecycle
	// ─────────────────────────────────────────────
	// Always ensure goroutines can terminate:
	// - Use WaitGroup to wait for completion
	// - Use channels to signal done
	// - Use context.Context for cancellation
	// - A goroutine that never terminates = goroutine leak
}

func worker(id int, tasks <-chan int, results chan<- string) {
	for task := range tasks {
		time.Sleep(10 * time.Millisecond) // simulate work
		results <- fmt.Sprintf("Worker %d processed task %d", id, task)
	}
}
