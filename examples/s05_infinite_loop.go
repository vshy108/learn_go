//go:build ignore

// Section 5, Topic 41: Infinite Loops and Breaking Out
//
// Go's for {} is the infinite loop. No while keyword exists.
// Use break, return, or os.Exit to terminate.
//
// Run: go run examples/s05_infinite_loop.go

package main

import "fmt"

func main() {
	fmt.Println("=== Infinite Loops ===")
	fmt.Println()

	// 1. Basic infinite loop with break
	fmt.Println("-- Break on condition --")
	count := 0
	for {
		count++
		if count > 5 {
			break
		}
		fmt.Printf("  count=%d\n", count)
	}
	fmt.Println("  Broke out at count =", count)

	// 2. Infinite with continue
	fmt.Println("\n-- Skip evens --")
	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Printf("  odd: %d\n", i)
	}

	// 3. Labeled break
	fmt.Println("\n-- Labeled break --")
outer:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == 1 && y == 1 {
				fmt.Println("  Breaking out of both loops")
				break outer
			}
			fmt.Printf("  x=%d, y=%d\n", x, y)
		}
	}

	// 4. Simulate while loop
	fmt.Println("\n-- Simulated while --")
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("  First power of 2 >= 100:", n)

	// 5. Simulate do-while
	fmt.Println("\n-- Simulated do-while --")
	val := 0
	for {
		val++
		fmt.Printf("  val=%d\n", val)
		if val >= 3 {
			break
		}
	}
}
