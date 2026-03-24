//go:build ignore

// Section 17, Topic 122: Unit Testing Basics (go test)
//
// Go has built-in testing — no framework needed!
//
// Conventions:
//   - Test files: xxx_test.go (same package)
//   - Test functions: func TestXxx(t *testing.T) { ... }
//   - Run: go test ./...
//
// Test file must NOT have package main — it uses the package being tested.
//
// Commands:
//   go test             — run tests in current package
//   go test ./...       — run all tests recursively
//   go test -v          — verbose output
//   go test -run Regex  — run matching tests only
//   go test -count=1    — disable test caching
//
// GOTCHA: Test files must end with _test.go.
// GOTCHA: Test function names must start with Test (capital T).
// GOTCHA: Tests are NOT in package main.
//
// Run: go run examples/s17_unit_testing.go

package main

import "fmt"

// ─────────────────────────────────────────────
// Functions to test (normally in a separate file)
// ─────────────────────────────────────────────
func add(a, b int) int { return a + b }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=== Unit Testing Basics ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Example test file: math_test.go
	// ─────────────────────────────────────────────
	fmt.Println("-- Example test file --")
	fmt.Print(`
// math_test.go
package mypackage

import "testing"

func TestAdd(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
        t.Errorf("add(2, 3) = %d, want 5", result)
    }
}

func TestDivide(t *testing.T) {
    result, err := divide(10, 2)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if result != 5.0 {
        t.Errorf("divide(10, 2) = %f, want 5.0", result)
    }
}

func TestDivideByZero(t *testing.T) {
    _, err := divide(10, 0)
    if err == nil {
        t.Error("expected error for division by zero")
    }
}
`)

	// ─────────────────────────────────────────────
	// Testing functions
	// ─────────────────────────────────────────────
	fmt.Println("-- t.Error vs t.Fatal --")
	fmt.Println("  t.Error():  logs error, continues test")
	fmt.Println("  t.Errorf(): Error with formatting")
	fmt.Println("  t.Fatal():  logs error, STOPS test immediately")
	fmt.Println("  t.Fatalf(): Fatal with formatting")
	fmt.Println("  t.Log():    logs message (shown with -v)")
	fmt.Println("  t.Skip():   skip this test")
	fmt.Println("  t.Helper(): mark as helper (better error location)")

	// Quick demo:
	fmt.Println("\n-- Demo --")
	fmt.Println("add(2, 3) =", add(2, 3))
	result, _ := divide(10, 3)
	fmt.Printf("divide(10, 3) = %.2f\n", result)
	fmt.Println("isPalindrome(\"racecar\") =", isPalindrome("racecar"))
	fmt.Println("isPalindrome(\"hello\") =", isPalindrome("hello"))
}
