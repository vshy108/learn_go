//go:build ignore

// Section 13, Topic 100: Race Conditions and -race Flag
//
// Data race: two goroutines access the same variable concurrently,
// and at least one is a write.
//
// Go provides a built-in race detector:
//   go run -race program.go
//   go test -race ./...
//
// GOTCHA: Race detector only finds races that actually execute.
// GOTCHA: No data race ≠ no concurrency bugs (logic races still possible).
//
// Run: go run -race examples/s13_race_conditions.go

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("=== Race Conditions ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Data race example (run with -race to detect)
	// ─────────────────────────────────────────────
	// The race detector would flag this:
	// var count int
	// var wg sync.WaitGroup
	// for i := 0; i < 100; i++ {
	//     wg.Add(1)
	//     go func() { defer wg.Done(); count++ }()
	// }
	// wg.Wait()
	// fmt.Println(count) // undefined behavior!

	// ─────────────────────────────────────────────
	// 2. Fix with Mutex
	// ─────────────────────────────────────────────
	fmt.Println("-- Fix: Mutex --")
	var mu sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("Mutex count: %d\n", count)

	// ─────────────────────────────────────────────
	// 3. Fix with atomic operations
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Fix: atomic --")
	var atomicCount atomic.Int64
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomicCount.Add(1)
		}()
	}
	wg.Wait()
	fmt.Printf("Atomic count: %d\n", atomicCount.Load())

	// ─────────────────────────────────────────────
	// 4. Fix with channels
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Fix: channel --")
	ch := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1
		}()
	}
	wg.Wait()
	close(ch)

	total := 0
	for v := range ch {
		total += v
	}
	fmt.Printf("Channel count: %d\n", total)

	// ─────────────────────────────────────────────
	// 5. Race detector usage
	// ─────────────────────────────────────────────
	// go run -race main.go
	// go test -race ./...
	// go build -race -o app ./cmd/app
	//
	// Output on race:
	// WARNING: DATA RACE
	// Write at 0x00c000126008 by goroutine 7:
	//   ...
	// Previous read at 0x00c000126008 by goroutine 8:
	//   ...
}
