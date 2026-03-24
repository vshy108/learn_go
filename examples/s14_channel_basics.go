//go:build ignore

// Section 14, Topic 103: Channel Basics
//
// Channels are typed conduits for communication between goroutines.
//   ch := make(chan Type)    — unbuffered channel
//   ch <- value             — send to channel
//   value := <-ch           — receive from channel
//
// Channels are the primary synchronization mechanism in Go.
// "Don't communicate by sharing memory; share memory by communicating."
//
// GOTCHA: Unbuffered channels block until both sender and receiver are ready.
// GOTCHA: Sending to a nil channel blocks forever.
// GOTCHA: Receiving from a nil channel blocks forever.
//
// Run: go run examples/s14_channel_basics.go

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Channel Basics ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Create and use a channel
	// ─────────────────────────────────────────────
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine!" // send
	}()

	msg := <-ch // receive (blocks until message arrives)
	fmt.Println(msg)

	// ─────────────────────────────────────────────
	// 2. Channel as synchronization
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Synchronization --")
	done := make(chan bool)

	go func() {
		fmt.Println("  Working...")
		time.Sleep(50 * time.Millisecond)
		fmt.Println("  Done!")
		done <- true
	}()

	<-done // wait for goroutine to signal completion
	fmt.Println("Main: goroutine finished")

	// ─────────────────────────────────────────────
	// 3. Multiple values through channel
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multiple values --")
	nums := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			nums <- i
		}
		close(nums) // signal no more values
	}()

	for n := range nums { // range over channel until closed
		fmt.Printf("  Received: %d\n", n)
	}

	// ─────────────────────────────────────────────
	// 4. Channel types
	// ─────────────────────────────────────────────
	_ = make(chan int)   // bidirectional
	_ = make(chan<- int) // send-only
	_ = make(<-chan int) // receive-only
	// Direction is typically restricted in function signatures:
	// func producer(out chan<- int) { out <- 42 }
	// func consumer(in <-chan int) { v := <-in }

	// ─────────────────────────────────────────────
	// 5. Zero value of channel is nil
	// ─────────────────────────────────────────────
	var nilCh chan int
	fmt.Printf("\nnil channel: %v (is nil: %t)\n", nilCh, nilCh == nil)
	// <-nilCh    // blocks forever!
	// nilCh <- 1 // blocks forever!
}
