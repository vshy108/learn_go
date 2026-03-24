package learn_go_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"
)

// skipFiles lists examples that cannot run in automated tests:
// - network listeners / HTTP servers that block
// - files that deliberately demonstrate log.Fatal / os.Exit
var skipFiles = map[string]string{
	"s18_net_http.go":        "starts HTTP server (blocks)",
	"s19_middleware_pattern.go": "references http.ListenAndServe",
}

// slowFiles get a longer timeout (they use time.Sleep for demos)
var slowFiles = map[string]bool{
	"s13_goroutine_basics.go":   true,
	"s13_goroutine_patterns.go": true,
	"s13_waitgroup.go":          true,
	"s14_channel_basics.go":     true,
	"s14_context_cancellation.go": true,
	"s14_fan_out_fan_in.go":     true,
	"s14_select.go":             true,
	"s18_time.go":               true,
	"s19_concurrency_patterns.go": true,
}

// TestAllExamples discovers every s*_*.go file in examples/ and
// runs it with `go run`, verifying exit code 0 (compiles + no panic).
func TestAllExamples(t *testing.T) {
	files, err := filepath.Glob("examples/s*.go")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("no example files found — run from repo root")
	}

	sort.Strings(files)
	t.Logf("Discovered %d example files", len(files))

	for _, file := range files {
		base := filepath.Base(file)
		t.Run(base, func(t *testing.T) {
			t.Parallel()

			if reason, skip := skipFiles[base]; skip {
				t.Skipf("skipped: %s", reason)
			}

			timeout := 10 * time.Second
			if slowFiles[base] {
				timeout = 30 * time.Second
			}

			cmd := exec.Command("go", "run", file)
			cmd.Env = append(os.Environ(), "GOFLAGS=-buildvcs=false")

			done := make(chan error, 1)
			var output []byte
			go func() {
				var runErr error
				output, runErr = cmd.CombinedOutput()
				done <- runErr
			}()

			select {
			case runErr := <-done:
				if runErr != nil {
					t.Errorf("go run %s failed:\n%s\n%v", base, output, runErr)
				}
			case <-time.After(timeout):
				cmd.Process.Kill()
				t.Errorf("go run %s timed out after %v", base, timeout)
			}
		})
	}
}

// TestExamplesBySection runs examples grouped by section.
// Each section is a subtest so you can run: go test -run TestExamplesBySection/s05
func TestExamplesBySection(t *testing.T) {
	files, err := filepath.Glob("examples/s*.go")
	if err != nil {
		t.Fatal(err)
	}

	sections := make(map[string][]string)
	for _, f := range files {
		base := filepath.Base(f)
		// Extract section prefix: "s01", "s02", etc.
		parts := strings.SplitN(base, "_", 2)
		if len(parts) >= 1 {
			sections[parts[0]] = append(sections[parts[0]], f)
		}
	}

	sectionNames := make([]string, 0, len(sections))
	for s := range sections {
		sectionNames = append(sectionNames, s)
	}
	sort.Strings(sectionNames)

	sectionLabels := map[string]string{
		"s01": "Basics",
		"s02": "Variables_Constants",
		"s03": "Data_Types",
		"s04": "Functions",
		"s05": "Control_Flow",
		"s06": "Arrays_Slices",
		"s07": "Maps",
		"s08": "Strings",
		"s09": "Structs",
		"s10": "Pointers",
		"s11": "Interfaces",
		"s12": "Error_Handling",
		"s13": "Goroutines",
		"s14": "Channels",
		"s15": "Packages_Modules",
		"s16": "Generics",
		"s17": "Testing",
		"s18": "Stdlib",
		"s19": "Advanced_Patterns",
		"s20": "Reflection_Unsafe",
	}

	for _, sec := range sectionNames {
		secFiles := sections[sec]
		sort.Strings(secFiles)
		label := sec
		if l, ok := sectionLabels[sec]; ok {
			label = sec + "_" + l
		}

		t.Run(label, func(t *testing.T) {
			for _, file := range secFiles {
				base := filepath.Base(file)
				f := file // capture
				t.Run(base, func(t *testing.T) {
					t.Parallel()

					if reason, skip := skipFiles[base]; skip {
						t.Skipf("skipped: %s", reason)
					}

					timeout := 10 * time.Second
					if slowFiles[base] {
						timeout = 30 * time.Second
					}

					cmd := exec.Command("go", "run", f)
					cmd.Env = append(os.Environ(), "GOFLAGS=-buildvcs=false")

					done := make(chan error, 1)
					var output []byte
					go func() {
						var runErr error
						output, runErr = cmd.CombinedOutput()
						done <- runErr
					}()

					select {
					case runErr := <-done:
						if runErr != nil {
							t.Errorf("go run %s failed:\n%s\n%v", base, output, runErr)
						}
					case <-time.After(timeout):
						cmd.Process.Kill()
						t.Errorf("go run %s timed out after %v", base, timeout)
					}
				})
			}
		})
	}
}

// TestExamplesCompile verifies all example files compile without running them.
// Faster than TestAllExamples — use for quick CI checks.
func TestExamplesCompile(t *testing.T) {
	files, err := filepath.Glob("examples/s*.go")
	if err != nil {
		t.Fatal(err)
	}

	sort.Strings(files)

	for _, file := range files {
		base := filepath.Base(file)
		t.Run(base, func(t *testing.T) {
			t.Parallel()

			cmd := exec.Command("go", "build", "-o", os.DevNull, file)
			cmd.Env = append(os.Environ(), "GOFLAGS=-buildvcs=false")
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("go build %s failed:\n%s\n%v", base, output, err)
			}
		})
	}
}
