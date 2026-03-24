//go:build ignore

// Section 18, Topic 133: time Package
//
// Go's time package handles times, durations, and formatting.
//
// GOTCHA: Go uses a reference time for formatting: Mon Jan 2 15:04:05 MST 2006
//         (1/2 3:4:5 in 2006 — each component is unique!)
//
// Run: go run examples/s18_time.go

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== time Package ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Current time
	// ─────────────────────────────────────────────
	now := time.Now()
	fmt.Println("Now:", now)
	fmt.Println("UTC:", now.UTC())
	fmt.Printf("Year: %d, Month: %s, Day: %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", now.Hour(), now.Minute(), now.Second())
	fmt.Println("Weekday:", now.Weekday())
	fmt.Println("Unix timestamp:", now.Unix())

	// ─────────────────────────────────────────────
	// 2. Time formatting (reference time!)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Formatting --")
	// The reference time: Mon Jan 2 15:04:05 MST 2006
	fmt.Println("RFC3339:", now.Format(time.RFC3339))
	fmt.Println("Custom:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("Date only:", now.Format("2006-01-02"))
	fmt.Println("Time only:", now.Format("15:04:05"))
	fmt.Println("US format:", now.Format("01/02/2006"))
	fmt.Println("With day:", now.Format("Mon, 02 Jan 2006"))

	// ─────────────────────────────────────────────
	// 3. Parsing time strings
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Parsing --")
	t, err := time.Parse("2006-01-02", "2024-06-15")
	if err != nil {
		fmt.Println("Parse error:", err)
	} else {
		fmt.Println("Parsed:", t)
	}

	t2, _ := time.Parse(time.RFC3339, "2024-01-15T10:30:00Z")
	fmt.Println("RFC3339:", t2)

	// ─────────────────────────────────────────────
	// 4. Duration
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Duration --")
	d := 2*time.Hour + 30*time.Minute
	fmt.Println("Duration:", d)
	fmt.Printf("In seconds: %.0f\n", d.Seconds())
	fmt.Printf("In minutes: %.0f\n", d.Minutes())

	future := now.Add(d)
	fmt.Println("Now +2h30m:", future.Format("15:04:05"))

	// ─────────────────────────────────────────────
	// 5. Time comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Comparison --")
	t1, _ := time.Parse("2006-01-02", "2024-01-01")
	t3, _ := time.Parse("2006-01-02", "2024-12-31")
	fmt.Println("Before:", t1.Before(t3)) // true
	fmt.Println("After:", t1.After(t3))   // false
	fmt.Println("Equal:", t1.Equal(t1))   // true
	fmt.Println("Diff:", t3.Sub(t1))      // 8760h0m0s (365 days)

	// ─────────────────────────────────────────────
	// 6. Ticker and Timer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Timer --")
	timer := time.NewTimer(50 * time.Millisecond)
	<-timer.C
	fmt.Println("Timer fired!")

	// Sleep:
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Slept for: %v\n", time.Since(start))
}
