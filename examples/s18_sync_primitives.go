//go:build ignore

// Section 18, Topic 135: sync Package (Once, Pool, Map)
//
// sync provides additional synchronization primitives beyond Mutex.
//
//   sync.Once  — execute exactly once (e.g., initialization)
//   sync.Pool  — reusable object pool (reduces GC pressure)
//   sync.Map   — concurrent-safe map (no manual locking)
//
// Run: go run examples/s18_sync_primitives.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== sync Primitives ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. sync.Once — initialize exactly once
	// ─────────────────────────────────────────────
	fmt.Println("-- sync.Once --")
	var once sync.Once
	var config string

	initConfig := func() {
		fmt.Println("  Initializing config (runs once)")
		config = "initialized"
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(initConfig) // only first call executes
			fmt.Printf("  Goroutine %d: config=%s\n", id, config)
		}(i)
	}
	wg.Wait()

	// ─────────────────────────────────────────────
	// 2. sync.Pool — object reuse
	// ─────────────────────────────────────────────
	fmt.Println("\n-- sync.Pool --")
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("  Creating new buffer")
			return make([]byte, 0, 1024)
		},
	}

	// Get from pool (creates new if empty):
	buf := pool.Get().([]byte)
	buf = append(buf, "hello"...)
	fmt.Printf("  Buffer: %q\n", buf)

	// Return to pool:
	buf = buf[:0] // reset
	pool.Put(buf)

	// Get again (reuses pooled object):
	buf2 := pool.Get().([]byte)
	fmt.Printf("  Reused buffer cap: %d\n", cap(buf2))

	// ─────────────────────────────────────────────
	// 3. sync.Map — concurrent-safe map
	// ─────────────────────────────────────────────
	fmt.Println("\n-- sync.Map --")
	var m sync.Map

	// Store:
	m.Store("key1", "value1")
	m.Store("key2", 42)
	m.Store("key3", true)

	// Load:
	if val, ok := m.Load("key1"); ok {
		fmt.Printf("  key1: %v\n", val)
	}

	// LoadOrStore (atomic get-or-set):
	actual, loaded := m.LoadOrStore("key4", "new")
	fmt.Printf("  key4: %v (loaded existing: %t)\n", actual, loaded)

	// Range (iterate):
	fmt.Println("  All entries:")
	m.Range(func(key, value any) bool {
		fmt.Printf("    %v: %v\n", key, value)
		return true // continue iteration
	})

	// Delete:
	m.Delete("key2")

	// ─────────────────────────────────────────────
	// When to use sync.Map vs regular map+Mutex:
	// ─────────────────────────────────────────────
	// sync.Map is better when:
	//   - Many goroutines read, few write (read-heavy)
	//   - Keys are stable (not frequently added/deleted)
	// Regular map + Mutex is better when:
	//   - You need type safety (sync.Map uses any)
	//   - Write-heavy workloads
	//   - Simple iteration patterns
}
