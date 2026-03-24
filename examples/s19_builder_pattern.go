//go:build ignore

// Section 19, Topic 141: Builder Pattern
//
// Builder pattern constructs complex objects step by step.
// Less common in Go than functional options, but useful for fluent APIs.
//
// Run: go run examples/s19_builder_pattern.go

package main

import (
	"fmt"
	"strings"
)

// ─────────────────────────────────────────────
// Query Builder
// ─────────────────────────────────────────────
type QueryBuilder struct {
	table      string
	conditions []string
	orderBy    string
	limit      int
	columns    []string
}

func NewQuery(table string) *QueryBuilder {
	return &QueryBuilder{
		table:   table,
		columns: []string{"*"},
	}
}

func (q *QueryBuilder) Select(cols ...string) *QueryBuilder {
	q.columns = cols
	return q // return self for chaining
}

func (q *QueryBuilder) Where(cond string) *QueryBuilder {
	q.conditions = append(q.conditions, cond)
	return q
}

func (q *QueryBuilder) OrderBy(col string) *QueryBuilder {
	q.orderBy = col
	return q
}

func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	q.limit = n
	return q
}

func (q *QueryBuilder) Build() string {
	var sb strings.Builder
	sb.WriteString("SELECT ")
	sb.WriteString(strings.Join(q.columns, ", "))
	sb.WriteString(" FROM ")
	sb.WriteString(q.table)

	if len(q.conditions) > 0 {
		sb.WriteString(" WHERE ")
		sb.WriteString(strings.Join(q.conditions, " AND "))
	}
	if q.orderBy != "" {
		sb.WriteString(" ORDER BY ")
		sb.WriteString(q.orderBy)
	}
	if q.limit > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT %d", q.limit))
	}

	return sb.String()
}

func main() {
	fmt.Println("=== Builder Pattern ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Simple query
	// ─────────────────────────────────────────────
	q1 := NewQuery("users").Build()
	fmt.Println("Simple:", q1)

	// ─────────────────────────────────────────────
	// 2. Complex query (fluent chaining)
	// ─────────────────────────────────────────────
	q2 := NewQuery("users").
		Select("id", "name", "email").
		Where("age >= 18").
		Where("active = true").
		OrderBy("name ASC").
		Limit(10).
		Build()
	fmt.Println("Complex:", q2)

	// ─────────────────────────────────────────────
	// 3. Another example
	// ─────────────────────────────────────────────
	q3 := NewQuery("orders").
		Select("id", "total").
		Where("status = 'pending'").
		OrderBy("created_at DESC").
		Limit(5).
		Build()
	fmt.Println("Orders:", q3)

	// ─────────────────────────────────────────────
	// Builder vs Functional Options:
	// ─────────────────────────────────────────────
	// Builder: good for multi-step construction, fluent API
	// Functional Options: good for constructor configuration
}
