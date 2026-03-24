//go:build ignore

// Section 14, Topic 110: Context for Cancellation
//
// context.Context carries deadlines, cancellation signals, and values
// across API boundaries and goroutines.
//
// Key functions:
//   context.Background()      — root context
//   context.WithCancel(ctx)   — cancel signal
//   context.WithTimeout(ctx)  — auto-cancel after duration
//   context.WithDeadline(ctx) — auto-cancel at time
//   context.WithValue(ctx)    — attach key-value pairs
//
// GOTCHA: Always defer cancel() to prevent resource leaks.
// GOTCHA: Don't store context in structs — pass as first parameter.
// GOTCHA: context.Value is untyped — use custom key types to avoid collisions.
//
// Run: go run examples/s14_context_cancellation.go

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Context Cancellation ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. WithCancel — manual cancellation
	// ─────────────────────────────────────────────
	fmt.Println("-- WithCancel --")
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("  Goroutine: cancelled!", ctx.Err())
				return
			default:
				fmt.Println("  Goroutine: working...")
				time.Sleep(20 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(60 * time.Millisecond)
	cancel() // signal cancellation
	time.Sleep(30 * time.Millisecond)

	// ─────────────────────────────────────────────
	// 2. WithTimeout — auto-cancel after duration
	// ─────────────────────────────────────────────
	fmt.Println("\n-- WithTimeout --")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel2()

	result, err := fetchData(ctx2)
	if err != nil {
		fmt.Println("  Fetch error:", err)
	} else {
		fmt.Println("  Fetch result:", result)
	}

	// ─────────────────────────────────────────────
	// 3. WithValue — passing request-scoped data
	// ─────────────────────────────────────────────
	fmt.Println("\n-- WithValue --")
	type contextKey string
	const userKey contextKey = "user"

	ctx3 := context.WithValue(context.Background(), userKey, "Alice")
	processRequest(ctx3, userKey)

	// ─────────────────────────────────────────────
	// Context guidelines:
	// 1. First param: func(ctx context.Context, ...)
	// 2. Don't store in structs
	// 3. Always defer cancel()
	// 4. Use context.TODO() as placeholder during refactoring
	// 5. Don't pass nil context — use context.Background()
}

func fetchData(ctx context.Context) (string, error) {
	select {
	case <-time.After(30 * time.Millisecond):
		return "data from server", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func processRequest(ctx context.Context, key any) {
	if user, ok := ctx.Value(key).(string); ok {
		fmt.Printf("  Processing request for user: %s\n", user)
	}
}
