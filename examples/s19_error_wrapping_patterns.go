//go:build ignore

// Section 19, Topic 143: Error Wrapping Patterns
//
// Advanced error patterns for production Go code.
//
// Run: go run examples/s19_error_wrapping_patterns.go

package main

import (
	"errors"
	"fmt"
)

// ─────────────────────────────────────────────
// 1. Sentinel errors with context
// ─────────────────────────────────────────────
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrValidation   = errors.New("validation error")
)

// ─────────────────────────────────────────────
// 2. Custom error with multiple wrapping (Go 1.20+)
// ─────────────────────────────────────────────
type MultiError struct {
	errs []error
}

func (me *MultiError) Error() string {
	msgs := make([]string, len(me.errs))
	for i, err := range me.errs {
		msgs[i] = err.Error()
	}
	return fmt.Sprintf("multiple errors: %v", msgs)
}

func (me *MultiError) Unwrap() []error {
	return me.errs
}

func (me *MultiError) Add(err error) {
	me.errs = append(me.errs, err)
}

// ─────────────────────────────────────────────
// 3. Domain error type
// ─────────────────────────────────────────────
type AppError struct {
	Code    string
	Message string
	Err     error // wrapped error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(code, msg string, err error) *AppError {
	return &AppError{Code: code, Message: msg, Err: err}
}

func main() {
	fmt.Println("=== Error Wrapping Patterns ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Chain wrapping
	// ─────────────────────────────────────────────
	fmt.Println("-- Chain wrapping --")
	err := serviceLayer()
	fmt.Println("Error:", err)

	// Unwrap through chain:
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Root cause: not found")
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		fmt.Printf("App error: code=%s, msg=%s\n", appErr.Code, appErr.Message)
	}

	// ─────────────────────────────────────────────
	// 2. Multi-error (Go 1.20+ multiple unwrap)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multi-error --")
	me := &MultiError{}
	me.Add(fmt.Errorf("field name: %w", ErrValidation))
	me.Add(fmt.Errorf("field email: %w", ErrValidation))
	me.Add(ErrUnauthorized)

	fmt.Println("Error:", me)

	// errors.Is checks all wrapped errors:
	fmt.Println("Has validation:", errors.Is(me, ErrValidation))
	fmt.Println("Has unauthorized:", errors.Is(me, ErrUnauthorized))
	fmt.Println("Has not found:", errors.Is(me, ErrNotFound))

	// ─────────────────────────────────────────────
	// 3. fmt.Errorf with %w (simple wrapping)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- fmt.Errorf %w --")
	err = fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", ErrNotFound))
	fmt.Println("Error:", err)
	fmt.Println("Is ErrNotFound:", errors.Is(err, ErrNotFound))

	// ─────────────────────────────────────────────
	// Best practices:
	// ─────────────────────────────────────────────
	// 1. Wrap errors with context: fmt.Errorf("doing X: %w", err)
	// 2. Use sentinel errors for known conditions
	// 3. Use custom error types for structured data
	// 4. Check errors with errors.Is/As (not ==)
	// 5. Handle error once (don't log AND return)
}

func dataLayer() error {
	return ErrNotFound
}

func repoLayer() error {
	err := dataLayer()
	if err != nil {
		return NewAppError("REPO_001", "user not found in database", err)
	}
	return nil
}

func serviceLayer() error {
	err := repoLayer()
	if err != nil {
		return fmt.Errorf("service: get user: %w", err)
	}
	return nil
}
