//go:build ignore

// Section 5, Topic 41: Infinite Loops and Breaking Out
//
// Go's infinite loop is: for { }
// This is the idiomatic way to write servers, event loops, retry loops, etc.
//
// Run: go run examples/s05_infinite_loop.go

package main

import "fmt"

func main() {
	fmt.Println("=== Infinite Loops ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic infinite loop with break
	// ─────────────────────────────────────────────
	fmt.Println("-- Counting to 5 --")
	i := 0
	for {
		if i >= 5 {
			break
		}




















































}	// Go's for cannot return a value.	// Rust's loop can return a value: let x = loop { break 42; };	// Both support labeled breaks for nested loops.	// Rust: loop { ... break }	// Go:   for { ... break }	// ─────────────────────────────────────────────	// Comparison: Go vs Rust	// ─────────────────────────────────────────────	// }	//     go handleConnection(conn)	//     if err != nil { continue }	//     conn, err := listener.Accept()	// for {	// ─────────────────────────────────────────────	// 4. Typical server pattern (conceptual)	// ─────────────────────────────────────────────	}		fmt.Printf("Attempt %d failed, retrying...\n", attempts)		}			break			fmt.Printf("Succeeded after %d attempts\n", attempts)		if success {		success := attempts >= 3 // simulate: succeeds on 3rd try		attempts++	for {	attempts := 0	fmt.Println("\n-- Retry pattern --")	// ─────────────────────────────────────────────	// 3. Retry pattern	// ─────────────────────────────────────────────	fmt.Println()	}		}			break		if n > 20 {		n *= 2		fmt.Printf("n=%d ", n)	for {	n := 1	fmt.Println("\n-- Do-while pattern --")	// ─────────────────────────────────────────────	// 2. Do-while pattern (execute at least once)	// ─────────────────────────────────────────────	fmt.Println()	}		i++		fmt.Printf("%d ", i)