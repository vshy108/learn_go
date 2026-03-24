//go:build ignore

// Section 3, Topic 16: Basic Types Overview — int, float64, string, bool
//
// Go has a small set of built-in types. Unlike Rust's many integer types that
// are all explicit, Go keeps it simple with sensible defaults:
//   - int:     platform-sized integer (32 or 64 bit)
//   - float64: 64-bit floating point (the default for float literals)
//   - string:  immutable sequence of bytes (UTF-8 by convention)
//   - bool:    true or false
//
// GOTCHA: Go has NO implicit type conversions — even int32 to int64 requires explicit cast.
// GOTCHA: int is NOT the same as int64 on 64-bit platforms — they're distinct types.
//
// Run: go run examples/s03_basic_types.go

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println("=== Basic Types Overview ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. int — the default integer type
	// ─────────────────────────────────────────────
	n := 42 // inferred as int
	fmt.Printf("Type: %T, Value: %d, Size: %d bytes\n", n, n, unsafe.Sizeof(n))
	// Size is platform-dependent: 8 bytes on 64-bit, 4 bytes on 32-bit

	// ─────────────────────────────────────────────
	// 2. float64 — the default float type
	// ─────────────────────────────────────────────
	f := 3.14 // inferred as float64
	fmt.Printf("Type: %T, Value: %f, Size: %d bytes\n", f, f, unsafe.Sizeof(f))

	// ─────────────────────────────────────────────
	// 3. string — immutable UTF-8 bytes
	// ─────────────────────────────────────────────
	s := "Hello, 世界"
	fmt.Printf("Type: %T, Value: %s\n", s, s)
	fmt.Printf("  len(s)=%d bytes (NOT characters!)\n", len(s))
	// "Hello, " = 7 bytes, "世界" = 6 bytes (3 bytes each in UTF-8)
	// Total: 13 bytes, but only 9 characters (runes)
	runes := []rune(s)
	fmt.Printf("  rune count=%d characters\n", len(runes))

	// ─────────────────────────────────────────────
	// 4. bool — true or false
	// ─────────────────────────────────────────────
	b := true
	fmt.Printf("Type: %T, Value: %t, Size: %d byte\n", b, b, unsafe.Sizeof(b))
	// bool is 1 byte (not 1 bit) — same as in Rust and C

	// ─────────────────────────────────────────────
	// 5. All built-in types at a glance
	// ─────────────────────────────────────────────
	fmt.Println("\n-- All built-in types --")
	fmt.Println("Integers:  int, int8, int16, int32, int64")
	fmt.Println("Unsigned:  uint, uint8, uint16, uint32, uint64, uintptr")
	fmt.Println("Floats:    float32, float64")
	fmt.Println("Complex:   complex64, complex128")
	fmt.Println("Other:     bool, string, byte (=uint8), rune (=int32)")

	// ─────────────────────────────────────────────
	// 6. reflect.TypeOf vs %T
	// ─────────────────────────────────────────────
	fmt.Println("\n-- reflect.TypeOf --")
	fmt.Println("int:", reflect.TypeOf(n))
	fmt.Println("float64:", reflect.TypeOf(f))
	fmt.Println("string:", reflect.TypeOf(s))
	fmt.Println("bool:", reflect.TypeOf(b))

	// ─────────────────────────────────────────────
	// 7. GOTCHA: No implicit conversions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- No implicit conversions --")
	var i32 int32 = 10
	var i64 int64 = 20
	// sum := i32 + i64     // ERROR: mismatched types int32 and int64
	sum := int64(i32) + i64 // Must explicitly convert
	fmt.Printf("int32(%d) + int64(%d) = int64(%d)\n", i32, i64, sum)

	// Even int and int64 are different types:
	var myInt int = 100
	// var myInt64 int64 = myInt  // ERROR: cannot use myInt (type int) as int64
	var myInt64 int64 = int64(myInt) // explicit conversion required
	fmt.Printf("int(%d) → int64(%d)\n", myInt, myInt64)
}
