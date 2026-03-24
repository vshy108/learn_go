//go:build ignore

// Section 16, Topic 117: Generic Functions (Go 1.18+)
//
// Generics let you write functions and types parameterized by type.
// Syntax:
//   func Name[T constraint](param T) T { ... }
//
// Type parameters appear in square brackets before the function params.
// Constraints define what operations are allowed on T.
//
// GOTCHA: Generics were added in Go 1.18 — not available in older versions.
// GOTCHA: Use generics when the logic is the same for multiple types.
//         Don't overuse them — Go favors simplicity.
//
// Run: go run examples/s16_generic_functions.go

package main

import (
	"fmt"
	"strings"
)

// ─────────────────────────────────────────────
// 1. Basic generic function
// ─────────────────────────────────────────────
func Max[T int | float64 | string](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// ─────────────────────────────────────────────
// 2. any constraint (least restrictive)
// ─────────────────────────────────────────────
func Print[T any](val T) {
	fmt.Printf("  %v (%T)\n", val, val)
}

// ─────────────────────────────────────────────
// 3. comparable constraint (supports == and !=)
// ─────────────────────────────────────────────
func Contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// ─────────────────────────────────────────────
// 4. Multiple type parameters
// ─────────────────────────────────────────────
func Map[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func main() {
	fmt.Println("=== Generic Functions ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Using Max
	// ─────────────────────────────────────────────
	fmt.Println("Max(3, 5):", Max(3, 5))
	fmt.Println("Max(3.14, 2.72):", Max(3.14, 2.72))
	fmt.Println("Max(\"apple\", \"banana\"):", Max("apple", "banana"))

	// Explicit type argument (usually inferred, so not needed here):
	fmt.Println("Max(10, 20):", Max(10, 20))

	// ─────────────────────────────────────────────
	// Using Print
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Print --")
	Print(42)
	Print("hello")
	Print(3.14)
	Print(true)

	// ─────────────────────────────────────────────
	// Using Contains
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Contains --")
	fmt.Println("Contains [1,2,3] 2:", Contains([]int{1, 2, 3}, 2))
	fmt.Println("Contains [\"a\",\"b\"] \"c\":", Contains([]string{"a", "b"}, "c"))

	// ─────────────────────────────────────────────
	// Using Map
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Map --")
	nums := []int{1, 2, 3, 4, 5}
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("Doubled:", doubled)

	words := []string{"hello", "world"}
	upper := Map(words, strings.ToUpper)
	fmt.Println("Upper:", upper)

	strs := Map(nums, func(n int) string {
		return fmt.Sprintf("item-%d", n)
	})
	fmt.Println("Strings:", strs)
}
