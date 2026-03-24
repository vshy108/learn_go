//go:build ignore

// Section 3, Topic 21: Explicit Type Conversion
//
// Go has NO implicit type conversion (unlike C, Java, Python).
// You must always convert explicitly using T(value) syntax.
//
// This applies everywhere:
//   - Between different integer types (int32 → int64)
//   - Between int and float (int → float64)
//   - Between int and uint (can change sign!)
//   - Between string and []byte/[]rune
//
// GOTCHA: Converting a larger int to a smaller one silently truncates.
// GOTCHA: float → int truncates toward zero (no rounding).
// GOTCHA: Negative int → uint produces a large positive number.
//
// Run: go run examples/s03_type_conversion.go

package main

import "fmt"

func main() {
	fmt.Println("=== Explicit Type Conversion ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Between integer sizes
	// ─────────────────────────────────────────────
	fmt.Println("-- Integer conversions --")
	var a int32 = 42
	var b int64 = int64(a) // explicit: int32 → int64
	var c int8 = int8(a)   // explicit: int32 → int8 (fits)
	fmt.Printf("int32(%d) → int64(%d), int8(%d)\n", a, b, c)

	// GOTCHA: Truncation when value doesn't fit
	var big int32 = 300
	var small int8 = int8(big) // 300 % 256 = 44
	fmt.Printf("int32(300) → int8(%d) — truncated!\n", small)

	var neg int32 = -1
	var u uint32 = uint32(neg) // -1 → 4294967295
	fmt.Printf("int32(-1) → uint32(%d) — two's complement!\n", u)

	// ─────────────────────────────────────────────
	// 2. Float ↔ Int conversion
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Float ↔ Int --")
	f := 3.99
	i := int(f) // truncates toward zero (NOT rounds!)
	fmt.Printf("int(3.99) = %d (truncated, not rounded)\n", i)

	negF := -3.99
	negI := int(negF)
	fmt.Printf("int(-3.99) = %d (truncated toward zero)\n", negI)

	// int → float
	n := 42
	fl := float64(n)
	fmt.Printf("float64(42) = %f\n", fl)

	// GOTCHA: Large int may lose precision as float:
	bigInt := 9007199254740993 // 2^53 + 1
	bigFloat := float64(bigInt)
	fmt.Printf("int %d → float64 %f (precision lost!)\n", bigInt, bigFloat)
	fmt.Printf("Back to int: %d\n", int(bigFloat)) // may differ!

	// ─────────────────────────────────────────────
	// 3. Between float32 and float64
	// ─────────────────────────────────────────────
	fmt.Println("\n-- float32 ↔ float64 --")
	var f64 float64 = 3.14159265358979
	var f32 float32 = float32(f64) // loses precision
	fmt.Printf("float64(%.15f) → float32(%.15f)\n", f64, f32)
	// float64 back:
	var f64b float64 = float64(f32)
	fmt.Printf("float32 back to float64: %.15f (precision already lost)\n", f64b)

	// ─────────────────────────────────────────────
	// 4. String ↔ []byte / []rune
	// ─────────────────────────────────────────────
	fmt.Println("\n-- String conversions --")
	s := "Hello, 世界"
	bytes := []byte(s)
	runes := []rune(s)
	fmt.Printf("string → []byte: %v\n", bytes)
	fmt.Printf("string → []rune: %v\n", runes)
	fmt.Printf("[]byte → string: %s\n", string(bytes))
	fmt.Printf("[]rune → string: %s\n", string(runes))

	// int → string: converts to Unicode character, NOT digit string!
	fmt.Printf("string(65) = %s (rune 'A', NOT \"65\")\n", string(rune(65)))
	// Use strconv.Itoa() or fmt.Sprintf() for number-to-string conversion.

	// ─────────────────────────────────────────────
	// 5. Conversion vs assertion (for interfaces)
	// ─────────────────────────────────────────────
	// Type conversion: T(value) — between compatible concrete types
	// Type assertion:  value.(T) — extract concrete type from interface
	// These are fundamentally different operations!
	var iface interface{} = 42
	val := iface.(int) // type assertion (runtime check)
	fmt.Printf("\ninterface{}(42).(int) = %d\n", val)

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust
	// ─────────────────────────────────────────────
	// Go:   int64(x), float64(x)       — T(value) syntax
	// Rust: x as i64, x as f64         — `as` keyword
	// Rust: i64::from(x), x.into()     — trait-based conversion
	//
	// Rust has TryFrom/TryInto for fallible conversions.
	// Go silently truncates — there is no checked conversion built-in.
}
