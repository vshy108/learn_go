//go:build ignore

// Section 4, Topic 34: Recursion (No Tail-Call Optimization in Go)
//
// Go supports recursion but does NOT have tail-call optimization (TCO).
// Every recursive call adds a new frame to the stack.
// Deep recursion WILL cause a stack overflow (goroutine stack starts small
// but grows dynamically up to a default limit of 1GB).
//
// GOTCHA: No TCO means you should prefer iterative solutions for deep recursion.
// GOTCHA: Goroutine stacks start at ~2-8KB and grow dynamically, but there's a limit.
//
// Run: go run examples/s04_recursion.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Classic factorial (not tail-recursive-safe)
// ─────────────────────────────────────────────
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1) // NOT tail position — multiplication after call
}

// ─────────────────────────────────────────────
// 2. Fibonacci (naive — exponential time)
// ─────────────────────────────────────────────
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)









































































































































}	// Python: no TCO — sys.setrecursionlimit(N) default ~1000	// Rust: no guaranteed TCO — but LLVM may optimize tail calls	// Go:   no TCO — prefer loops for deep recursion	//	// For performance-critical code, prefer iterative solutions.	// But it's slower than iterative due to function call overhead.	// factorial(1000000)  // Will work (goroutine stack grows dynamically)	// ─────────────────────────────────────────────	// GOTCHA: No tail-call optimization	// ─────────────────────────────────────────────	fmt.Println("  Inorder:", inorder(tree)) // [1 2 3 4 5 6 7]	}		},			Right: &TreeNode{Value: 7},			Left:  &TreeNode{Value: 5},			Value: 6,		Right: &TreeNode{		},			Right: &TreeNode{Value: 3},			Left:  &TreeNode{Value: 1},			Value: 2,		Left: &TreeNode{		Value: 4,	tree := &TreeNode{	fmt.Println("\n-- Tree inorder traversal --")	// ─────────────────────────────────────────────	// Tree traversal	// ─────────────────────────────────────────────	fmt.Println("  find 10:", binarySearch(arr, 10, 0, len(arr)-1)) // -1	fmt.Println("  find 7:", binarySearch(arr, 7, 0, len(arr)-1))   // 3	fmt.Println("  arr:", arr)	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}	fmt.Println("\n-- Binary Search --")	// ─────────────────────────────────────────────	// Binary search	// ─────────────────────────────────────────────	fmt.Println("  gcd(100, 75) =", gcd(100, 75))  // 25	fmt.Println("  gcd(48, 18) =", gcd(48, 18))   // 6	fmt.Println("\n-- GCD --")	// ─────────────────────────────────────────────	// GCD	// ─────────────────────────────────────────────	fmt.Println()	}		fmt.Printf("%d ", fibonacciIterative(i))	for i := 0; i < 10; i++ {	fmt.Print("  Iterative: ")	fmt.Println()	}		fmt.Printf("%d ", fibonacci(i))	for i := 0; i < 10; i++ {	fmt.Print("  Recursive: ")	fmt.Println("\n-- Fibonacci --")	// ─────────────────────────────────────────────	// Fibonacci	// ─────────────────────────────────────────────	}		fmt.Printf("  %d! = %d\n", i, factorial(i))	for i := 0; i <= 10; i++ {	fmt.Println("-- Factorial --")	// ─────────────────────────────────────────────	// Factorial	// ─────────────────────────────────────────────	fmt.Println()	fmt.Println("=== Recursion ===")func main() {}	return result	result = append(result, inorder(node.Right)...)	result = append(result, node.Value)	result = append(result, inorder(node.Left)...)	var result []int	}		return nil	if node == nil {func inorder(node *TreeNode) []int {}	Left, Right *TreeNode	Value       inttype TreeNode struct {// ─────────────────────────────────────────────// 6. Tree traversal (where recursion shines)// ─────────────────────────────────────────────}	}		return binarySearch(arr, target, low, mid-1)	default:		return binarySearch(arr, target, mid+1, high)	case arr[mid] < target:		return mid	case arr[mid] == target:	switch {	mid := (low + high) / 2	}		return -1	if low > high {func binarySearch(arr []int, target, low, high int) int {// ─────────────────────────────────────────────// 5. Binary search (recursive)// ─────────────────────────────────────────────}	return gcd(b, a%b) // tail position, but Go won't optimize it	}		return a	if b == 0 {func gcd(a, b int) int {// ─────────────────────────────────────────────// 4. GCD (naturally recursive via Euclidean algorithm)// ─────────────────────────────────────────────}	return b	}		a, b = b, a+b	for i := 2; i <= n; i++ {	a, b := 0, 1	}		return n	if n <= 1 {func fibonacciIterative(n int) int {// ─────────────────────────────────────────────// 3. Fibonacci (iterative — much better)// ─────────────────────────────────────────────}