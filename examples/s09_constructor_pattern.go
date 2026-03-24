//go:build ignore

// Section 9, Topic 73: Constructor Function Pattern (NewXxx)
//
// Go has no constructors. By convention, use NewXxx functions
// to create and initialize structs, especially when:
//   - Fields need validation
//   - Maps/slices need initialization (avoid nil map panics)
//   - Default values should be set
//
// Convention: NewTypeName returns *TypeName
//
// Run: go run examples/s09_constructor_pattern.go

package main

import (
	"errors"
	"fmt"
)

// ─────────────────────────────────────────────
// Type with constructor
// ─────────────────────────────────────────────
type Server struct {
	Host     string
	Port     int
	handlers map[string]func()
}

// Constructor — initializes internal state
func NewServer(host string, port int) *Server {
	return &Server{
		Host:     host,
		Port:     port,
		handlers: make(map[string]func()), // avoids nil map panic
	}
}

func (s *Server) AddHandler(path string, handler func()) {
	s.handlers[path] = handler
}

func (s *Server) String() string {
	return fmt.Sprintf("Server(%s:%d, handlers=%d)", s.Host, s.Port, len(s.handlers))
}

// ─────────────────────────────────────────────
// Constructor with validation
// ─────────────────────────────────────────────
type User struct {
	Name  string
	Email string
	Age   int
}

func NewUser(name, email string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if age < 0 || age > 150 {
		return nil, errors.New("invalid age")
	}
	return &User{Name: name, Email: email, Age: age}, nil
}

// ─────────────────────────────────────────────
// Constructor with defaults
// ─────────────────────────────────────────────
type Config struct {
	Host    string
	Port    int
	Timeout int
	Debug   bool
}

func DefaultConfig() Config {
	return Config{
		Host:    "localhost",
		Port:    8080,
		Timeout: 30,
		Debug:   false,
	}
}

func main() {
	fmt.Println("=== Constructor Pattern ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic constructor
	// ─────────────────────────────────────────────
	srv := NewServer("localhost", 8080)
	srv.AddHandler("/", func() { fmt.Println("root") })
	srv.AddHandler("/api", func() { fmt.Println("api") })
	fmt.Println(srv)

	// Without constructor — nil map panic:
	// bad := Server{Host: "localhost", Port: 8080}
	// bad.AddHandler("/", func(){})  // PANIC: nil map

	// ─────────────────────────────────────────────
	// 2. Constructor with validation
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Validation --")
	u, err := NewUser("Alice", "alice@example.com", 30)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v\n", *u)
	}

	_, err = NewUser("", "bad@example.com", 30)
	fmt.Println("Empty name:", err)

	_, err = NewUser("Bob", "bob@example.com", -5)
	fmt.Println("Bad age:", err)

	// ─────────────────────────────────────────────
	// 3. Default config pattern
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Defaults --")
	cfg := DefaultConfig()
	fmt.Printf("Default: %+v\n", cfg)

	// Override specific fields:
	cfg.Port = 9090
	cfg.Debug = true
	fmt.Printf("Custom:  %+v\n", cfg)
}
