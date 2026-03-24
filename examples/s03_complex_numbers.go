//go:build ignore

// Section 3, Topic 19: Complex Numbers — complex64 and complex128
//
// Go has built-in complex number support (unusual for a systems language!).
//   - complex64:  float32 real + float32 imaginary
//   - complex128: float64 real + float64 imaginary (default)
//
// Built-in functions: complex(), real(), imag()
// Math support: cmplx package for complex math operations.
//
// GOTCHA: The default complex literal type is complex128 (not complex64).
// GOTCHA: You cannot mix complex64 and complex128 without conversion.
//
// Run: go run examples/s03_complex_numbers.go

package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println("=== Complex Numbers ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Creating complex numbers
	// ─────────────────────────────────────────────
	c1 := 3 + 4i              // complex128 (literal)
	c2 := complex(3.0, 4.0)   // complex128 (built-in function)
	var c3 complex64 = 3 + 4i // explicit complex64
	fmt.Printf("c1 = %v  (type: %T)\n", c1, c1)
	fmt.Printf("c2 = %v  (type: %T)\n", c2, c2)
	fmt.Printf("c3 = %v  (type: %T)\n", c3, c3)

	// ─────────────────────────────────────────────
	// 2. Extracting real and imaginary parts
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Real and imaginary parts --")
	fmt.Printf("real(c1) = %f\n", real(c1))
	fmt.Printf("imag(c1) = %f\n", imag(c1))

	// ─────────────────────────────────────────────
	// 3. Arithmetic
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Arithmetic --")
	a := 1 + 2i
	b := 3 + 4i
	fmt.Printf("(%v) + (%v) = %v\n", a, b, a+b)
	fmt.Printf("(%v) - (%v) = %v\n", a, b, a-b)
	fmt.Printf("(%v) * (%v) = %v\n", a, b, a*b)
	fmt.Printf("(%v) / (%v) = %v\n", a, b, a/b)

	// ─────────────────────────────────────────────
	// 4. cmplx package functions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- cmplx package --")
	fmt.Printf("Abs(3+4i)   = %f (magnitude)\n", cmplx.Abs(3+4i))
	fmt.Printf("Phase(1+1i) = %f (angle in radians)\n", cmplx.Phase(1+1i))
	fmt.Printf("Conj(3+4i)  = %v (conjugate)\n", cmplx.Conj(3+4i))
	fmt.Printf("Sqrt(-1)    = %v\n", cmplx.Sqrt(-1)) // (0+1i)
	fmt.Printf("Exp(πi)     = %v (Euler's identity ≈ -1)\n", cmplx.Exp(1i*cmplx.Inf()))

	// Euler's identity: e^(iπ) + 1 = 0
	euler := cmplx.Exp(1i*complex(0, 0)+complex(0, 1)*complex(3.14159265358979, 0)) + 1
	fmt.Printf("e^(iπ)+1    ≈ %v (should be ≈ 0)\n", euler)

	// ─────────────────────────────────────────────
	// 5. Zero value
	// ─────────────────────────────────────────────
	var zero complex128
	fmt.Printf("\nZero value: %v\n", zero) // (0+0i)

	// ─────────────────────────────────────────────
	// Comparison with other languages
	// ─────────────────────────────────────────────
	// Go:     3 + 4i (built-in type)
	// Python: 3 + 4j (built-in complex)
	// Rust:   num::complex::Complex::new(3.0, 4.0) (external crate)
	// C:      _Complex double c = 3.0 + 4.0*I (C99+)
}
