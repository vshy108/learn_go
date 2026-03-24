//go:build ignore

// Section 12, Topic 91: Custom Error Types
//
// Custom errors carry structured data beyond a message string.
// Any type implementing Error() string satisfies the error interface.
//
// Common patterns:
//   - Struct with fields (most common)
//   - Named type (type MyError string)
//   - Constant errors (sentinel errors)
//
// GOTCHA: Use pointer receiver for Error() so error comparison works correctly.
//
// Run: go run examples/s12_custom_errors.go

package main

import (
	"fmt"
	"time"
)

// ─────────────────────────────────────────────
// 1. Struct error with context
// ─────────────────────────────────────────────
type HTTPError struct {
	StatusCode int
	Message    string
	URL        string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s (url: %s)", e.StatusCode, e.Message, e.URL)
}

// ─────────────────────────────────────────────
// 2. Error with timestamp
// ─────────────────────────────────────────────
type AppError struct {
	Code    string
	Message string
	Time    time.Time
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Time.Format("15:04:05"), e.Code, e.Message)
}

// ─────────────────────────────────────────────
// 3. Named error type
// ─────────────────────────────────────────────
type Role string

const (
	Admin  Role = "admin"
	Editor Role = "editor"
	Viewer Role = "viewer"
)

type PermissionError struct {
	User     string
	Required Role
	HasRole  Role
}

func (e *PermissionError) Error() string {
	return fmt.Sprintf("user %s has role %s, needs %s", e.User, e.HasRole, e.Required)
}

func main() {
	fmt.Println("=== Custom Error Types ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Using HTTPError
	// ─────────────────────────────────────────────
	err := fetchURL("https://example.com/api")
	if err != nil {
		fmt.Println("Error:", err)

		// Type assert to get details:
		if httpErr, ok := err.(*HTTPError); ok {
			fmt.Printf("  Status: %d\n", httpErr.StatusCode)
			fmt.Printf("  URL: %s\n", httpErr.URL)
		}
	}

	// ─────────────────────────────────────────────
	// Using AppError
	// ─────────────────────────────────────────────
	fmt.Println("\n-- AppError --")
	appErr := &AppError{
		Code:    "DB_CONN",
		Message: "database connection failed",
		Time:    time.Now(),
	}
	fmt.Println(appErr)

	// ─────────────────────────────────────────────
	// Using PermissionError
	// ─────────────────────────────────────────────
	fmt.Println("\n-- PermissionError --")
	if err := checkPermission("alice", Viewer, Admin); err != nil {
		fmt.Println("Error:", err)
		if pe, ok := err.(*PermissionError); ok {
			fmt.Printf("  User %s needs role: %s\n", pe.User, pe.Required)
		}
	}
}

func fetchURL(url string) error {
	return &HTTPError{
		StatusCode: 404,
		Message:    "not found",
		URL:        url,
	}
}

func checkPermission(user string, has, needs Role) error {
	if has != needs {
		return &PermissionError{User: user, Required: needs, HasRole: has}
	}
	return nil
}
