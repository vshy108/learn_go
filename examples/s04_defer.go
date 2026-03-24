//go:build ignore

// Section 4, Topic 32: defer — Deferred Function Calls
//
// `defer` schedules a function call to execute when the surrounding function
// returns. Deferred calls are pushed onto a stack and executed in LIFO order.
//
// Primary use cases:
//   - Cleanup: closing files, database connections, mutexes
//   - Matching: open/close, lock/unlock, begin/end
//
// GOTCHA: Deferred function ARGUMENTS are evaluated immediately, not when
//         the deferred function runs.
// GOTCHA: defer runs when the FUNCTION returns, not when the BLOCK ends.
// GOTCHA: Deferred functions can read and modify named return values.
// GOTCHA: defer in a loop can leak resources if the loop runs many times.
//
// Run: go run examples/s04_defer.go

package main

import "fmt"







































































































}	fmt.Println(result)	result := a / b // panics if b == 0	fmt.Printf("  %d / %d = ", a, b)	}()		}			fmt.Printf("  Recovered from panic: %v\n", r)		if r := recover(); r != nil {	defer func() {func safeDivide(a, b int) {}	fmt.Println("  (defer would close file here)")	fmt.Println("  Processing...")	// defer f.Close()  // guaranteed to run when function returns	// if err != nil { return }	// f, err := os.Open("file.txt")	fmt.Println("  Opening file...")func processFile() {}	return x * 2 // sets result = 10, then defer adds 10 → returns 20	}()		result += 10 // modifies the named return value!	defer func() {func doubleAndAdd(x int) (result int) {// Named return + defer = can modify the return value}	fmt.Println("  doubleAndAdd(5) =", doubleAndAdd(5))	fmt.Println("\n-- defer modifying named returns --")	// ─────────────────────────────────────────────	// 4. defer with named returns	// ─────────────────────────────────────────────	// The closure will print y=200 because it captures y by reference	y = 200	}()		fmt.Printf("  Closure y=%d (captures variable by reference)\n", y)	defer func() {	y := 100	// To capture the CURRENT value at defer time, use a closure:	// Output: Current x=20, then Deferred x=10	fmt.Printf("  Current x=%d\n", x)	x = 20	defer fmt.Printf("  Deferred x=%d (captured at defer time)\n", x)	x := 10	fmt.Println("\n-- Arguments evaluated immediately --")	// ─────────────────────────────────────────────	// 3. GOTCHA: Arguments evaluated immediately	// ─────────────────────────────────────────────	defer fmt.Println("  Third deferred (runs first)")	defer fmt.Println("  Second deferred (runs second)")	defer fmt.Println("  First deferred (runs last)")	fmt.Println("-- LIFO order --")	// ─────────────────────────────────────────────	// 2. LIFO order — last defer runs first	// ─────────────────────────────────────────────func deferDemo() {}	safeDivide(10, 0) // would panic without recover	fmt.Println("\n-- defer with recover --")	// ─────────────────────────────────────────────	// 7. defer with panic/recover	// ─────────────────────────────────────────────	// FIX: wrap in a helper function or use explicit close	// }	//     defer f.Close()  // 10000 files stay open until function returns!	//     f, _ := os.Open(files[i])	// for i := 0; i < 10000; i++ {	// BAD: defers accumulate until function returns	fmt.Println("\n-- defer in loops (be careful!) --")	// ─────────────────────────────────────────────	// 6. GOTCHA: defer in loops	// ─────────────────────────────────────────────	processFile()	fmt.Println("\n-- Practical cleanup pattern --")	// ─────────────────────────────────────────────	// 5. Practical: defer for cleanup	// ─────────────────────────────────────────────	deferDemo()	fmt.Println()	// Output: Start, End, Deferred (runs last)	fmt.Println("End")	defer fmt.Println("Deferred (runs last)")	fmt.Println("Start")	fmt.Println("-- Basic defer --")	// ─────────────────────────────────────────────	// 1. Basic defer — runs after function returns	// ─────────────────────────────────────────────	fmt.Println()	fmt.Println("=== defer Keyword ===")func main() {