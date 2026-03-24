//go:build ignore

// Section 19, Topic 140: Functional Options Pattern
//
// Functional options provide clean, extensible configuration for constructors.
// Alternative to builder pattern or config structs with many optional fields.
//
// Pattern:
//   type Option func(*Config)
//   func WithPort(p int) Option { return func(c *Config) { c.Port = p } }
//   func NewServer(opts ...Option) *Server { ... }
//
// Run: go run examples/s19_functional_options.go

package main

import "fmt"

// ─────────────────────────────────────────────
// Server with functional options
// ─────────────────────────────────────────────
type Server struct {
	host    string
	port    int
	timeout int
	maxConn int
	tls     bool
}

type Option func(*Server)

func WithHost(host string) Option {
	return func(s *Server) { s.host = host }
}

func WithPort(port int) Option {
	return func(s *Server) { s.port = port }
}

func WithTimeout(seconds int) Option {
	return func(s *Server) { s.timeout = seconds }
}

func WithMaxConnections(n int) Option {
	return func(s *Server) { s.maxConn = n }
}

func WithTLS() Option {
	return func(s *Server) { s.tls = true }
}

func NewServer(opts ...Option) *Server {
	// Defaults:
	s := &Server{
		host:    "localhost",
		port:    8080,
		timeout: 30,
		maxConn: 100,
		tls:     false,
	}

	// Apply options:
	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) String() string {
	proto := "http"
	if s.tls {
		proto = "https"
	}
	return fmt.Sprintf("%s://%s:%d (timeout=%ds, maxConn=%d)",
		proto, s.host, s.port, s.timeout, s.maxConn)
}

func main() {
	fmt.Println("=== Functional Options ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. All defaults
	// ─────────────────────────────────────────────
	s1 := NewServer()
	fmt.Println("Defaults:", s1)

	// ─────────────────────────────────────────────
	// 2. Custom options
	// ─────────────────────────────────────────────
	s2 := NewServer(
		WithHost("0.0.0.0"),
		WithPort(443),
		WithTLS(),
		WithTimeout(60),
	)
	fmt.Println("Custom:", s2)

	// ─────────────────────────────────────────────
	// 3. Partial customization
	// ─────────────────────────────────────────────
	s3 := NewServer(WithPort(9000))
	fmt.Println("Just port:", s3)

	// ─────────────────────────────────────────────
	// Benefits:
	// ─────────────────────────────────────────────
	// - Clean API: NewServer() works with no args
	// - Extensible: add new options without breaking callers
	// - Self-documenting: WithTLS() is clear
	// - No nil/zero confusion: defaults are explicit
	//
	// Used by:
	// - google.golang.org/grpc
	// - go.uber.org/zap
	// - Many popular Go libraries
}
