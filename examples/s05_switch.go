//go:build ignore

// Section 5, Topic 38: switch Statement — No Fallthrough by Default
//
// Go's switch is much cleaner than C/Java:
//   - No `break` needed (cases don't fall through by default)
//   - Cases can have multiple values
//   - Cases can have expressions (not just constants)
//   - switch without a condition is like if/else if
//
// GOTCHA: Cases do NOT fall through by default (opposite of C).
//         Use `fallthrough` keyword to explicitly fall through.
// GOTCHA: `fallthrough` is unconditional — it doesn't check the next case's condition.
// GOTCHA: switch can have an initializer like if.
//
// Run: go run examples/s05_switch.go

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("=== switch Statement ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic switch
	// ─────────────────────────────────────────────
	fmt.Println("-- Basic switch --")
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday":
		fmt.Println("Second day")
	case "Friday":
		fmt.Println("Almost weekend!")
	default:
		fmt.Println("Some other day")
	}

	// ─────────────────────────────────────────────
	// 2. Multiple values per case
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multiple values per case --")
	today := time.Now().Weekday()
	switch today {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!")
	default:
		fmt.Println("Weekday")
	}

	// ─────────────────────────────────────────────
	// 3. Switch without condition (like if/else if)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Conditionless switch --")
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// ─────────────────────────────────────────────
	// 4. Switch with initializer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Switch with initializer --")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Unknown:", os)
	}

	// ─────────────────────────────────────────────
	// 5. Case expressions (not just constants)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Expression cases --")
	n := 15
	switch {
	case n%15 == 0:
		fmt.Println("FizzBuzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(n)
	}

	// ─────────────────────────────────────────────
	// 6. fallthrough (explicit, unconditional)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- fallthrough --")
	x := 1
	switch x {
	case 1:
		fmt.Println("One")
		fallthrough // unconditionally enters next case
	case 2:
		fmt.Println("Two (via fallthrough)")
		// no fallthrough here, so it stops
	case 3:
		fmt.Println("Three (not reached)")
	}

	// GOTCHA: fallthrough doesn't check the next case's condition!
	// switch 1 {
	// case 1:
	//     fallthrough
	// case 999:          // this RUNS even though 1 != 999
	//     fmt.Println("!")
	// }

	// ─────────────────────────────────────────────
	// 7. break in switch (exits the switch, not a loop)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- break in switch --")
	switch 1 {
	case 1:
		fmt.Println("Before break")
		if true {
			break // exits the switch
		}
		fmt.Println("After break (never reached)")
	}

	// To break an outer loop from inside a switch, use labels:
	fmt.Println("\n-- break with label --")
Loop:
	for i := 0; i < 5; i++ {
		switch i {
		case 3:
			fmt.Println("Breaking outer loop at i=3")
			break Loop // breaks the for loop, not just the switch
		}
		fmt.Printf("i=%d ", i)
	}
	fmt.Println()

	// ─────────────────────────────────────────────
	// Comparison: Go vs C vs Rust
	// ─────────────────────────────────────────────
	// Go:   no fallthrough by default, use `fallthrough` to opt in
	// C:    fallthrough by default, use `break` to opt out
	// Rust: match — exhaustive, no fallthrough, can return values
}
