//go:build ignore

// Section 16, Topic 118: Type Constraints
//
// Constraints restrict what types can be used as type parameters.
// Constraints are interfaces (with Go 1.18 extensions).
//
// Built-in constraints:
//   any          — no restriction (alias for interface{})
//   comparable   — supports == and !=
//
// Standard library: golang.org/x/exp/constraints
//   constraints.Ordered   — supports <, >, <=, >=
//   constraints.Integer   — all integer types
//   constraints.Float     — all float types
//   constraints.Signed    — signed integers
//   constraints.Unsigned  — unsigned integers
//
// GOTCHA: ~ in constraints means "underlying type" (includes named types).
//
// Run: go run examples/s16_type_constraints.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Custom constraint with type union
// ─────────────────────────────────────────────
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// ─────────────────────────────────────────────
// 2. Constraint with method requirement
// ─────────────────────────────────────────────
type Stringer interface {
	String() string
}

func JoinStrings[T Stringer](items []T, sep string) string {
	result := ""
	for i, item := range items {
		if i > 0 {
			result += sep
		}
		result += item.String()
	}
	return result
}

// ─────────────────────────────────────────────
// 3. Combined constraint (methods + types)
// ─────────────────────────────────────────────
type OrderedStringer interface {
	~int | ~string
	String() string
}

// ─────────────────────────────────────────────
// 4. comparable constraint
// ─────────────────────────────────────────────
func Contains[T comparable](s []T, target T) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func Index[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}

// Named type (~ includes it in Number constraint)
type Celsius float64
type Score int

func main() {
	fmt.Println("=== Type Constraints ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Number constraint
	// ─────────────────────────────────────────────
	fmt.Println("Sum ints:", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println("Sum floats:", Sum([]float64{1.1, 2.2, 3.3}))

	// Named types work because of ~:
	temps := []Celsius{36.5, 37.0, 38.5}
	fmt.Println("Sum temps:", Sum(temps))

	scores := []Score{90, 85, 92}
	fmt.Println("Sum scores:", Sum(scores))

	// ─────────────────────────────────────────────
	// comparable constraint
	// ─────────────────────────────────────────────
	fmt.Println("\n-- comparable --")
	fmt.Println("Contains 3:", Contains([]int{1, 2, 3}, 3))
	fmt.Println("Contains 'x':", Contains([]string{"a", "b"}, "x"))
	fmt.Println("Index of 'b':", Index([]string{"a", "b", "c"}, "b"))

	// ─────────────────────────────────────────────
	// ~ (tilde) = underlying type
	// ─────────────────────────────────────────────
	// ~int matches: int, type MyInt int, type Score int
	// int (without ~) matches ONLY: int
}
