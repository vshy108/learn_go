//go:build ignore

// Section 14, Topic 106: Range Over Channels and Closing
//
// close(ch): Signals no more values will be sent.
// range ch: Receives values until channel is closed.
//
// GOTCHA: Only the sender should close a channel (never the receiver).
// GOTCHA: Sending on a closed channel causes PANIC.
// GOTCHA: Receiving from a closed channel returns the zero value immediately.
// GOTCHA: Closing a nil channel panics.
// GOTCHA: Closing an already-closed channel panics.
//
// Run: go run examples/s14_range_close.go

package main

import "fmt"

func main() {
	fmt.Println("=== Range Over Channels ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Range over channel
	// ─────────────────────────────────────────────
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("  %d\n", v)
	}
	fmt.Println("Channel closed, range exited")

	// ─────────────────────────────────────────────
	// 2. Detecting closed channel (comma-ok)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Comma-ok --")
	ch2 := make(chan string, 2)
	ch2 <- "hello"
	close(ch2)

	v, ok := <-ch2
	fmt.Printf("value=%q, ok=%t\n", v, ok) // "hello", true

	v, ok = <-ch2
	fmt.Printf("value=%q, ok=%t\n", v, ok) // "", false (closed)

	// ─────────────────────────────────────────────
	// 3. Receiving from closed channel returns zero
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Closed channel returns zero --")
	ch3 := make(chan int, 1)
	ch3 <- 42
	close(ch3)

	fmt.Println(<-ch3) // 42
	fmt.Println(<-ch3) // 0 (zero value, channel closed)
	fmt.Println(<-ch3) // 0 (can receive forever)

	// ─────────────────────────────────────────────
	// 4. GOTCHA: Panics
	// ─────────────────────────────────────────────
	// close(nilCh)       → panic: close of nil channel
	// close(closedCh)    → panic: close of closed channel
	// closedCh <- value  → panic: send on closed channel

	// ─────────────────────────────────────────────
	// 5. When to close channels
	// ─────────────────────────────────────────────
	// Close when:
	//   - Using range to iterate (receiver needs to know when to stop)
	//   - Signaling completion to multiple receivers
	//
	// Don't close when:
	//   - Only one receiver (use done channel or WaitGroup instead)
	//   - Channel is shared by multiple senders (who closes?)
}
