//go:build ignore

// Section 2, Topic 9: Zero Values for All Types
//
// Every type in Go has a zero value - the default when not initialized.
// There is NO undefined/null for most types.
//
// GOTCHA: Zero value of a pointer is nil (the only "null" in Go).
// GOTCHA: Zero value of a slice/map/channel is nil, but they behave differently.
//
// Run: go run examples/s02_zero_values.go

package main

import "fmt"

func main() {
	fmt.Println("=== Zero Values ===")
	fmt.Println()

	// Numeric types
	var i int
	var i8 int8
	var i64 int64
	var u uint
	var f32 float32
	var f64 float64
	var c64 complex64
	fmt.Println("-- Numeric --")
	fmt.Printf("int: %d, int8: %d, int64: %d, uint: %d\n", i, i8, i64, u)
	fmt.Printf("float32: %f, float64: %f, complex64: %v\n", f32, f64, c64)

	// String and bool
	var s string
	var b bool
	fmt.Println("\n-- String & Bool --")
	fmt.Printf("string: %q (empty, not nil), bool: %t\n", s, b)

	// Pointer, slice, map, channel, function, interface
	var ptr *int
	var sl []int
	var m map[string]int
	var ch chan int
	var fn func()
	var iface interface{}
	fmt.Println("\n-- Reference types (all nil) --")
	fmt.Printf("*int: %v, []int: %v, map: %v, chan: %v, func: %v, interface: %v\n",
		ptr, sl, m, ch, fn, iface)
	fmt.Printf("slice==nil: %t, map==nil: %t\n", sl == nil, m == nil)

	// Struct zero value
	type Point struct {
		X, Y int
	}
	var p Point
	fmt.Println("\n-- Struct --")
	fmt.Printf("Point: %+v (fields are zero-valued)\n", p)

	// Array zero value
	var arr [3]int
	fmt.Println("\n-- Array --")
	fmt.Printf("[3]int: %v\n", arr)
}
