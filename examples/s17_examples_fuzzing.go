//go:build ignore

// Section 17, Topic 128: Example Tests and Fuzzing
//
// Example tests: Testable documentation. They show up in godoc.
// Fuzz tests (Go 1.18+): Automated testing with generated inputs.
//
// Run: go run examples/s17_examples_fuzzing.go

package main

import "fmt"

func main() {
	fmt.Println("=== Example Tests & Fuzzing ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Example tests (testable documentation)
	// ─────────────────────────────────────────────
	fmt.Println("-- Example tests --")
	fmt.Print(`
// In math_test.go:
func ExampleAdd() {
    fmt.Println(add(2, 3))
    // Output: 5
}

func ExampleAdd_negative() {
    fmt.Println(add(-1, 1))
    // Output: 0
}

// Examples appear in godoc documentation!
// go test verifies the // Output: comment matches.
`)

	// ─────────────────────────────────────────────
	// 2. Example naming convention
	// ─────────────────────────────────────────────
	fmt.Println("-- Naming --")
	fmt.Println("  ExampleFoo()        — example for function Foo")
	fmt.Println("  ExampleBar_suffix() — example variant for Bar")
	fmt.Println("  ExampleMyType()     — example for type MyType")
	fmt.Println("  ExampleMyType_Method() — for method")
	fmt.Println("  Example()           — package-level example")

	// ─────────────────────────────────────────────
	// 3. Fuzz testing (Go 1.18+)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Fuzz testing --")
	fmt.Print(`
// In reverse_test.go:
func FuzzReverse(f *testing.F) {
    // Seed corpus (initial test cases):
    f.Add("hello")
    f.Add("world")
    f.Add("")
    f.Add("12345")

    // Fuzz function:
    f.Fuzz(func(t *testing.T, s string) {
        reversed := Reverse(s)
        doubleReversed := Reverse(reversed)
        if s != doubleReversed {
            t.Errorf("double reverse mismatch: %q → %q → %q",
                s, reversed, doubleReversed)
        }
    })
}
`)

	// ─────────────────────────────────────────────
	// 4. Fuzz commands
	// ─────────────────────────────────────────────
	fmt.Println("-- Fuzz commands --")
	fmt.Println("  go test -fuzz=FuzzReverse          # run fuzzer")
	fmt.Println("  go test -fuzz=FuzzReverse -fuzztime=30s  # 30 seconds")
	fmt.Println("  go test -fuzz=FuzzReverse -fuzztime=1000x  # 1000 iterations")
	fmt.Println()
	fmt.Println("  Corpus: testdata/fuzz/FuzzReverse/")
	fmt.Println("  Failed inputs saved there for regression testing")

	// ─────────────────────────────────────────────
	// 5. Fuzz supported types
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Fuzz types --")
	fmt.Println("  string, []byte, int, int8-64, uint, uint8-64,")
	fmt.Println("  float32, float64, bool, rune")
}
