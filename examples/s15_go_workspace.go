//go:build ignore

// Section 15, Topic 116: Go Workspace (go.work)
//
// Go workspaces (Go 1.18+) let you work on multiple modules simultaneously.
// Useful for multi-module repositories or developing a library alongside its consumer.
//
// Commands:
//   go work init ./module1 ./module2  — create go.work
//   go work use ./another-module      — add module to workspace
//   go work sync                      — sync workspace with modules
//
// GOTCHA: go.work should NOT be committed (it's for local development).
// GOTCHA: go.work replaces the need for `replace` directives in go.mod.
//
// Run: go run examples/s15_go_workspace.go

package main

import "fmt"

func main() {
	fmt.Println("=== Go Workspace (go.work) ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. go.work file
	// ─────────────────────────────────────────────
	fmt.Println("-- go.work --")
	fmt.Print(`
go 1.22

use (
    ./app
    ./lib
    ./shared
)
`)

	// ─────────────────────────────────────────────
	// 2. Example project structure
	// ─────────────────────────────────────────────
	fmt.Println("-- Structure --")
	fmt.Print(`
  mymonorepo/
  ├── go.work          ← workspace file
  ├── app/
  │   ├── go.mod       ← module: mycompany.com/app
  │   └── main.go
  ├── lib/
  │   ├── go.mod       ← module: mycompany.com/lib
  │   └── utils.go
  └── shared/
      ├── go.mod       ← module: mycompany.com/shared
      └── types.go
`)

	// ─────────────────────────────────────────────
	// 3. Without workspace (old way)
	// ─────────────────────────────────────────────
	// Had to use `replace` in go.mod:
	// replace mycompany.com/lib => ../lib
	// Messy and error-prone. Must remember to remove before publishing.

	// ─────────────────────────────────────────────
	// 4. Workspace commands
	// ─────────────────────────────────────────────
	fmt.Println("-- Commands --")
	fmt.Println("  go work init ./app ./lib    # create workspace")
	fmt.Println("  go work use ./new-module    # add module")
	fmt.Println("  go work sync                # sync dependencies")
	fmt.Println("  GOWORK=off go build ./...   # ignore workspace")

	// ─────────────────────────────────────────────
	// 5. Best practices
	// ─────────────────────────────────────────────
	// - Add go.work to .gitignore
	// - Each module should build independently
	// - Use workspace only for local development
	// - CI/CD should NOT use workspace (each module builds separately)
}
