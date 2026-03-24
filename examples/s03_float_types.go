//go:build ignore

// Section 3, Topic 18: Float Types — float32 and float64
//
// Go has two floating-point types, both following IEEE 754:
//   - float32: single precision, ~7 significant digits, 4 bytes
//   - float64: double precision, ~15 significant digits, 8 bytes
//
// Go defaults to float64 for floating-point literals (same as Rust).
//
// GOTCHA: float32 loses precision quickly — use float64 unless memory is critical.
// GOTCHA: Floating-point comparison with == is unreliable.
// GOTCHA: NaN != NaN (IEEE 754 spec). Use math.IsNaN() to check.
//
// Run: go run examples/s03_float_types.go

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Float Types ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. float32 vs float64
	// ─────────────────────────────────────────────
	var f32 float32 = 3.14159265358979
	var f64 float64 = 3.14159265358979
	fmt.Println("-- Precision comparison --")
	fmt.Printf("float32: %.20f (only ~7 digits accurate)\n", f32)
	fmt.Printf("float64: %.20f (about ~15 digits accurate)\n", f64)

	// ─────────────────────────────────────────────
	// 2. Ranges and special values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Ranges --")
	fmt.Printf("float32 max:    %e\n", math.MaxFloat32)
	fmt.Printf("float32 min:    %e (smallest positive)\n", math.SmallestNonzeroFloat32)
	fmt.Printf("float64 max:    %e\n", math.MaxFloat64)
	fmt.Printf("float64 min:    %e (smallest positive)\n", math.SmallestNonzeroFloat64)

	// ─────────────────────────────────────────────
	// 3. Special IEEE 754 values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Special values --")
	inf := math.Inf(1)
	negInf := math.Inf(-1)
	nan := math.NaN()

	fmt.Printf("+Inf:  %f\n", inf)
	fmt.Printf("-Inf:  %f\n", negInf)
	fmt.Printf("NaN:   %f\n", nan)

	// Infinity arithmetic:
	fmt.Printf("1/+Inf = %f\n", 1.0/inf)     // 0
	fmt.Printf("+Inf + 1 = %f\n", inf+1)     // +Inf
	fmt.Printf("+Inf - Inf = %f\n", inf-inf) // NaN

	// ─────────────────────────────────────────────
	// 4. GOTCHA: NaN comparisons
	// ─────────────────────────────────────────────
	fmt.Println("\n-- NaN gotchas --")
	fmt.Printf("NaN == NaN:  %t (always false!)\n", nan == nan)
	fmt.Printf("NaN != NaN:  %t (always true!)\n", nan != nan)
	fmt.Printf("NaN > 0:     %t\n", nan > 0)
	fmt.Printf("NaN < 0:     %t\n", nan < 0)
	fmt.Printf("math.IsNaN:  %t (correct way to check)\n", math.IsNaN(nan))

	// ─────────────────────────────────────────────
	// 5. GOTCHA: Floating-point equality
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Equality gotchas --")
	a := 0.1 + 0.2
	fmt.Printf("0.1 + 0.2 = %.20f\n", a)
	fmt.Printf("0.1 + 0.2 == 0.3? %t (usually false!)\n", a == 0.3)

	// Better approach: compare with epsilon
	epsilon := 1e-9
	fmt.Printf("Within epsilon? %t\n", math.Abs(a-0.3) < epsilon)

	// ─────────────────────────────────────────────
	// 6. GOTCHA: Integer division vs float division
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Division --")
	fmt.Printf("7 / 2 = %d (integer division, truncates)\n", 7/2)
	fmt.Printf("7.0 / 2.0 = %f (float division)\n", 7.0/2.0)
	fmt.Printf("float64(7) / 2 = %f\n", float64(7)/2)

	// Division by zero:
	// fmt.Println(1 / 0)      // compile error: division by zero
	var zero float64
	fmt.Printf("1.0 / 0.0 = %f (+Inf)\n", 1.0/zero)
	fmt.Printf("0.0 / 0.0 = %f (NaN)\n", zero/zero)

	// ─────────────────────────────────────────────
	// 7. Common math functions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Math functions --")
	fmt.Printf("math.Sqrt(2)  = %f\n", math.Sqrt(2))
	fmt.Printf("math.Pow(2,10)= %f\n", math.Pow(2, 10))
	fmt.Printf("math.Ceil(2.3)= %f\n", math.Ceil(2.3))
	fmt.Printf("math.Floor(2.7)=%f\n", math.Floor(2.7))
	fmt.Printf("math.Round(2.5)=%f\n", math.Round(2.5))
	fmt.Printf("math.Abs(-5)  = %f\n", math.Abs(-5))
	fmt.Printf("math.Pi       = %f\n", math.Pi)
	fmt.Printf("math.E        = %f\n", math.E)
}
