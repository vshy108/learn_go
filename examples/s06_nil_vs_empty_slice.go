//go:build ignore

// Section 6, Topic 47: nil Slice vs Empty Slice vs Zero-Length Slice
//
// Three similar but subtly different constructs:
//   var s []int          → nil slice (nil pointer, len=0, cap=0)
//   s := []int{}         → empty slice (non-nil pointer, len=0, cap=0)
//   s := make([]int, 0)  → empty slice (non-nil pointer, len=0, cap=0)
//
// In MOST cases they behave identically (append, len, range all work).
// The difference matters for JSON serialization and nil checks.
//
// GOTCHA: nil slice marshals to JSON "null", empty slice to "[]".
// GOTCHA: reflect.DeepEqual(nil, []int{}) is false!
//
// Run: go run examples/s06_nil_vs_empty_slice.go

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("=== nil vs Empty Slice ===")
	fmt.Println()













































































}	// - Use `make([]int, 0, cap)` when you know approximate capacity	//   (e.g., for JSON APIs returning [])	// - Use `s := []int{}` or `make([]int, 0)` when you need a non-nil empty slice	// - Use `var s []int` (nil) when "no data" is the natural zero state	// ─────────────────────────────────────────────	// Best practice	// ─────────────────────────────────────────────	fmt.Printf("DeepEqual([]int{}, []{}):%t\n", reflect.DeepEqual([]int{}, []int{}))	fmt.Printf("DeepEqual(nil, []int{}): %t\n", reflect.DeepEqual([]int(nil), []int{}))	fmt.Printf("DeepEqual(nil, nil):     %t\n", reflect.DeepEqual([]int(nil), []int(nil)))	fmt.Println("\n-- DeepEqual --")	// ─────────────────────────────────────────────	// 6. GOTCHA: reflect.DeepEqual	// ─────────────────────────────────────────────	// API consumers often prefer [] over null, so initialize to empty!	fmt.Printf("empty: %s\n", j2) // {"items":[]}	j2, _ := json.Marshal(r2)	r2 := Response{Items: []int{}}	// empty slice → []	fmt.Printf("nil:   %s\n", j1) // {"items":null}	j1, _ := json.Marshal(r1)	r1 := Response{Items: nil}	// nil slice → null	}		Items []int `json:"items"`	type Response struct {	fmt.Println("\n-- JSON serialization --")	// ─────────────────────────────────────────────	// 5. GOTCHA: JSON serialization difference	// ─────────────────────────────────────────────	fmt.Println("\nRange over nil: no panic")	}		fmt.Println("This never runs")	for range s {	var s []int	// ─────────────────────────────────────────────	// 4. range works on all (including nil)	// ─────────────────────────────────────────────	fmt.Printf("emptyMake: %v\n", emptyMake)	fmt.Printf("emptyLit:  %v\n", emptyLit)	fmt.Printf("nilSlice:  %v\n", nilSlice)	emptyMake = append(emptyMake, 1, 2, 3)	emptyLit = append(emptyLit, 1, 2, 3)	nilSlice = append(nilSlice, 1, 2, 3)	fmt.Println("\n-- append --")	// ─────────────────────────────────────────────	// 3. append works on all	// ─────────────────────────────────────────────	fmt.Printf("emptyMake: len=%d, cap=%d\n", len(emptyMake), cap(emptyMake))	fmt.Printf("emptyLit:  len=%d, cap=%d\n", len(emptyLit), cap(emptyLit))	fmt.Printf("nilSlice:  len=%d, cap=%d\n", len(nilSlice), cap(nilSlice))	fmt.Println("\n-- len and cap --")	// ─────────────────────────────────────────────	// 2. len and cap (all the same)	// ─────────────────────────────────────────────	fmt.Printf("emptyMake == nil: %t\n", emptyMake == nil) // false	fmt.Printf("emptyLit == nil:  %t\n", emptyLit == nil)  // false	fmt.Printf("nilSlice == nil:  %t\n", nilSlice == nil)  // true	fmt.Println("-- nil check --")	// ─────────────────────────────────────────────	// 1. nil check	// ─────────────────────────────────────────────	emptyMake := make([]int, 0)	emptyLit := []int{}	var nilSlice []int