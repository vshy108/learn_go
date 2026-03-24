//go:build ignore

// Section 8, Topic 63: Raw String Literals (Backtick)
//
// Go has two string literal forms:
//   Interpreted: "hello\nworld" — processes escape sequences
//   Raw:         `hello\nworld` — no escape processing, can span multiple lines
//
// GOTCHA: Raw strings cannot contain backticks.
// GOTCHA: Raw strings preserve ALL characters including \n literally.
// GOTCHA: Carriage returns (\r) are stripped from raw strings on all platforms.
//
// Run: go run examples/s08_raw_strings.go

package main

import "fmt"

func main() {
	fmt.Println("=== Raw String Literals ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Interpreted vs raw
	// ─────────────────────────────────────────────
	fmt.Println("-- Interpreted vs Raw --")
	interpreted := "Hello\nWorld\t!"
	raw := `Hello\nWorld\t!`

	fmt.Printf("Interpreted: %s\n", interpreted) // processes \n, \t
	fmt.Printf("Raw:         %s\n", raw)         // literal \n, \t

	// ─────────────────────────────────────────────
	// 2. Multi-line strings
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multi-line --")
	query := `SELECT name, age
FROM users
WHERE age > 18
ORDER BY name`
	fmt.Println(query)

	// ─────────────────────────────────────────────
	// 3. Regex patterns (no double-escaping)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Regex patterns --")
	// With interpreted strings, need double backslash:
	regexInterpreted := "\\d+\\.\\d+"
	// With raw strings, single backslash:
	regexRaw := `\d+\.\d+`

	fmt.Printf("Interpreted regex: %s\n", regexInterpreted)
	fmt.Printf("Raw regex:         %s\n", regexRaw)
	// Both produce the same pattern

	// ─────────────────────────────────────────────
	// 4. File paths (especially Windows)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- File paths --")
	winPath := `C:\Users\Alice\Documents\file.txt`
	fmt.Println("Windows path:", winPath)

	// ─────────────────────────────────────────────
	// 5. JSON and HTML templates
	// ─────────────────────────────────────────────
	fmt.Println("\n-- JSON literal --")
	jsonStr := `{
    "name": "Alice",
    "age": 30,
    "hobbies": ["coding", "reading"]
}`
	fmt.Println(jsonStr)

	fmt.Println("\n-- HTML --")
	html := `<!DOCTYPE html>
<html>
<body>
    <h1>Hello, World!</h1>
</body>
</html>`
	fmt.Println(html)

	// ─────────────────────────────────────────────
	// 6. GOTCHA: Cannot include backtick in raw string
	// ─────────────────────────────────────────────
	// This is impossible:
	// s := `text with ` backtick`  // syntax error

	// Workaround: use concatenation
	withBacktick := "code block: `inline code`"
	fmt.Println("\n" + withBacktick)

	// Or use interpreted string:
	fmt.Println("backtick: `")

	// ─────────────────────────────────────────────
	// 7. Quotes inside raw strings work fine
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Quotes in raw strings --")
	withQuotes := `She said "hello" and he said 'hi'`
	fmt.Println(withQuotes)
}
