//go:build ignore

// Section 13, Topic 99: sync.Mutex
//
// Mutex provides mutual exclusion for shared data.
//   mu.Lock()   — acquire the lock (blocks if already held)
//   mu.Unlock() — release the lock
//
// Also: sync.RWMutex — multiple readers OR one writer.
//
// GOTCHA: Forgetting to unlock → deadlock. Always use defer mu.Unlock().
// GOTCHA: Mutex is NOT reentrant — locking twice from same goroutine = deadlock.
// GOTCHA: Don't copy a Mutex (pass by pointer).
//
// Run: go run examples/s13_mutex.go

package main

import (
	"fmt"
	"sync"
)

// ─────────────────────────────────────────────
// Thread-safe counter
// ─────────────────────────────────────────────
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	fmt.Println("=== sync.Mutex ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Race condition without mutex
	// ─────────────────────────────────────────────
	fmt.Println("-- Without mutex (race condition) --")
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // DATA RACE! Multiple goroutines write concurrently
		}()
	}
	wg.Wait()
	fmt.Printf("Counter (unsafe): %d (expected 1000, may differ)\n", counter)

	// ─────────────────────────────────────────────
	// 2. Fixed with Mutex
	// ─────────────────────────────────────────────
	fmt.Println("\n-- With mutex --")
	var mu sync.Mutex
	safeCounter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			safeCounter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter (safe): %d\n", safeCounter) // always 1000

	// ─────────────────────────────────────────────
	// 3. struct with embedded Mutex
	// ─────────────────────────────────────────────
	fmt.Println("\n-- SafeCounter struct --")
	sc := &SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n%5)
			sc.Inc(key)
		}(i)
	}
	wg.Wait()
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		fmt.Printf("  %s: %d\n", key, sc.Get(key))
	}

	// ─────────────────────────────────────────────
	// 4. RWMutex (multiple readers)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- RWMutex --")
	var rwMu sync.RWMutex
	data := "initial"

	// Multiple readers can read concurrently:
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rwMu.RLock()
			defer rwMu.RUnlock()
			fmt.Printf("  Reader %d: %s\n", id, data)
		}(i)
	}

	// Writer gets exclusive access:
	wg.Add(1)
	go func() {
		defer wg.Done()
		rwMu.Lock()
		defer rwMu.Unlock()
		data = "updated"
		fmt.Println("  Writer: updated data")
	}()

	wg.Wait()
	fmt.Printf("Final: %s\n", data)
}
