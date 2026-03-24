//go:build ignore

// Section 4, Topic 27: Multiple Return Values
//
// Go functions can return multiple values. This is used extensively for:
//   - Returning a result + error (the most common pattern)
//   - Returning a value + ok boolean (map lookup, type assertion)
//   - Returning multiple computed results
//
// GOTCHA: You MUST accept all return values, or explicitly discard with _.
// GOTCHA: Multiple return values are NOT tuples — you can't store them in a
//         single variable or pass them directly to most functions.
//
// Run: go run examples/s04_multiple_returns.go

package main

import (
	"errors"
	"fmt"
	"math"
)

// ─────────────────────────────────────────────
// 1. Basic multiple returns
// ─────────────────────────────────────────────
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// ─────────────────────────────────────────────
// 2. Multiple computed results
// ─────────────────────────────────────────────
func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// ─────────────────────────────────────────────
// 3. Three return values
// ─────────────────────────────────────────────
func parseCoordinate(input string) (float64, float64, error) {
	var lat, lon float64
	_, err := fmt.Sscanf(input, "%f,%f", &lat, &lon)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid coordinate: %w", err)
	}
	return lat, lon, nil
}

func main() {
	fmt.Println("=== Multiple Return Values ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Using result + error pattern
	// ─────────────────────────────────────────────
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ─────────────────────────────────────────────
	// Discarding return values with _
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Discarding values --")
	val, _ := divide(10, 3) // ignore error (not recommended in production!)
	fmt.Printf("Value only: %.4f\n", val)

	// ─────────────────────────────────────────────
	// Multiple computed values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multiple results --")
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	lo, hi := minMax(nums)
	fmt.Printf("nums: %v → min=%d, max=%d\n", nums, lo, hi)

	// ─────────────────────────────────────────────
	// Three return values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Three returns --")
	lat, lon, err := parseCoordinate("37.7749,-122.4194")
	if err == nil {
		fmt.Printf("Lat=%.4f, Lon=%.4f\n", lat, lon)
	}
	_, _, err = parseCoordinate("invalid")
	if err != nil {
		fmt.Println("Parse error:", err)
	}

	// ─────────────────────────────────────────────
	// GOTCHA: Can't directly pass multiple returns to multi-arg functions
	// ─────────────────────────────────────────────
	// This works (special case — single multi-return as sole argument):
	fmt.Println(divide(10, 3)) // OK: Println accepts ...interface{}

	// But this doesn't work in general:
	// math.Pow(divide(10, 3))  // ERROR: multiple-value in single-value context
	// Fix:
	v, _ := divide(10, 3)
	fmt.Printf("Pow: %.2f\n", math.Pow(v, 2))

	// ─────────────────────────────────────────────
	// GOTCHA: Must handle ALL return values
	// ─────────────────────────────────────────────
	// divide(10, 3)  // ERROR: divide(10, 3) (value of type (float64, error)) used as value
	// You must either assign both or use _:
	_, _ = divide(10, 3) // explicitly discard both

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust
	// ─────────────────────────────────────────────
	// Go:   return val, err  (multiple returns, error as value)
	// Rust: return Ok(val)   (Result<T, E> enum, single return)
	// Go:   more verbose but explicit error handling
	// Rust: pattern matching with match/if let/?
}
