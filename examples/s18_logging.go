//go:build ignore

// Section 18, Topic 136: log and slog Packages
//
// log:  Basic logging (Go 1.0+)
// slog: Structured logging (Go 1.21+) — recommended for new code
//
// GOTCHA: log.Fatal calls os.Exit(1) after logging.
// GOTCHA: log.Panic calls panic() after logging.
//
// Run: go run examples/s18_logging.go

package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	// ─────────────────────────────────────────────
	// 1. Basic log package
	// ─────────────────────────────────────────────
	log.Println("=== Logging ===")
	log.Println("Basic log message")
	log.Printf("Formatted: name=%s, age=%d\n", "Alice", 30)

	// Custom logger:
	logger := log.New(os.Stdout, "[APP] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Custom logger message")

	// ─────────────────────────────────────────────
	// 2. Log flags
	// ─────────────────────────────────────────────
	// log.Ldate        — date: 2024/03/15
	// log.Ltime        — time: 14:30:00
	// log.Lmicroseconds — microsecond resolution
	// log.Lshortfile   — file.go:123
	// log.Llongfile    — /full/path/file.go:123
	// log.LUTC         — UTC time

	// ─────────────────────────────────────────────
	// 3. slog — structured logging (Go 1.21+)
	// ─────────────────────────────────────────────
	log.Println("")
	slog.Info("Starting application",
		"version", "1.0.0",
		"port", 8080,
	)

	slog.Warn("Deprecated API called",
		"endpoint", "/api/v1/users",
		"replacement", "/api/v2/users",
	)

	slog.Error("Failed to connect",
		"host", "db.example.com",
		"error", "timeout",
	)

	// ─────────────────────────────────────────────
	// 4. slog with groups
	// ─────────────────────────────────────────────
	slog.Info("Request processed",
		slog.Group("request",
			"method", "GET",
			"path", "/api/users",
		),
		slog.Group("response",
			"status", 200,
			"duration_ms", 42,
		),
	)

	// ─────────────────────────────────────────────
	// 5. JSON handler
	// ─────────────────────────────────────────────
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("JSON log",
		"user", "alice",
		"action", "login",
	)

	// ─────────────────────────────────────────────
	// 6. Log levels
	// ─────────────────────────────────────────────
	// slog.LevelDebug  (-4)
	// slog.LevelInfo   (0)   — default
	// slog.LevelWarn   (4)
	// slog.LevelError  (8)

	// Set minimum level:
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn, // only warn and error
	})
	warnLogger := slog.New(textHandler)
	warnLogger.Info("This won't show") // below warn
	warnLogger.Warn("This will show")
	warnLogger.Error("This will also show")

	// ─────────────────────────────────────────────
	// log.Fatal and log.Panic:
	// ─────────────────────────────────────────────
	// log.Fatal("bye")  → logs, then os.Exit(1)
	// log.Panic("oh no") → logs, then panic()
}
