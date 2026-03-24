//go:build ignore

// Section 6, Topic 47: nil slice vs empty slice vs zero-length slice
//
// GOTCHA: nil slice and empty slice behave the same for len, cap, append, range
//         but they are NOT the same in JSON marshaling.
// GOTCHA: var s []int is nil, []int{} and make([]int, 0) are non-nil empty.
//
// Run: go run examples/s06_nil_vs_empty_slice.go

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("=== nil vs Empty Slice ===")
	fmt.Println()

	// 1. Three ways to get a "zero-length" slice
	var nilSlice []int          // nil
	emptyLiteral := []int{}     // non-nil, len=0
	emptyMake := make([]int, 0) // non-nil, len=0

	fmt.Printf("nilSlice:     len=%d, cap=%d, nil=%t\n",
		len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("emptyLiteral: len=%d, cap=%d, nil=%t\n",
		len(emptyLiteral), cap(emptyLiteral), emptyLiteral == nil)
	fmt.Printf("emptyMake:    len=%d, cap=%d, nil=%t\n",
		len(emptyMake), cap(emptyMake), emptyMake == nil)

	// 2. All behave the same for operations
	fmt.Println("\n-- Same behavior --")
	nilSlice = append(nilSlice, 1)
	emptyLiteral = append(emptyLiteral, 1)
	emptyMake = append(emptyMake, 1)
	fmt.Println("After append(1):", nilSlice, emptyLiteral, emptyMake)

	// 3. JSON difference!
	fmt.Println("\n-- JSON marshaling (IMPORTANT) --")
	type Response struct {
		Items []int `json:"items"`
	}

	r1 := Response{Items: nil}
	r2 := Response{Items: []int{}}

	j1, _ := json.Marshal(r1)
	j2, _ := json.Marshal(r2)
	fmt.Printf("nil slice JSON:   %s\n", j1) // {"items":null}
	fmt.Printf("empty slice JSON: %s\n", j2)  // {"items":[]}

	// 4. Comparison
	fmt.Println("\n-- Comparison --")
	var s1 []int
	fmt.Println("s1 == nil:", s1 == nil) // true
	// s1 == s2 // ERROR: slices can only be compared to nil
}
