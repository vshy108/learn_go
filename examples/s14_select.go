//go:build ignore

// Section 14, Topic 107: Select Statement
//
// select lets a goroutine wait on multiple channel operations.
// Like switch, but each case is a channel operation.
//
// GOTCHA: If multiple cases are ready, one is chosen at RANDOM.
// GOTCHA: select without default blocks until a case is ready.
// GOTCHA: select with default never blocks (non-blocking).
//
// Run: go run examples/s14_select.go

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Select Statement ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic select
	// ─────────────────────────────────────────────
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(30 * time.Millisecond)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "two"
	}()

	// Wait for whichever arrives first:
	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg) // likely this one (faster)
	}

	// ─────────────────────────────────────────────
	// 2. Select with timeout
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Timeout --")
	ch := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Got:", res)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timeout!")
	}

	// ─────────────────────────────────────────────
	// 3. Non-blocking with default
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Non-blocking --")
	ch3 := make(chan int, 1)

	// Non-blocking receive:
	select {
	case v := <-ch3:
		fmt.Println("Received:", v)
	default:
		fmt.Println("No value ready")
	}

	// Non-blocking send:
	select {
	case ch3 <- 42:
		fmt.Println("Sent 42")
	default:
		fmt.Println("Channel full")
	}

	// ─────────────────────────────────────────────
	// 4. Select in a loop (event loop pattern)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Event loop --")
	tick := time.NewTicker(30 * time.Millisecond)
	defer tick.Stop()
	done := make(chan bool)

	go func() {
		time.Sleep(130 * time.Millisecond)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("  Done!")
			goto end
		case t := <-tick.C:
			fmt.Printf("  Tick at %s\n", t.Format("04:05.000"))
		}
	}
end:

	// ─────────────────────────────────────────────
	// 5. Empty select blocks forever
	// ─────────────────────────────────────────────
	// select {} // blocks forever (useful for keeping main alive)
}
