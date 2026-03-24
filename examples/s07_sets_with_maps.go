//go:build ignore

// Section 7, Topic 57: Implementing Sets with Maps
//
// Go has no built-in set type. Use map[T]bool or map[T]struct{}.
// map[T]struct{} uses zero memory for values (preferred).
//
// Run: go run examples/s07_sets_with_maps.go

package main

import "fmt"

// Set type using map[string]struct{}
type StringSet map[string]struct{}

func NewStringSet(items ...string) StringSet {
	s := make(StringSet)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func (s StringSet) Add(item string) {
	s[item] = struct{}{}
}

func (s StringSet) Remove(item string) {
	delete(s, item)
}

func (s StringSet) Contains(item string) bool {
	_, ok := s[item]
	return ok
}

func (s StringSet) Size() int {
	return len(s)
}

// Union: items in either set
func (s StringSet) Union(other StringSet) StringSet {
	result := NewStringSet()
	for k := range s {
		result.Add(k)
	}
	for k := range other {
		result.Add(k)
	}
	return result
}

// Intersection: items in both sets
func (s StringSet) Intersection(other StringSet) StringSet {
	result := NewStringSet()
	for k := range s {
		if other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

func main() {
	fmt.Println("=== Sets with Maps ===")
	fmt.Println()

	// 1. Create and use
	fruits := NewStringSet("apple", "banana", "cherry")
	fmt.Println("Contains apple:", fruits.Contains("apple"))
	fmt.Println("Contains grape:", fruits.Contains("grape"))
	fmt.Println("Size:", fruits.Size())

	// 2. Add and remove
	fruits.Add("grape")
	fruits.Remove("banana")
	fmt.Println("\nAfter add grape, remove banana:")
	for item := range fruits {
		fmt.Println(" ", item)
	}

	// 3. Set operations
	fmt.Println("\n-- Set operations --")
	a := NewStringSet("apple", "banana", "cherry")
	b := NewStringSet("banana", "date", "elderberry")

	union := a.Union(b)
	inter := a.Intersection(b)

	fmt.Print("A:            ")
	for k := range a {
		fmt.Print(k, " ")
	}
	fmt.Print("\nB:            ")
	for k := range b {
		fmt.Print(k, " ")
	}
	fmt.Print("\nUnion:        ")
	for k := range union {
		fmt.Print(k, " ")
	}
	fmt.Print("\nIntersection: ")
	for k := range inter {
		fmt.Print(k, " ")
	}
	fmt.Println()

	// 4. Quick set with map[T]bool (simpler but wastes 1 byte per entry)
	fmt.Println("\n-- map[T]bool alternative --")
	seen := map[int]bool{}
	nums := []int{1, 2, 3, 2, 1, 4}
	for _, n := range nums {
		seen[n] = true
	}
	fmt.Println("Unique count:", len(seen))
}
