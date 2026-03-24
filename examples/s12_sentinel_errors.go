//go:build ignore

// Section 12, Topic 93: Sentinel Errors
//
// Sentinel errors are package-level error values used for comparison.
// Convention: var ErrXxx = errors.New("...")
//
// Stdlib examples:
//   io.EOF
//   os.ErrNotExist
//   sql.ErrNoRows
//
// GOTCHA: Once exported, sentinel errors become part of your API.
// GOTCHA: Use errors.Is() to compare (supports wrapping), not ==.
//
// Run: go run examples/s12_sentinel_errors.go

package main

import (
	"errors"
	"fmt"
)

// ─────────────────────────────────────────────
// Define sentinel errors
// ─────────────────────────────────────────────
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrConflict     = errors.New("conflict: resource already exists")
	ErrValidation   = errors.New("validation error")
)

func main() {
	fmt.Println("=== Sentinel Errors ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Using sentinel errors
	// ─────────────────────────────────────────────
	err := findUser(0)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found — show 404 page")
	}

	err = findUser(42)
	if err == nil {
		fmt.Println("User found!")
	}

	// ─────────────────────────────────────────────
	// 2. Wrapped sentinels still match
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Wrapped sentinels --")
	err = getProfile(0)
	fmt.Println("Error:", err) // "get profile: find user: not found"
	if errors.Is(err, ErrNotFound) {
		fmt.Println("→ Detected ErrNotFound through wrapping")
	}

	// ─────────────────────────────────────────────
	// 3. Switching on sentinel errors
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Switch pattern --")
	testCases := []int{0, -1, 42}
	for _, id := range testCases {
		err := processRequest(id)
		switch {
		case err == nil:
			fmt.Printf("  ID %d: success\n", id)
		case errors.Is(err, ErrNotFound):
			fmt.Printf("  ID %d: 404 Not Found\n", id)
		case errors.Is(err, ErrUnauthorized):
			fmt.Printf("  ID %d: 401 Unauthorized\n", id)
		default:
			fmt.Printf("  ID %d: 500 %v\n", id, err)
		}
	}

	// ─────────────────────────────────────────────
	// 4. Stdlib sentinel errors
	// ─────────────────────────────────────────────
	// io.EOF          — end of input
	// os.ErrNotExist  — file doesn't exist
	// os.ErrExist     — file already exists
	// os.ErrPermission — permission denied
	// context.Canceled — context was canceled
	// context.DeadlineExceeded — deadline passed
}

func findUser(id int) error {
	if id == 0 {
		return ErrNotFound
	}
	return nil
}

func getProfile(id int) error {
	err := findUser(id)
	if err != nil {
		return fmt.Errorf("get profile: find user: %w", err)
	}
	return nil
}

func processRequest(id int) error {
	if id < 0 {
		return ErrUnauthorized
	}
	return findUser(id)
}
