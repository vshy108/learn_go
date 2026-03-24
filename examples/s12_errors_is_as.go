//go:build ignore

// Section 12, Topic 95: errors.Is and errors.As
//
// errors.Is: Tests if any error in the chain matches a target value.
// errors.As: Tests if any error in the chain matches a target type.
//
// These are the proper way to inspect errors since Go 1.13.
//
// GOTCHA: Don't use == for wrapped errors. Use errors.Is.
// GOTCHA: errors.As takes a **pointer** to the target type.
//
// Run: go run examples/s12_errors_is_as.go

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// Custom error type
type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: not found", e.Name)
}

// Sentinel error
var ErrDatabase = errors.New("database error")

func main() {
	fmt.Println("=== errors.Is and errors.As ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. errors.Is with sentinel
	// ─────────────────────────────────────────────
	err := queryUser("alice")
	fmt.Println("Error:", err)

	if errors.Is(err, ErrDatabase) {
		fmt.Println("→ Is database error (matched through wrapping)")
	}

	// == would fail on wrapped error:
	fmt.Printf("== comparison: %t\n", err == ErrDatabase) // false

	// ─────────────────────────────────────────────
	// 2. errors.As with custom type
	// ─────────────────────────────────────────────
	fmt.Println("\n-- errors.As --")
	err = findItem("widget")
	fmt.Println("Error:", err)

	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		fmt.Printf("→ NotFoundError: Name=%s\n", nfe.Name)
	}

	// ─────────────────────────────────────────────
	// 3. Real-world: checking os errors
	// ─────────────────────────────────────────────
	fmt.Println("\n-- os errors --")
	_, err = os.Open("/nonexistent-file-12345")
	if err != nil {
		fmt.Println("Error:", err)

		// errors.Is with os sentinel:
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("→ File does not exist")
		}

		// errors.As with *fs.PathError:
		var pathErr *fs.PathError
		if errors.As(err, &pathErr) {
			fmt.Printf("→ Op: %s, Path: %s\n", pathErr.Op, pathErr.Path)
		}
	}

	// ─────────────────────────────────────────────
	// 4. Multiple wrapping layers
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Deep wrapping --")
	deepErr := fmt.Errorf("layer3: %w",
		fmt.Errorf("layer2: %w",
			fmt.Errorf("layer1: %w", ErrDatabase)))
	fmt.Println("Error:", deepErr)
	fmt.Printf("Is ErrDatabase: %t\n", errors.Is(deepErr, ErrDatabase)) // true!
}

func queryUser(name string) error {
	return fmt.Errorf("query user %s: %w", name, ErrDatabase)
}

func findItem(name string) error {
	return fmt.Errorf("repository: %w", &NotFoundError{Name: name})
}
