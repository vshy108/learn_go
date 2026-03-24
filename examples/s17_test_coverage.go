//go:build ignore

// Section 17, Topic 127: Test Coverage
//
// Go has built-in code coverage analysis.
//
// Commands:
//   go test -cover                    — show coverage percentage
//   go test -coverprofile=cover.out   — generate coverage data
//   go tool cover -html=cover.out     — open in browser
//   go tool cover -func=cover.out     — per-function coverage
//
// GOTCHA: 100% coverage doesn't mean bug-free.
// GOTCHA: Coverage counts executed lines, not correctness.
//
// Run: go run examples/s17_test_coverage.go

package main

import "fmt"

func main() {
	fmt.Println("=== Test Coverage ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Generating coverage
	// ─────────────────────────────────────────────
	fmt.Println("-- Commands --")
	fmt.Println("  go test -cover ./...")
	fmt.Println("  → PASS coverage: 78.5% of statements")
	fmt.Println()
	fmt.Println("  go test -coverprofile=cover.out ./...")
	fmt.Println("  go tool cover -html=cover.out")
	fmt.Println("  → Opens browser with highlighted source")
	fmt.Println()
	fmt.Println("  go tool cover -func=cover.out")
	fmt.Println("  → Per-function coverage breakdown")

	// ─────────────────────────────────────────────
	// 2. Coverage output
	// ─────────────────────────────────────────────
	fmt.Println("\n-- go tool cover -func output --")
	fmt.Print(`
mypackage/math.go:5:    Add         100.0%
mypackage/math.go:9:    Divide      85.7%
mypackage/math.go:20:   Factorial   60.0%
total:                  (statements) 78.5%
`)

	// ─────────────────────────────────────────────
	// 3. Coverage modes
	// ─────────────────────────────────────────────
	fmt.Println("-- Coverage modes --")
	fmt.Println("  -covermode=set    — was each statement executed? (default)")
	fmt.Println("  -covermode=count  — how many times?")
	fmt.Println("  -covermode=atomic — like count, thread-safe (for -race)")

	// ─────────────────────────────────────────────
	// 4. In CI/CD
	// ─────────────────────────────────────────────
	fmt.Println("\n-- CI/CD --")
	fmt.Print(`
# Fail if coverage < 80%:
go test -coverprofile=cover.out ./...
COVERAGE=$(go tool cover -func=cover.out | grep total | awk '{print $3}' | tr -d '%')
if [ $(echo "$COVERAGE < 80" | bc) -eq 1 ]; then
    echo "Coverage $COVERAGE% is below 80%"
    exit 1
fi
`)

	// ─────────────────────────────────────────────
	// 5. Coverage best practices
	// ─────────────────────────────────────────────
	fmt.Println("-- Best practices --")
	fmt.Println("  ✓ Aim for meaningful coverage, not 100%")
	fmt.Println("  ✓ Cover error paths and edge cases")
	fmt.Println("  ✓ Use -coverprofile in CI/CD")
	fmt.Println("  ✓ Review uncovered code to decide if tests are needed")
	fmt.Println("  ✗ Don't write trivial tests just for coverage")
}
