//go:build ignore

// Section 3, Topic 22: strconv - String <-> Number Conversions
//
// The strconv package converts between strings and basic types.
//
// GOTCHA: Atoi returns (int, error) - always check the error.
// GOTCHA: ParseFloat always uses float64, not float32.
// GOTCHA: string(65) gives "A" (rune), not "65". Use strconv.Itoa(65) for "65".
//
// Run: go run examples/s03_string_conversions.go

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== strconv Conversions ===")
	fmt.Println()

	// 1. Itoa: int -> string
	s := strconv.Itoa(42)
	fmt.Printf("Itoa(42) = %q (%T)\n", s, s)

	// 2. Atoi: string -> int
	n, err := strconv.Atoi("123")
	fmt.Printf("Atoi(\"123\") = %d, err=%v\n", n, err)

	n2, err2 := strconv.Atoi("abc")
	fmt.Printf("Atoi(\"abc\") = %d, err=%v\n", n2, err2)

	// 3. ParseFloat
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("ParseFloat(\"3.14\") = %f (%T)\n", f, f)

	// 4. ParseBool
	b, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool(\"true\") = %t\n", b)

	b2, _ := strconv.ParseBool("1")
	fmt.Printf("ParseBool(\"1\") = %t\n", b2)

	// 5. ParseInt with base
	hex, _ := strconv.ParseInt("FF", 16, 64)
	fmt.Printf("ParseInt(\"FF\", 16) = %d\n", hex)

	bin, _ := strconv.ParseInt("1010", 2, 64)
	fmt.Printf("ParseInt(\"1010\", 2) = %d\n", bin)

	// 6. FormatFloat
	fs := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Printf("FormatFloat(3.14159, 2 decimals) = %q\n", fs)

	// 7. FormatBool
	bs := strconv.FormatBool(true)
	fmt.Printf("FormatBool(true) = %q\n", bs)

	// 8. Common gotcha: string() vs Itoa
	fmt.Println("\n-- GOTCHA: string(65) vs Itoa(65) --")
	fmt.Printf("string(65) = %q (rune conversion!)\n", string(rune(65)))
	fmt.Printf("Itoa(65)   = %q (number to string)\n", strconv.Itoa(65))

	// 9. Sprintf alternative
	fmt.Println("\n-- fmt.Sprintf alternative --")
	mixed := fmt.Sprintf("name=%s, age=%d, score=%.1f", "Alice", 30, 95.5)
	fmt.Println(mixed)
}
