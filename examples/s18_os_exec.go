//go:build ignore

// Section 18, Topic 130: os and os/exec
//
// os: Operating system functions (files, env, process).
// os/exec: Run external commands.
//
// Run: go run examples/s18_os_exec.go

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("=== os and os/exec ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Environment variables
	// ─────────────────────────────────────────────
	fmt.Println("HOME:", os.Getenv("HOME"))
	fmt.Println("PATH exists:", os.Getenv("PATH") != "")

	// Set (for this process only):
	os.Setenv("MY_VAR", "hello")
	fmt.Println("MY_VAR:", os.Getenv("MY_VAR"))

	// ─────────────────────────────────────────────
	// 2. Working directory
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Working directory --")
	wd, _ := os.Getwd()
	fmt.Println("CWD:", wd)

	// ─────────────────────────────────────────────
	// 3. File operations
	// ─────────────────────────────────────────────
	fmt.Println("\n-- File operations --")

	// Create temp file:
	tmpFile := filepath.Join(os.TempDir(), "go_example.txt")
	err := os.WriteFile(tmpFile, []byte("Hello, Go!\n"), 0644)
	if err != nil {
		fmt.Println("Write error:", err)
	} else {
		fmt.Println("Wrote:", tmpFile)
	}

	// Read file:
	data, err := os.ReadFile(tmpFile)
	if err != nil {
		fmt.Println("Read error:", err)
	} else {
		fmt.Printf("Read: %q\n", string(data))
	}

	// File info:
	info, err := os.Stat(tmpFile)
	if err == nil {
		fmt.Printf("Size: %d bytes, Mode: %s\n", info.Size(), info.Mode())
	}

	// Remove:
	os.Remove(tmpFile)

	// Check if file exists:
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		fmt.Println("File removed successfully")
	}

	// ─────────────────────────────────────────────
	// 4. os.Args (command line arguments)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- os.Args --")
	fmt.Printf("Program: %s\n", os.Args[0])
	fmt.Printf("Args: %v\n", os.Args[1:])

	// ─────────────────────────────────────────────
	// 5. os/exec — run external commands
	// ─────────────────────────────────────────────
	fmt.Println("\n-- os/exec --")

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "echo", "Hello from exec!")
	} else {
		cmd = exec.Command("echo", "Hello from exec!")
	}

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Exec error:", err)
	} else {
		fmt.Printf("Output: %s", output)
	}

	// CombinedOutput captures both stdout and stderr:
	cmd2 := exec.Command("go", "version")
	out2, err := cmd2.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Go version: %s", out2)
	}
}
