//go:build ignore

// Section 18, Topic 129: os and io Packages
//
// os: Operating system functions (files, env, process).
// io: Core I/O interfaces and helpers.
//
// Run: go run examples/s18_os_io.go

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== os and io Packages ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Environment variables
	// ─────────────────────────────────────────────
	fmt.Println("-- Environment --")
	home, _ := os.UserHomeDir()
	fmt.Println("HOME:", home)
	fmt.Println("PATH:", os.Getenv("PATH")[:50]+"...")
	fmt.Println("GOPATH:", os.Getenv("GOPATH"))

	// ─────────────────────────────────────────────
	// 2. File operations
	// ─────────────────────────────────────────────
	fmt.Println("\n-- File I/O --")

	// Write file:
	tmpFile := filepath.Join(os.TempDir(), "go_example.txt")
	err := os.WriteFile(tmpFile, []byte("Hello, Go!\nLine 2\n"), 0644)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
	fmt.Println("Wrote:", tmpFile)

	// Read file:
	data, err := os.ReadFile(tmpFile)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Printf("Read: %q\n", string(data))

	// ─────────────────────────────────────────────
	// 3. File with Open/Close
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Open/Close --")
	f, err := os.Open(tmpFile)
	if err != nil {
		fmt.Println("Open error:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 5)
	n, _ := f.Read(buf)
	fmt.Printf("Read %d bytes: %q\n", n, buf[:n])

	// ─────────────────────────────────────────────
	// 4. io utilities
	// ─────────────────────────────────────────────
	fmt.Println("\n-- io utilities --")

	// Seek back to start:
	f.Seek(0, io.SeekStart)

	// ReadAll:
	all, _ := io.ReadAll(f)
	fmt.Printf("ReadAll: %q\n", string(all))

	// ─────────────────────────────────────────────
	// 5. Directory operations
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Directories --")
	cwd, _ := os.Getwd()
	fmt.Println("CWD:", cwd)

	entries, _ := os.ReadDir(".")
	fmt.Printf("Files in '.': %d entries\n", len(entries))
	for i, e := range entries {
		if i >= 3 {
			fmt.Println("  ...")
			break
		}
		fmt.Printf("  %s (dir=%t)\n", e.Name(), e.IsDir())
	}

	// Cleanup:
	os.Remove(tmpFile)

	// ─────────────────────────────────────────────
	// 6. os.Args and os.Exit
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Process --")
	fmt.Println("Args:", os.Args)
	fmt.Println("PID:", os.Getpid())
	// os.Exit(1)  — exits immediately (defers DON'T run!)
}
