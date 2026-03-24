//go:build ignore

// Section 8, Topic 62: strconv Package — String Conversions
//
// The strconv package converts between strings and basic Go types.
// It's the standard way to parse numbers from strings and format them back.
//
// GOTCHA: Atoi/Itoa work only with int. Use ParseInt/ParseFloat for specific types.
// GOTCHA: Parse functions return errors — always check them!
// GOTCHA: FormatFloat requires format, precision, and bitSize parameters.
//
// Run: go run examples/s08_strconv.go

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== strconv Package ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Itoa / Atoi (int ↔ string)
	// ─────────────────────────────────────────────
	fmt.Println("-- Itoa / Atoi --")
	s := strconv.Itoa(42)
	fmt.Printf("Itoa(42) = %q (type: %T)\n", s, s)

	n, err := strconv.Atoi("42")
	fmt.Printf("Atoi(\"42\") = %d, err=%v\n", n, err)

	_, err = strconv.Atoi("hello")
	fmt.Printf("Atoi(\"hello\") err: %v\n", err)

	_, err = strconv.Atoi("3.14")
	fmt.Printf("Atoi(\"3.14\") err: %v\n", err) // fails! not an integer

	// ─────────────────────────────────────────────
	// 2. ParseInt / ParseFloat / ParseBool
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Parse functions --")

	// ParseInt(s, base, bitSize)
	i64, _ := strconv.ParseInt("255", 10, 64) // decimal
	fmt.Printf("ParseInt(\"255\", 10) = %d\n", i64)

	hex, _ := strconv.ParseInt("FF", 16, 64) // hexadecimal
	fmt.Printf("ParseInt(\"FF\", 16) = %d\n", hex)

	bin, _ := strconv.ParseInt("1010", 2, 64) // binary
	fmt.Printf("ParseInt(\"1010\", 2) = %d\n", bin)

	// ParseFloat(s, bitSize)
	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Printf("ParseFloat(\"3.14159\") = %f\n", f)

	// ParseBool
	b, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool(\"true\") = %t\n", b)

	// Valid bool strings: "1", "t", "T", "TRUE", "true", "True",
	//                     "0", "f", "F", "FALSE", "false", "False"

	// ─────────────────────────────────────────────
	// 3. Format functions (type → string)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Format functions --")
	fmt.Printf("FormatInt(255, 16) = %q\n", strconv.FormatInt(255, 16)) // "ff"
	fmt.Printf("FormatInt(255, 2) = %q\n", strconv.FormatInt(255, 2))   // "11111111"
	fmt.Printf("FormatFloat(3.14, 'f', 2, 64) = %q\n",
		strconv.FormatFloat(3.14, 'f', 2, 64)) // "3.14"
	fmt.Printf("FormatFloat(3.14, 'e', 3, 64) = %q\n",
		strconv.FormatFloat(3.14, 'e', 3, 64)) // "3.140e+00"
	fmt.Printf("FormatBool(true) = %q\n", strconv.FormatBool(true))

	// ─────────────────────────────────────────────
	// 4. Append functions (append formatted to []byte)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Append functions --")
	buf := []byte("value: ")
	buf = strconv.AppendInt(buf, 42, 10)
	fmt.Printf("AppendInt: %s\n", buf)

	buf2 := []byte("pi=")
	buf2 = strconv.AppendFloat(buf2, 3.14159, 'f', 5, 64)
	fmt.Printf("AppendFloat: %s\n", buf2)

	// ─────────────────────────────────────────────
	// 5. Quote / Unquote
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Quote / Unquote --")
	quoted := strconv.Quote("Hello, 世界!\n")
	fmt.Printf("Quote: %s\n", quoted) // "Hello, 世界!\n"

	unquoted, _ := strconv.Unquote(`"Hello\tWorld"`)
	fmt.Printf("Unquote: %s\n", unquoted) // Hello	World

	// ─────────────────────────────────────────────
	// 6. Error handling pattern
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Error handling --")
	input := "not_a_number"
	val, err := strconv.Atoi(input)
	if err != nil {
		if numErr, ok := err.(*strconv.NumError); ok {
			fmt.Printf("NumError: Func=%s, Num=%s, Err=%v\n",
				numErr.Func, numErr.Num, numErr.Err)
		}
	} else {
		fmt.Printf("Parsed: %d\n", val)
	}
}
