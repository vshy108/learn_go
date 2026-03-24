//go:build ignore

// Section 13, Topic 97: Goroutine Basics
//
// A goroutine is a lightweight thread managed by the Go runtime.
// Start one with the `go` keyword: go f()
//
// Goroutines are multiplexed onto OS threads (M:N scheduling).
// They start with ~2KB stack (grows as needed). You can run millions.
//
// GOTCHA: main() is itself a goroutine. When main returns, ALL goroutines die.
// GOTCHA: Goroutine execution order is non-deterministic.
// GOTCHA: No way to "kill" a goroutine from outside — use channels/context.
//
// Run: go run examples/s13_goroutine_basics.go

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Goroutine Basics ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Starting a goroutine
	// ─────────────────────────────────────────────
	go sayHello("World")
	go sayHello("Go")
	go sayHello("Goroutine")

	// Must wait or main exits immediately:
	time.Sleep(100 * time.Millisecond)

	// ─────────────────────────────────────────────
	// 2. Anonymous goroutine (closure)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Anonymous goroutine --")
	go func() {
		fmt.Println("  Hello from anonymous goroutine!")
	}()

	go func(msg string) {
		fmt.Printf("  Message: %s\n", msg)
	}("passed as argument")

	time.Sleep(50 * time.Millisecond)

	// ─────────────────────────────────────────────
	// 3. GOTCHA: Main exits = all goroutines die
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Main goroutine --")
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("This will NEVER print if main exits first")
	}()
	fmt.Println("Main is about to finish this section")
	// The goroutine above won't complete

	// ─────────────────────────────────────────────
	// 4. GOTCHA: Closure variable capture
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Closure gotcha --")
	// BAD: All goroutines share the same `i`
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("  BAD: i=%d\n", i) // likely prints 3, 3, 3
		}()
	}
	time.Sleep(50 * time.Millisecond)

	// GOOD: Pass i as argument
	fmt.Println("  ---")
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("  GOOD: n=%d\n", n) // prints 0, 1, 2 (any order)
		}(i)
	}
	time.Sleep(50 * time.Millisecond)

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust
	// ─────────────────────────────────────────────
	// Go:   go f()              — green thread, GC manages memory
	// Rust: std::thread::spawn  — OS thread, ownership ensures safety
	//       tokio::spawn        — async task (with async runtime)
	// Go's goroutines are cheaper and simpler to use.
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
