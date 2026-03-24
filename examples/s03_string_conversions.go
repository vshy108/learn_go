//go:build ignore

// Section 3, Topic 22: strconv — String ↔ Number Conversions
//
// The `strconv` package provides functions for converting between strings
// and basic types. This is Go's equivalent of Python's int(), str() or
// Rust's .parse::<T>() and .to_string().
//
// Key functions:
//   - Atoi / Itoa:     string ↔ int (convenience wrappers)
//   - ParseInt/FormatInt: string ↔ int with base and size control
//   - ParseFloat/FormatFloat: string ↔ float with precision control
//   - ParseBool/FormatBool: string ↔ bool
//
// GOTCHA: Atoi returns (int, error) — you MUST handle the error.
//         Go does NOT throw exceptions; errors are return values.
//
// Run: go run examples/s03_string_conversions.go

package main















































































































}	// fmt.Sscanf("42", "%d", &n) → parse from string	// fmt.Sprintf("%d", 42) → "42" is another way to convert int→string	//	// Rust:   "42".parse::<i32>() → Ok(42) (returns Result)	// Go:     strconv.Atoi("42") → (42, nil) (returns error, not exception)	// Python: int("42") → 42 (or raises ValueError)	// ─────────────────────────────────────────────	// 6. GOTCHA: Atoi is NOT the same as int()	// ─────────────────────────────────────────────	fmt.Printf("FormatBool(false) = %s\n", strconv.FormatBool(false))	fmt.Printf("FormatBool(true)  = %s\n", strconv.FormatBool(true))	fmt.Println("\n-- FormatBool --")	// "yes" and "no" are NOT valid — returns error!	}		fmt.Printf("  ParseBool(%q) = %t, err=%v\n", val, b, err)		b, err := strconv.ParseBool(val)	for _, val := range []string{"true", "TRUE", "1", "T", "false", "0", "F", "yes"} {	// Also:    "0", "f", "F", "FALSE", "false", "False"	// Accepts: "1", "t", "T", "TRUE", "true", "True"	fmt.Println("\n-- ParseBool --")	// ─────────────────────────────────────────────	// 5. ParseBool / FormatBool	// ─────────────────────────────────────────────	fmt.Printf("FormatFloat(3.14159, 'g', -1, 64) = %s\n", strconv.FormatFloat(3.14159, 'g', -1, 64))	fmt.Printf("FormatFloat(3.14159, 'e', 4, 64)  = %s\n", strconv.FormatFloat(3.14159, 'e', 4, 64))	fmt.Printf("FormatFloat(3.14159, 'f', 2, 64)  = %s\n", strconv.FormatFloat(3.14159, 'f', 2, 64))	//   prec: number of digits after decimal (-1 = smallest exact representation)	//   fmt: 'f' decimal, 'e' scientific, 'g' compact, 'b' binary exponent	// FormatFloat(f, fmt, prec, bitSize)	fmt.Println("\n-- FormatFloat --")	fmt.Printf("ParseFloat(\"+Inf\") = %f\n", f)	f, _ = strconv.ParseFloat("+Inf", 64)	fmt.Printf("ParseFloat(\"NaN\") = %f\n", f)	f, _ = strconv.ParseFloat("NaN", 64)	// Special values:	fmt.Printf("ParseFloat(\"1e10\", 64) = %f\n", f)	f, _ = strconv.ParseFloat("1e10", 64)	fmt.Printf("ParseFloat(\"3.14159\", 64) = %f\n", f)	f, _ := strconv.ParseFloat("3.14159", 64) // 64 = float64	fmt.Println("\n-- ParseFloat --")	// ─────────────────────────────────────────────	// 4. ParseFloat / FormatFloat	// ─────────────────────────────────────────────	fmt.Printf("FormatInt(42, 36) = %s (base36)\n", strconv.FormatInt(42, 36))	fmt.Printf("FormatInt(42, 16) = %s (hex)\n", strconv.FormatInt(42, 16))	fmt.Printf("FormatInt(42, 8)  = %s (octal)\n", strconv.FormatInt(42, 8))	fmt.Printf("FormatInt(42, 2)  = %s (binary)\n", strconv.FormatInt(42, 2))	fmt.Println("\n-- FormatInt --")	// ─────────────────────────────────────────────	// 3. FormatInt — int to string with base	// ─────────────────────────────────────────────	fmt.Printf("ParseInt(\"999\", 10, 8): err=%v\n", err)	_, err = strconv.ParseInt("999", 10, 8) // max int8 = 127	// Overflow detection:	fmt.Printf("ParseInt(\"0o77\", 0, 64) = %d\n", v)	v, _ = strconv.ParseInt("0o77", 0, 64)    // auto-detect octal	fmt.Printf("ParseInt(\"0b1010\", 0, 64) = %d\n", v)	v, _ = strconv.ParseInt("0b1010", 0, 64)  // auto-detect binary	fmt.Printf("ParseInt(\"0xFF\", 0, 64) = %d\n", v)	v, _ = strconv.ParseInt("0xFF", 0, 64)    // auto-detect hex	fmt.Printf("ParseInt(\"42\", 10, 64) = %d\n", v)	v, _ := strconv.ParseInt("42", 10, 64)   // decimal, fits int64	// base=0 auto-detects: 0x→hex, 0o→octal, 0b→binary, else decimal	// ParseInt(s, base, bitSize) → (int64, error)	fmt.Println("\n-- ParseInt --")	// ─────────────────────────────────────────────	// 2. ParseInt — more control (base, bit size)	// ─────────────────────────────────────────────	fmt.Printf("Itoa(42) = %q\n", s) // "42"	s := strconv.Itoa(42)	fmt.Println("\n-- Itoa (int → string) --")	fmt.Printf("Atoi(\"  42  \") = %d, err=%v (no trimming!)\n", n, err) // error!	n, err = strconv.Atoi("  42  ")	fmt.Printf("Atoi(\"\") = %d, err=%v\n", n, err) // 0, error	n, err = strconv.Atoi("")	fmt.Printf("Atoi(\"not_a_number\") = %d, err=%v\n", n, err) // 0, error	n, err = strconv.Atoi("not_a_number")	fmt.Printf("Atoi(\"42\") = %d, err=%v\n", n, err)	n, err := strconv.Atoi("42")	fmt.Println("-- Atoi (string → int) --")	// ─────────────────────────────────────────────	// 1. Atoi / Itoa — the quick int converters	// ─────────────────────────────────────────────	fmt.Println()	fmt.Println("=== strconv Package ===")func main() {)	"strconv"	"fmt"import (