//go:build ignore

// Section 14, Topic 108: Pipeline Pattern
//
// Pipelines chain stages connected by channels.
// Each stage: receives from input channel → processes → sends to output channel.
//
// This is a core Go concurrency pattern.
//
// Run: go run examples/s14_pipeline.go

package main

import "fmt"

func main() {
	fmt.Println("=== Pipeline Pattern ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Simple pipeline: generate → square → print
	// ─────────────────────────────────────────────
	fmt.Println("-- generate → square → print --")

	// Stage 1: Generate numbers
	nums := generate(2, 3, 4, 5)

	// Stage 2: Square each number
	squared := square(nums)

	// Stage 3: Print results
	for v := range squared {
		fmt.Printf("  %d\n", v) // 4, 9, 16, 25
	}

	// ─────────────────────────────────────────────
	// 2. Longer pipeline: generate → double → addOne → print
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Longer pipeline --")
	stage1 := generate(1, 2, 3, 4, 5)
	stage2 := mapChan(stage1, func(x int) int { return x * 2 })
	stage3 := mapChan(stage2, func(x int) int { return x + 1 })

	for v := range stage3 {
		fmt.Printf("  %d\n", v) // 3, 5, 7, 9, 11
	}

	// ─────────────────────────────────────────────
	// 3. Filter pipeline
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Filter pipeline --")
	all := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	evens := filterChan(all, func(x int) bool { return x%2 == 0 })
	doubled := mapChan(evens, func(x int) int { return x * 2 })

	for v := range doubled {
		fmt.Printf("  %d\n", v) // 4, 8, 12, 16, 20
	}

	// ─────────────────────────────────────────────
	// Pipeline benefits:
	// ─────────────────────────────────────────────
	// - Each stage is independent and testable
	// - Stages can run concurrently on different cores
	// - Memory-efficient: processes one item at a time
	// - Composable: mix and match stages
}

// generate sends values to a channel and closes it
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square reads from in, squares each value, sends to out
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

// mapChan applies a function to each value
func mapChan(in <-chan int, f func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- f(v)
		}
		close(out)
	}()
	return out
}

// filterChan keeps values where predicate is true
func filterChan(in <-chan int, pred func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if pred(v) {
				out <- v
			}
		}
		close(out)
	}()
	return out
}
