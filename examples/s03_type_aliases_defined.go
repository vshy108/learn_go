//go:build ignore

// Section 3, Topic 24: Type Aliases vs Defined Types
//
// Go has two ways to create new type names:
//   - Type alias:  `type MyInt = int` (same type, just another name)
//   - Defined type: `type MyInt int`  (NEW type, different from int)
//
// Type aliases were added in Go 1.9 primarily for gradual code migration.
// Defined types are the standard way to create domain-specific types.
//
// GOTCHA: Defined types do NOT inherit methods from the underlying type.
//         But they DO inherit the underlying type's behavior (operations, etc.).
// GOTCHA: byte is a type alias for uint8; rune is a type alias for int32.
// GOTCHA: A defined type is NOT assignable to its underlying type without conversion.
//
// Run: go run examples/s03_type_aliases_defined.go

package main

import "fmt"

// ─────────────────────────────────────────────
// Defined types (new distinct types)
// ─────────────────────────────────────────────
type Celsius float64
type Fahrenheit float64
type UserID int64

// Methods can be defined on defined types:
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// ─────────────────────────────────────────────
// Type alias (same type, different name)
// ─────────────────────────────────────────────
type Temperature = float64 // alias — Temperature IS float64

func main() {
	fmt.Println("=== Type Aliases vs Defined Types ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Defined types are distinct
	// ─────────────────────────────────────────────
	fmt.Println("-- Defined types --")
	var boiling Celsius = 100
	var bodyTemp Fahrenheit = 98.6
	fmt.Printf("Boiling: %.1f°C = %.1f°F\n", boiling, boiling.ToFahrenheit())
	fmt.Printf("Body:    %.1f°F = %.1f°C\n", bodyTemp, bodyTemp.ToCelsius())

	// Celsius and Fahrenheit are DIFFERENT types:
	// boiling = bodyTemp  // ERROR: cannot use bodyTemp (Fahrenheit) as Celsius
	// This prevents mixing up units — type safety!

	// ─────────────────────────────────────────────
	// 2. Defined types require explicit conversion
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Explicit conversion required --")
	var c Celsius = 100
	var f float64 = float64(c) // must convert explicitly
	fmt.Printf("Celsius(100) → float64(%f)\n", f)

	// Can do arithmetic with underlying type through conversion:
	result := c + Celsius(10) // Celsius + Celsius = Celsius
	fmt.Printf("100°C + 10°C = %.1f°C\n", result)

	// ─────────────────────────────────────────────
	// 3. Type aliases are the SAME type
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Type alias --")
	var temp Temperature = 98.6
	var plain float64 = temp // NO conversion needed — same type!
	fmt.Printf("Temperature=%f, float64=%f\n", temp, plain)
	fmt.Printf("Temperature type: %T\n", temp) // float64 (not Temperature!)

	// ─────────────────────────────────────────────
	// 4. Built-in aliases: byte and rune
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Built-in aliases --")
	var b byte = 42
	var u uint8 = b // no conversion — byte IS uint8
	fmt.Printf("byte(%d) = uint8(%d) — same type\n", b, u)

	var r rune = 'A'
	var i int32 = r // no conversion — rune IS int32
	fmt.Printf("rune(%c) = int32(%d) — same type\n", r, i)

	// ─────────────────────────────────────────────
	// 5. Methods on defined types
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Methods on defined types --")
	var id UserID = 12345
	fmt.Printf("UserID: %d (type: %T)\n", id, id) // learn_go.UserID (or main.UserID)

	// You CANNOT define methods on type aliases (unless the alias
	// is defined in the same package as the original type).
	// type MyAlias = int
	// func (m MyAlias) Double() int { ... }  // ERROR

	// ─────────────────────────────────────────────
	// 6. When to use which
	// ─────────────────────────────────────────────
	// Defined type (`type X int`):
	//   - Create domain types with methods (UserID, Celsius, etc.)
	//   - Prevent mixing up similar types (type safety)
	//   - Add methods to existing types
	//
	// Type alias (`type X = int`):
	//   - Gradual code migration (rename types across packages)
	//   - Provide alternative names for complex types
	//   - byte=uint8 and rune=int32 are built-in examples

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust
	// ─────────────────────────────────────────────
	// Go defined type:  type Celsius float64   (new type, explicit conversion)
	// Rust newtype:     struct Celsius(f64);   (newtype pattern with tuple struct)
	// Go alias:         type Temp = float64    (same type)
	// Rust alias:       type Temp = f64;       (same type)
}
