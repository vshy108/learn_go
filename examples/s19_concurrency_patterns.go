//go:build ignore

// Section 19, Topic 144: Concurrency Patterns (Worker Pool, Rate Limiter, etc.)
//
// Advanced concurrency patterns commonly used in production Go.
//
// Run: go run examples/s19_concurrency_patterns.go

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Concurrency Patterns ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Worker Pool
	// ─────────────────────────────────────────────
	fmt.Println("-- Worker Pool --")
	jobs := make(chan int, 10)
	results := make(chan string, 10)

	// Start pool of 3 workers:
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				time.Sleep(10 * time.Millisecond)
				results <- fmt.Sprintf("worker %d: job %d done", id, job)
			}
		}(w)
	}

	// Submit jobs:
	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait and collect:
	go func() { wg.Wait(); close(results) }()
	for r := range results {
		fmt.Printf("  %s\n", r)
	}

	// ─────────────────────────────────────────────
	// 2. Semaphore (bounded concurrency)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Semaphore --")
	sem := make(chan struct{}, 2) // max 2 concurrent
	var wg2 sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			sem <- struct{}{}        // acquire
			defer func() { <-sem }() // release

			fmt.Printf("  Task %d: running (max 2 concurrent)\n", id)
			time.Sleep(20 * time.Millisecond)
		}(i)
	}
	wg2.Wait()

	// ─────────────────────────────────────────────
	// 3. Rate Limiter with time.Ticker
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Rate Limiter --")
	rateLimiter := time.NewTicker(30 * time.Millisecond)
	defer rateLimiter.Stop()

	for i := 1; i <= 3; i++ {
		<-rateLimiter.C // wait for tick
		fmt.Printf("  Request %d at %s\n", i, time.Now().Format("04:05.000"))
	}

	// ─────────────────────────────────────────────
	// 4. Context-based cancellation
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Context cancellation --")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	ch := make(chan string, 1)
	go func() {
		time.Sleep(30 * time.Millisecond)
		ch <- "result"
	}()

	select {
	case result := <-ch:
		fmt.Println("  Got:", result)
	case <-ctx.Done():
		fmt.Println("  Timeout!")
	}

	// ─────────────────────────────────────────────
	// 5. errgroup pattern (conceptual)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- errgroup (concept) --")
	fmt.Println(`
import "golang.org/x/sync/errgroup"

g, ctx := errgroup.WithContext(context.Background())

g.Go(func() error {
    return fetchData(ctx, "url1")
})

g.Go(func() error {
    return fetchData(ctx, "url2")
})

if err := g.Wait(); err != nil {
    // First error cancels all goroutines
    log.Fatal(err)
}
`)
}
