//go:build ignore

// Section 3, Topic 23: Integer Overflow Behavior
//
// Go integer overflow wraps around silently (no panic, no error).
// This is fundamentally different from Rust, which panics on overflow
// in debug mode.
//
// For constant expressions, overflow IS caught at compile time.
// For runtime expressions, overflow wraps using two's complement.
//
// GOTCHA: There are NO overflow-checking arithmetic operations built-in.
//         You must implement overflow checks yourself or use math/bits.
// GOTCHA: Unsigned overflow is well-defined (wraps). Signed overflow
//         is also well-defined in Go (wraps), unlike C (undefined behavior).
//
// Run: go run examples/s03_overflow.go

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Integer Overflow ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Compile-time overflow detection
	// ─────────────────────────────────────────────
	// Constants are checked at compile time:
	// const x int8 = 128  // ERROR: constant 128 overflows int8
	// const y uint8 = -1  // ERROR: constant -1 overflows uint8

	// ─────────────────────────────────────────────
	// 2. Runtime overflow — wraps silently
	// ─────────────────────────────────────────────
	fmt.Println("-- Unsigned overflow (wrapping) --")
	var u uint8 = 255
	fmt.Printf("uint8: %d + 1 = %d (wraps to 0)\n", u, u+1)

	var u2 uint8 = 0
	fmt.Printf("uint8: %d - 1 = %d (wraps to 255)\n", u2, u2-1)

	fmt.Println("\n-- Signed overflow (wrapping) --")
	var s int8 = 127
	fmt.Printf("int8: %d + 1 = %d (wraps to -128)\n", s, s+1)

	var s2 int8 = -128
	fmt.Printf("int8: %d - 1 = %d (wraps to 127)\n", s2, s2-1)

	// ─────────────────────────────────────────────
	// 3. Multiplication overflow
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multiplication overflow --")
	var a uint8 = 200
	var b uint8 = 2
	fmt.Printf("uint8: %d * %d = %d (expected 400, got %d)\n", a, b, a*b, a*b)
	// 200 * 2 = 400, but 400 % 256 = 144

	// ─────────────────────────────────────────────
	// 4. Checking for overflow manually
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Manual overflow check --")
	x, y := int32(math.MaxInt32), int32(1)
	if willOverflowAdd(x, y) {
		fmt.Printf("int32: %d + %d would overflow!\n", x, y)
	}

	// ─────────────────────────────────────────────
	// 5. Using math constants for bounds checking
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Bounds checking with math constants --")
	val := 50000
	if val > math.MaxInt16 || val < math.MinInt16 {
		fmt.Printf("%d does not fit in int16 (range: %d to %d)\n",
			val, math.MinInt16, math.MaxInt16)
	}

	// ─────────────────────────────────────────────
	// 6. Float overflow
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Float overflow --")
	var f float64 = math.MaxFloat64
	fmt.Printf("MaxFloat64: %e\n", f)
	fmt.Printf("MaxFloat64 * 2: %f (becomes +Inf)\n", f*2) // +Inf, no panic

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust
	// ─────────────────────────────────────────────
	// Go:   wraps silently at runtime (always)
	// Rust: panics in debug, wraps in release mode
	// Rust: has checked_add(), saturating_add(), wrapping_add()
	// Go:   no built-in overflow-checked arithmetic
	// C:    signed overflow is undefined behavior!
}

func willOverflowAdd(a, b int32) bool {
	if b > 0 && a > math.MaxInt32-b {
		return true
	}
	if b < 0 && a < math.MinInt32-b {
		return true
	}
	return false
}
