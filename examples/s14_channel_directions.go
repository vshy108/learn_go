//go:build ignore

// Section 14, Topic 105: Channel Directions
//
// Channel direction restricts how a channel can be used in a function.
//   chan<- T    — send-only channel
//   <-chan T    — receive-only channel
//   chan T      — bidirectional
//
// Bidirectional channels are implicitly converted to directional.
// This provides compile-time safety.
//
// GOTCHA: You can't send on a receive-only channel or vice versa.
// GOTCHA: You can't close a receive-only channel.
//
// Run: go run examples/s14_channel_directions.go

package main

import "fmt"

func main() {
	fmt.Println("=== Channel Directions ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Producer (send-only) and Consumer (receive-only)
	// ─────────────────────────────────────────────
	ch := make(chan int, 5)
	go produce(ch)
	consume(ch)

	// ─────────────────────────────────────────────
	// 2. Pipeline pattern
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pipeline --")
	naturals := make(chan int)
	squares := make(chan int)

	go generateNumbers(naturals)
	go squareNumbers(naturals, squares)

	// Print first 5 squares:
	for i := 0; i < 5; i++ {
		fmt.Printf("  %d\n", <-squares)
	}

	// ─────────────────────────────────────────────
	// 3. Direction conversion
	// ─────────────────────────────────────────────
	// Bidirectional → directional: implicit (always allowed)
	// Directional → bidirectional: NOT allowed (compile error)
	ch2 := make(chan string)
	var sendOnly chan<- string = ch2 // ok
	var recvOnly <-chan string = ch2 // ok
	_ = sendOnly
	_ = recvOnly
	// var bidir chan string = sendOnly  // ERROR: cannot convert
}

// produce can only send to this channel
func produce(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out) // producer closes the channel
}

// consume can only receive from this channel
func consume(in <-chan int) {
	for v := range in {
		fmt.Printf("  Received: %d\n", v)
	}
}

func generateNumbers(out chan<- int) {
	for i := 1; ; i++ {
		out <- i
	}
}

func squareNumbers(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
}
