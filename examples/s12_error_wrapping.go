//go:build ignore

// Section 12, Topic 92: Error Wrapping (Go 1.13+)
//
// Error wrapping lets you add context while preserving the original error.
//   fmt.Errorf("context: %w", err)   — wraps err
//   errors.Unwrap(err)               — gets the wrapped error
//   errors.Is(err, target)           — checks error chain
//   errors.As(err, &target)          — type-checks error chain
//
// GOTCHA: Use %w (not %v) to wrap errors. %v loses the chain.
// GOTCHA: errors.Is checks the entire chain, == only checks the top level.
//
// Run: go run examples/s12_error_wrapping.go

package main

import (
	"errors"
	"fmt"
	"os"
)

// Sentinel errors:
var (
	ErrNotFound  = errors.New("not found")
	ErrForbidden = errors.New("forbidden")
	ErrInternal  = errors.New("internal error")
)

func main() {
	fmt.Println("=== Error Wrapping ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Wrapping with %w
	// ─────────────────────────────────────────────
	err := loadConfig("settings.json")
	fmt.Println("Error:", err)
	// Output: config: open file: settings.json: not found

	// ─────────────────────────────────────────────
	// 2. errors.Is — check if error chain contains target
	// ─────────────────────────────────────────────
	fmt.Println("\n-- errors.Is --")
	fmt.Printf("Is ErrNotFound: %t\n", errors.Is(err, ErrNotFound))   // true
	fmt.Printf("Is ErrForbidden: %t\n", errors.Is(err, ErrForbidden)) // false

	// Direct comparison FAILS for wrapped errors:
	fmt.Printf("== ErrNotFound: %t\n", err == ErrNotFound) // false (wrapped!)

	// ─────────────────────────────────────────────
	// 3. errors.Unwrap — peel one layer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- errors.Unwrap --")
	e1 := err
	for e1 != nil {
		fmt.Printf("  → %v\n", e1)
		e1 = errors.Unwrap(e1)
	}

	// ─────────────────────────────────────────────
	// 4. errors.As — type assertion on chain
	// ─────────────────────────────────────────────
	fmt.Println("\n-- errors.As --")
	wrappedPath := fmt.Errorf("processing: %w", &os.PathError{
		Op:   "open",
		Path: "/etc/secret",
		Err:  os.ErrPermission,
	})
	fmt.Println("Error:", wrappedPath)

	var pathErr *os.PathError
	if errors.As(wrappedPath, &pathErr) {
		fmt.Printf("  Path: %s\n", pathErr.Path)
		fmt.Printf("  Op: %s\n", pathErr.Op)
	}

	// errors.Is also works through wrapping:
	fmt.Printf("  Is ErrPermission: %t\n", errors.Is(wrappedPath, os.ErrPermission))

	// ─────────────────────────────────────────────
	// 5. %v vs %w
	// ─────────────────────────────────────────────
	fmt.Println("\n-- %v vs %w --")
	base := ErrNotFound
	withW := fmt.Errorf("wrapped: %w", base)                               // preserves chain
	withV := fmt.Errorf("wrapped: %v", base)                               // BREAKS chain!
	fmt.Printf("%%w: Is ErrNotFound: %t\n", errors.Is(withW, ErrNotFound)) // true
	fmt.Printf("%%v: Is ErrNotFound: %t\n", errors.Is(withV, ErrNotFound)) // false!
}

func openFile(name string) error {
	return fmt.Errorf("%s: %w", name, ErrNotFound)
}

func loadConfig(name string) error {
	err := openFile(name)
	if err != nil {
		return fmt.Errorf("config: open file: %w", err)
	}
	return nil
}
