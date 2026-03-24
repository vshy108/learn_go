//go:build ignore

// Section 14, Topic 104: Buffered Channels
//
// Buffered channels have capacity. Sends don't block until the buffer is full.
//   ch := make(chan Type, capacity)
//
// Unbuffered (capacity 0): send blocks until receiver is ready.
// Buffered (capacity N): send blocks only when buffer is full.
//
// GOTCHA: A full buffered channel blocks the sender.
// GOTCHA: Buffered channels can mask synchronization bugs.
//
// Run: go run examples/s14_buffered_channels.go

package main

import "fmt"

func main() {
	fmt.Println("=== Buffered Channels ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Buffered channel — no blocking until full
	// ─────────────────────────────────────────────
	ch := make(chan int, 3)
	ch <- 1 // doesn't block — buffer has space
	ch <- 2
	ch <- 3
	// ch <- 4  // would block! buffer is full

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3

	// ─────────────────────────────────────────────
	// 2. cap() and len() on channels
	// ─────────────────────────────────────────────
	fmt.Println("\n-- cap and len --")
	ch2 := make(chan string, 5)
	ch2 <- "a"
	ch2 <- "b"
	fmt.Printf("len: %d, cap: %d\n", len(ch2), cap(ch2)) // len=2, cap=5

	// ─────────────────────────────────────────────
	// 3. Unbuffered vs Buffered
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Unbuffered vs Buffered --")

	// Unbuffered: synchronous — sender and receiver must meet
	// make(chan int)     — capacity 0
	// make(chan int, 0)  — same thing

	// Buffered: asynchronous — up to capacity
	// make(chan int, 10) — capacity 10

	// ─────────────────────────────────────────────
	// 4. When to use buffered channels
	// ─────────────────────────────────────────────
	// - Known number of results (e.g., fan-out with known goroutine count)
	// - Rate limiting (channel as semaphore)
	// - Decoupling producer from consumer (slight)
	//
	// Prefer unbuffered when synchronization is needed.
	// Use buffered when you know the capacity ahead of time.

	// ─────────────────────────────────────────────
	// 5. Buffered channel as semaphore
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Semaphore pattern --")
	sem := make(chan struct{}, 3) // max 3 concurrent operations

	for i := 0; i < 10; i++ {
		sem <- struct{}{} // acquire (blocks when 3 are running)
		go func(id int) {
			defer func() { <-sem }() // release
			fmt.Printf("  Task %d running\n", id)
		}(i)
	}

	// Drain semaphore to wait for completion:
	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
	fmt.Println("All tasks complete")
}
