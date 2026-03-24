//go:build ignore

// Section 17, Topic 123: Table-Driven Tests
//
// Table-driven tests are Go's idiomatic way to test with multiple inputs.
// Define test cases as a slice of structs, loop over them.
//
// Benefits:
//   - Easy to add new test cases
//   - Each case is clearly named
//   - DRY — one test body, many inputs
//
// Run: go run examples/s17_table_driven_tests.go

package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func fizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

func main() {
	fmt.Println("=== Table-Driven Tests ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Example table-driven test
	// ─────────────────────────────────────────────
	fmt.Print(`
// In abs_test.go:
func TestAbs(t *testing.T) {
    tests := []struct {
        name  string
        input int
        want  int
    }{
        {"positive", 5, 5},
        {"negative", -3, 3},
        {"zero", 0, 0},
        {"min negative", -1, 1},
        {"large negative", -100, 100},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := abs(tt.input)
            if got != tt.want {
                t.Errorf("abs(%d) = %d, want %d",
                    tt.input, got, tt.want)
            }
        })
    }
}
`)

	// ─────────────────────────────────────────────
	// Demo of the pattern
	// ─────────────────────────────────────────────
	fmt.Println("-- Demo: abs --")
	absTests := []struct {
		input int
		want  int
	}{
		{5, 5}, {-3, 3}, {0, 0}, {-100, 100},
	}
	for _, tt := range absTests {
		got := abs(tt.input)
		status := "PASS"
		if got != tt.want {
			status = "FAIL"
		}
		fmt.Printf("  abs(%d) = %d (want %d) %s\n", tt.input, got, tt.want, status)
	}

	fmt.Println("\n-- Demo: fizzBuzz --")
	fbTests := []struct {
		input int
		want  string
	}{
		{1, "1"}, {3, "Fizz"}, {5, "Buzz"}, {15, "FizzBuzz"}, {7, "7"},
	}
	for _, tt := range fbTests {
		got := fizzBuzz(tt.input)
		status := "PASS"
		if got != tt.want {
			status = "FAIL"
		}
		fmt.Printf("  fizzBuzz(%d) = %q (want %q) %s\n", tt.input, got, tt.want, status)
	}

	// ─────────────────────────────────────────────
	// t.Run creates subtests
	// ─────────────────────────────────────────────
	// go test -run TestAbs/positive   — run only "positive" subtest
	// go test -run TestAbs/negative   — run only "negative" subtest
	// go test -v -run TestAbs         — verbose output for all subtests
}
