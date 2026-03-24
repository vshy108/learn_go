//go:build ignore

// Section 12, Topic 92: Error Handling Best Practices
//
// Go's philosophy: errors are values, handle them explicitly.
//
// Best practices:
// - Return errors, don't panic
// - Add context with fmt.Errorf("doing X: %w", err)
// - Use sentinel errors for expected conditions
// - Use custom error types for structured data
// - Handle error exactly once (log OR return, not both)
// - Use errors.Is/As for comparison (not ==)
// - Don't ignore errors with _
//
// Run: go run examples/s12_error_best_practices.go

package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrNotFound = errors.New("not found")

// ─────────────────────────────────────────────
// 1. Add context when wrapping
// ─────────────────────────────────────────────
func openFile(name string) (*os.File, error) {
	f, err := os.Open(name)
	if err != nil {
		// GOOD: add context about what we were doing
		return nil, fmt.Errorf("openFile %q: %w", name, err)
	}
	return f, nil
}

// ─────────────────────────────────────────────
// 2. Sentinel error usage
// ─────────────────────────────────────────────
func lookup(key string) (string, error) {
	if key == "key1" {
		return "value1", nil
	}
	return "", ErrNotFound
}

// ─────────────────────────────────────────────
// 3. Avoid panic for expected errors
// ─────────────────────────────────────────────
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide: division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== Error Best Practices ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Wrapping with context
	// ─────────────────────────────────────────────
	_, err := openFile("/nonexistent/file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		// Output: openFile "/nonexistent/file.txt": open /nonexistent/file.txt: no such file...
	}

	// ─────────────────────────────────────────────
	// 2. Checking sentinel errors
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Sentinel check --")
	_, err = lookup("key2")
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Key not found (expected)")
	}

	// ─────────────────────────────────────────────
	// 3. Don't ignore errors
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Don't ignore errors --")
	fmt.Println("BAD:  result, _ := divide(10, 0)  // silently ignores error")
	fmt.Println("GOOD: Check and handle every error")

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("  Error:", err)
	} else {
		fmt.Println("  Result:", result)
	}

	// ─────────────────────────────────────────────
	// 4. Handle once (don't log AND return)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Handle once --")
	fmt.Println("BAD:")
	fmt.Println("  log.Println(err)")
	fmt.Println("  return err  // caller also logs → duplicate logs")
	fmt.Println("GOOD:")
	fmt.Println("  return fmt.Errorf(\"context: %w\", err)  // wrap and return")
}
