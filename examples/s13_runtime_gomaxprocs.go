//go:build ignore

// Section 13, Topic 102: GOMAXPROCS and Runtime
//
// GOMAXPROCS controls how many OS threads can execute Go code simultaneously.
// Default: number of CPU cores (since Go 1.5).
//
// runtime package provides goroutine and scheduler info.
//
// GOTCHA: GOMAXPROCS limits parallelism, not concurrency.
//         You can have millions of goroutines with GOMAXPROCS=1.
//
// Run: go run examples/s13_runtime_gomaxprocs.go

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("=== GOMAXPROCS and Runtime ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Current GOMAXPROCS
	// ─────────────────────────────────────────────
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0)) // 0 = query, don't change
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
	fmt.Printf("Version: %s\n", runtime.Version())

	// ─────────────────────────────────────────────
	// 2. Changing GOMAXPROCS
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Changing GOMAXPROCS --")
	prev := runtime.GOMAXPROCS(1) // single-threaded
	fmt.Printf("Previous GOMAXPROCS: %d, Now: %d\n", prev, runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(prev) // restore

	// ─────────────────────────────────────────────
	// 3. Goroutine.Goexit()
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Goexit --")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("  deferred in exiting goroutine")
		fmt.Println("  About to Goexit")
		runtime.Goexit() // terminates goroutine, runs deferred
		fmt.Println("  This never prints")
	}()
	wg.Wait()
	fmt.Println("Main continues after Goexit")

	// ─────────────────────────────────────────────
	// 4. runtime.Gosched()
	// ─────────────────────────────────────────────
	// Yields the processor, allowing other goroutines to run.
	// Rarely needed — the scheduler is usually smart enough.
	runtime.Gosched()

	// ─────────────────────────────────────────────
	// 5. Goroutine count
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Goroutine count --")
	fmt.Printf("Before: %d goroutines\n", runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {} // block forever (for counting)
		}()
	}
	fmt.Printf("After starting 10: %d goroutines\n", runtime.NumGoroutine())
	// Note: these goroutines will be cleaned up when main exits

	// ─────────────────────────────────────────────
	// Concurrency vs Parallelism
	// ─────────────────────────────────────────────
	// Concurrency: dealing with many things at once (design)
	// Parallelism: doing many things at once (execution)
	// GOMAXPROCS=1: concurrent but not parallel
	// GOMAXPROCS>1: concurrent and potentially parallel
}
