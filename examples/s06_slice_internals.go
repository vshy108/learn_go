//go:build ignore

// Section 6, Topic 48: Slice Internals — Header, Shared Backing Array
//
// A slice is a 3-word struct (24 bytes on 64-bit):
//   type slice struct {
//       array unsafe.Pointer  // pointer to backing array
//       len   int             // number of elements
//       cap   int             // capacity of backing array from pointer
//   }
//
// Understanding this internal structure explains:
//   - Why slices are "reference-like" even though passed by value
//   - Why append may or may not affect other slices
//   - Why reslicing doesn't copy data
//
// Run: go run examples/s06_slice_internals.go

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== Slice Internals ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Slice header size
	// ─────────────────────────────────────────────
	fmt.Printf("Size of slice header: %d bytes\n", unsafe.Sizeof([]int{}))
	// 24 bytes on 64-bit: ptr(8) + len(8) + cap(8)






























































}	fmt.Printf("After append: safe=%v, data=%v (data unchanged!)\n", safe, data)	safe = append(safe, 999)	// Now append to `safe` will always allocate a new array:	fmt.Printf("safe=%v (len=%d, cap=%d)\n", safe, len(safe), cap(safe))	safe := data[1:3:3] // elements [1,2], cap restricted to 2	// Limit capacity to prevent callers from seeing/modifying beyond the slice:	data := []int{0, 1, 2, 3, 4, 5}	fmt.Println("\n-- Protection with full slice expression --")	// ─────────────────────────────────────────────	// 5. Protecting against shared array bugs	// ─────────────────────────────────────────────		z, &x[0] == &z[0])	fmt.Printf("z=append(x,4,5,6): %v, same array? %t\n",	z := append(x, 4, 5, 6)	// Append beyond capacity — new backing array:		y, &x[0] == &y[0])	fmt.Printf("y=append(x,4): %v, same array? %t\n",	y := append(x, 4)	// Append within capacity — same backing array:	fmt.Printf("x=%v (len=%d, cap=%d)\n", x, len(x), cap(x))	x[0], x[1], x[2] = 1, 2, 3	x := make([]int, 3, 5)	fmt.Println("\n-- Append and reallocation --")	// ─────────────────────────────────────────────	// 4. When append creates a new backing array	// ─────────────────────────────────────────────	fmt.Printf("b extended: %v (sees original data)\n", b)	b = b[:cap(b)] // extend to full capacity	// b can be extended up to its capacity:	fmt.Printf("a cap=%d, b cap=%d\n", cap(a), cap(b)) // b cap = cap(a) - 1	fmt.Printf("a=%v, b=%v\n", a, b)	b := a[1:3] // [2, 3] — shares backing array	a := []int{1, 2, 3, 4, 5}	fmt.Println("\n-- Reslicing --")	// ─────────────────────────────────────────────	// 3. Reslicing doesn't copy	// ─────────────────────────────────────────────	fmt.Printf("  s2:  %v (also changed!)\n", s2) // s2[0] is now 999	fmt.Printf("  s1:  %v\n", s1)	fmt.Printf("  arr: %v\n", arr)	fmt.Printf("After s1[1]=999:\n")	s1[1] = 999 // modifies arr[3]	// Modifying through one slice affects the other:	fmt.Printf("s2 = arr[3:7]: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))	fmt.Printf("s1 = arr[2:5]: %v (len=%d, cap=%d)\n", s1, len(s1), cap(s1))	fmt.Printf("arr: %v\n", arr)	s2 := arr[3:7]  // [3, 4, 5, 6]	s1 := arr[2:5]  // [2, 3, 4]	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}	fmt.Println("\n-- Shared backing array --")	// ─────────────────────────────────────────────	// 2. Shared backing array	// ─────────────────────────────────────────────