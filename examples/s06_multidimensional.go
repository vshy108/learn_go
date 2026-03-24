//go:build ignore

// Section 6, Topic 50: Multi-dimensional Slices
//
// Go has no built-in 2D slice. Use slice of slices.
// Each inner slice can have different lengths (jagged).
//
// Run: go run examples/s06_multidimensional.go

package main

import "fmt"

func main() {
	fmt.Println("=== Multi-dimensional Slices ===")
	fmt.Println()

	// 1. 2D slice literal
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("-- 3x3 matrix --")
	for _, row := range matrix {
		fmt.Println(" ", row)
	}
	fmt.Printf("matrix[1][2] = %d\n", matrix[1][2])

	// 2. Dynamic allocation with make
	fmt.Println("\n-- Dynamic 2D --")
	rows, cols := 3, 4
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = i*cols + j
		}
	}
	for _, row := range grid {
		fmt.Println(" ", row)
	}

	// 3. Jagged (irregular) slices
	fmt.Println("\n-- Jagged --")
	jagged := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
	}
	for i, row := range jagged {
		fmt.Printf("  row %d (len %d): %v\n", i, len(row), row)
	}

	// 4. 2D array (fixed size)
	fmt.Println("\n-- 2D array (fixed) --")
	var arr [2][3]int
	arr[0] = [3]int{1, 2, 3}
	arr[1] = [3]int{4, 5, 6}
	fmt.Println(" ", arr)
}
