//go:build ignore

// Section 17, Topic 124: Benchmarks
//
// Go has built-in benchmarking support.
//   func BenchmarkXxx(b *testing.B) { ... }
//
// Run: go test -bench=. -benchmem
//
// The benchmark function is called with b.N iterations.
// The framework adjusts b.N to get reliable timing.
//
// GOTCHA: Don't use b.N as input to your function — it's the iteration count.
// GOTCHA: Use b.ResetTimer() after expensive setup.
//
// Run: go run examples/s17_benchmarks.go

package main

import "fmt"

func main() {
	fmt.Println("=== Benchmarks ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Example benchmark
	// ─────────────────────────────────────────────
	fmt.Print(`
// string_test.go
package mypackage

import (
    "strings"
    "testing"
)

func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        s := ""
        for j := 0; j < 100; j++ {
            s += "x"
        }
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        for j := 0; j < 100; j++ {
            sb.WriteString("x")
        }
        _ = sb.String()
    }
}
`)

	// ─────────────────────────────────────────────
	// Output format
	// ─────────────────────────────────────────────
	fmt.Println("-- Typical output --")
	fmt.Print(`
BenchmarkStringConcat-8      100000    12500 ns/op    8192 B/op    99 allocs/op
BenchmarkStringBuilder-8    1000000     1050 ns/op     512 B/op     4 allocs/op

Columns:
  -8        = GOMAXPROCS
  100000    = iterations (b.N)
  12500 ns  = nanoseconds per operation
  8192 B    = bytes allocated per operation (-benchmem)
  99 allocs = allocations per operation (-benchmem)
`)

	// ─────────────────────────────────────────────
	// Benchmark commands
	// ─────────────────────────────────────────────
	fmt.Println("-- Commands --")
	fmt.Println("  go test -bench=.                  # all benchmarks")
	fmt.Println("  go test -bench=BenchmarkConcat     # specific benchmark")
	fmt.Println("  go test -bench=. -benchmem         # include memory stats")
	fmt.Println("  go test -bench=. -benchtime=5s     # run for 5 seconds")
	fmt.Println("  go test -bench=. -count=5          # run 5 times")
	fmt.Println("  go test -bench=. -cpuprofile=cpu.out  # CPU profile")

	// ─────────────────────────────────────────────
	// b.ResetTimer and b.StopTimer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Timer control --")
	fmt.Print(`
func BenchmarkWithSetup(b *testing.B) {
    // Expensive setup:
    data := generateLargeDataset()

    b.ResetTimer()  // exclude setup time

    for i := 0; i < b.N; i++ {
        process(data)
    }
}
`)
}
