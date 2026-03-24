//go:build ignore

// Section 6, Topic 50: Multi-Dimensional Slices
//
// Go doesn't have a built-in multi-dimensional array/slice type.
// You create them as slices of slices: [][]int
//
// GOTCHA: Rows can have different lengths (jagged arrays).
// GOTCHA: Each row is independently allocated — not contiguous memory.
//         For performance-critical apps, use a flat slice with manual indexing.
//
// Run: go run examples/s06_multidimensional.go

package main

import "fmt"

func main() {
	fmt.Println("=== Multi-Dimensional Slices ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. 2D slice with literal
	// ─────────────────────────────────────────────
	fmt.Println("-- 2D literal --")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i, row := range matrix {
		fmt.Printf("  Row %d: %v\n", i, row)
	}
	fmt.Printf("  matrix[1][2] = %d\n", matrix[1][2]) // 6

	// ─────────────────────────────────────────────
	// 2. 2D slice with make (dynamic size)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- 2D with make --")
	rows, cols := 3, 4
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	// Fill with values:
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid[i][j] = i*cols + j
		}
	}
	for i, row := range grid {
		fmt.Printf("  Row %d: %v\n", i, row)
	}

	// ─────────────────────────────────────────────
	// 3. Jagged (non-rectangular)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Jagged (different row lengths) --")
	triangle := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},



































}	}		fmt.Printf("  %v\n", row)	for _, row := range board {	board[2][2] = "X"	board[1][1] = "O"	board[0][0] = "X"	var board [3][3]string	fmt.Println("\n-- 2D array (fixed size) --")	// ─────────────────────────────────────────────	// 5. 2D array (fixed size)	// ─────────────────────────────────────────────	// This is cache-friendly and used in performance-critical code.	}		fmt.Printf("  Row %d: %v\n", i, row)		row := flat[i*c : (i+1)*c]	for i := 0; i < r; i++ {	// Access as 2D:	}		}			flat[i*c+j] = i*c + j		for j := 0; j < c; j++ {	for i := 0; i < r; i++ {	flat := make([]int, r*c)	r, c := 3, 4	fmt.Println("\n-- Flat slice (contiguous memory) --")	// ─────────────────────────────────────────────	// 4. Flat slice with manual indexing (performance)	// ─────────────────────────────────────────────	}		fmt.Printf("  %v\n", row)	for _, row := range triangle {	}