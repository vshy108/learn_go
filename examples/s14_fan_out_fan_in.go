//go:build ignore

// Section 14, Topic 111: Fan-Out / Fan-In Pattern
//
// Fan-Out: distribute work to multiple goroutines.
// Fan-In: merge results from multiple channels into one.
//
// This is a key pattern for parallel processing in Go.
//
// Run: go run examples/s14_fan_out_fan_in.go

package main

import (
	"fmt"
	"sync"
	"time"
)

// produce generates n messages on a channel
func produce(prefix string, n int, delay time.Duration) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			time.Sleep(delay)
			ch <- fmt.Sprintf("%s-%d", prefix, i)
		}
	}()
	return ch
}

// fanIn merges multiple channels into one
func fanIn(channels ...<-chan string) <-chan string {
	merged := make(chan string)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for val := range c {
				merged <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	fmt.Println("=== Fan-Out / Fan-In ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Fan-Out: 3 producers running concurrently
	// ─────────────────────────────────────────────
	ch1 := produce("A", 3, 20*time.Millisecond)
	ch2 := produce("B", 3, 30*time.Millisecond)
	ch3 := produce("C", 3, 10*time.Millisecond)

	// ─────────────────────────────────────────────
	// 2. Fan-In: merge all into one channel
	// ─────────────────────────────────────────────
	merged := fanIn(ch1, ch2, ch3)

	fmt.Println("Merged results (order varies):")
	for v := range merged {
		fmt.Printf("  %s\n", v)
	}

	// ─────────────────────────────────────────────
	// 3. Fan-Out workers pattern
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Fan-Out Workers --")
	jobs := []int{1, 2, 3, 4, 5, 6}
	results := make(chan string, len(jobs))

	var wg sync.WaitGroup
	numWorkers := 3
	jobCh := make(chan int, len(jobs))

	// Start workers (fan-out):
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobCh {
				time.Sleep(10 * time.Millisecond)
				results <- fmt.Sprintf("worker-%d processed job-%d", id, job)
			}
		}(w)
	}

	// Send jobs:
	for _, j := range jobs {
		jobCh <- j
	}
	close(jobCh)

	// Collect (fan-in):
	go func() { wg.Wait(); close(results) }()
	for r := range results {
		fmt.Printf("  %s\n", r)
	}
}
