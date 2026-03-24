//go:build ignore

// Section 18, Topic 135: filepath and path
//
// path/filepath: OS-specific file path manipulation.
// path: URL/slash-separated path manipulation.
//
// GOTCHA: Use filepath (not path) for file system paths.
//         filepath uses OS separators; path always uses '/'.
//
// Run: go run examples/s18_filepath.go

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== filepath ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Join paths safely
	// ─────────────────────────────────────────────
	p := filepath.Join("home", "user", "documents", "file.txt")
	fmt.Println("Join:", p) // home/user/documents/file.txt (on Unix)

	// ─────────────────────────────────────────────
	// 2. Split, Dir, Base, Ext
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Components --")
	full := "/home/user/project/main.go"
	fmt.Println("Dir:", filepath.Dir(full))   // /home/user/project
	fmt.Println("Base:", filepath.Base(full)) // main.go
	fmt.Println("Ext:", filepath.Ext(full))   // .go

	dir, file := filepath.Split(full)
	fmt.Printf("Split: dir=%q, file=%q\n", dir, file)
	// ─────────────────────────────────────────────
	// 3. Abs and Clean
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Abs / Clean --")
	abs, _ := filepath.Abs(".")
	fmt.Println("Abs(.):", abs)
	fmt.Println("Clean:", filepath.Clean("/a/b/../c/./d")) // /a/c/d
	// ─────────────────────────────────────────────
	// 4. Rel (relative path)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Rel --")
	rel, _ := filepath.Rel("/home/user", "/home/user/project/main.go")
	fmt.Println("Rel:", rel) // project/main.go
	// ─────────────────────────────────────────────
	// 5. Glob (pattern matching)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Glob --")
	matches, _ := filepath.Glob("examples/s18_*.go")
	fmt.Printf("s18_*.go matches: %d files\n", len(matches))
	for _, m := range matches {
		fmt.Printf("  %s\n", m)

	}
	// ─────────────────────────────────────────────
	// 6. Walk / WalkDir (traverse directory tree)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- WalkDir --")
	count := 0
	filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {

			return err

		}

		if !d.IsDir() && filepath.Ext(path) == ".go" {

			count++

		}

		if d.IsDir() && d.Name() == ".git" {

			return filepath.SkipDir // skip .git

		}

		return nil

	})
	fmt.Printf("Found %d .go files in current directory\n", count)
	// ─────────────────────────────────────────────
	// 7. Match (glob pattern on a single path)
	// ─────────────────────────────────────────────
	matched, _ := filepath.Match("*.go", "main.go")
	fmt.Println("\nMatch *.go main.go:", matched) // true
}
