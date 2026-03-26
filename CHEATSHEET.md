# Go Language Cheatsheet

> Quick reference for Go syntax, patterns, and gotchas.

---

## Table of Contents

1. [Basics](#1-basics)
2. [Variables & Constants](#2-variables--constants)
3. [Data Types](#3-data-types)
4. [Functions](#4-functions)
5. [Control Flow](#5-control-flow)
6. [Arrays & Slices](#6-arrays--slices)
7. [Maps](#7-maps)
8. [Strings](#8-strings)
9. [Structs & Methods](#9-structs--methods)
10. [Pointers](#10-pointers)
11. [Interfaces](#11-interfaces)
12. [Error Handling](#12-error-handling)
13. [Goroutines](#13-goroutines)
14. [Channels](#14-channels)
15. [Packages & Modules](#15-packages--modules)
16. [Generics](#16-generics-go-118)
17. [Testing](#17-testing)
18. [Standard Library Highlights](#18-standard-library-highlights)
19. [Advanced Patterns](#19-advanced-patterns)
20. [Reflection & Unsafe](#20-reflection--unsafe)

---

## 1. Basics

```go
// Every Go file starts with a package declaration
package main

// Grouped imports (preferred style)
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Hello, World!")
}
```

```sh
go run main.go       # compile + run
go build main.go     # compile to binary
go vet main.go       # static analysis
```

**Gotchas:**
- Brace `{` must be on the same line (semicolons are auto-inserted)
- Unused imports and variables are compile errors
- `gofmt` is the canonical formatter — use it

### fmt Verbs

| Verb | Description |
|------|-------------|
| `%v` | Default format |
| `%+v` | Struct with field names |
| `%#v` | Go syntax representation |
| `%T` | Type of the value |
| `%d` | Integer (base 10) |
| `%b` | Binary |
| `%x` | Hex (lowercase) |
| `%f` | Float (default precision) |
| `%e` | Scientific notation |
| `%s` | String |
| `%q` | Quoted string |
| `%p` | Pointer address |
| `%t` | Boolean |
| `%%` | Literal percent sign |

---

## 2. Variables & Constants

```go
// var declaration (explicit type)
var name string = "Go"
var age int          // zero value: 0

// Short declaration (type inferred, only inside functions)
count := 42
x, y := 10, 20

// Multiple declaration block
var (
    a int
    b string
    c bool
)

// Constants (must be compile-time evaluable)
const Pi = 3.14159
const (
    StatusOK    = 200
    StatusError = 500
)

// iota — auto-incrementing constant generator (resets per const block)
const (
    Red   = iota // 0
    Green        // 1
    Blue         // 2
)

// Bitmask with iota
const (
    Read    = 1 << iota // 1
    Write               // 2
    Execute             // 4
)

// Blank identifier — discard unwanted values
_, err := someFunc()

// Swap without temp
a, b = b, a
```

### Zero Values

| Type | Zero Value |
|------|-----------|
| `int`, `float64` | `0` |
| `bool` | `false` |
| `string` | `""` |
| Pointer, slice, map, channel, func, interface | `nil` |

**Gotchas:**
- `:=` only works inside functions
- `:=` can shadow outer variables in inner blocks
- Constants can be untyped — they have higher precision and adapt to context
- `iota` resets to 0 at each `const` block

---

## 3. Data Types

### Numeric Types

| Type | Size | Range |
|------|------|-------|
| `int8` | 1 byte | -128 to 127 |
| `int16` | 2 bytes | -32,768 to 32,767 |
| `int32` | 4 bytes | -2B to 2B |
| `int64` | 8 bytes | -9.2×10¹⁸ to 9.2×10¹⁸ |
| `uint8` (`byte`) | 1 byte | 0 to 255 |
| `uint16` | 2 bytes | 0 to 65,535 |
| `uint32` | 4 bytes | 0 to 4B |
| `uint64` | 8 bytes | 0 to 18.4×10¹⁸ |
| `int` | platform | 32 or 64 bit |
| `float32` | 4 bytes | ~7 decimal digits |
| `float64` | 8 bytes | ~15 decimal digits |
| `complex64` | 8 bytes | float32 real + imag |
| `complex128` | 16 bytes | float64 real + imag |

```go
// byte = uint8, rune = int32 (Unicode code point)
var b byte = 'A'      // 65
var r rune = '世'      // 19990

// Type conversion — MUST be explicit (no implicit casting)
var i int = 42
var f float64 = float64(i)
var u uint8 = uint8(f)

// Type alias vs defined type
type Celsius float64    // defined type — new type, no methods inherited
type MyInt = int        // alias — same type, interchangeable
```

**Gotchas:**
- Integer overflow wraps silently (no runtime panic)
- `int` is NOT an alias for `int64` — they are distinct types
- Converting `float64` to `int` truncates (no rounding)
- `byte` and `uint8` are interchangeable; `rune` and `int32` are interchangeable

---

## 4. Functions

```go
// Basic function
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values (naked return)
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // returns x and y
}

// Variadic function
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
// Call: sum(1, 2, 3) or sum(slice...)

// First-class functions
var op func(int, int) int = add

// Anonymous function / closure
increment := func() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}()

// Defer — executes in LIFO order when function returns
func readFile() {
    f, _ := os.Open("file.txt")
    defer f.Close() // runs when readFile() exits
    // ... use f
}

// init() — runs before main(), once per package, can have multiple per file
func init() {
    // setup code
}
```

**Gotchas:**
- Named returns: `defer` can modify return values
- Defer args are evaluated immediately, but the call is deferred
- `defer` in a loop creates one deferred call per iteration (can leak resources)
- Go has **no tail-call optimization** — deep recursion can stack overflow
- `init()` cannot be called explicitly

---

## 5. Control Flow

```go
// if — no parentheses, braces required
if x > 0 {
    fmt.Println("positive")
} else if x == 0 {
    fmt.Println("zero")
} else {
    fmt.Println("negative")
}

// if with initializer (scoped to if/else block)
if err := doSomething(); err != nil {
    log.Fatal(err)
}

// for — the only loop keyword in Go
for i := 0; i < 10; i++ { }     // C-style
for i < 10 { }                   // while-style
for { }                           // infinite loop

// for range
for i, v := range slice { }      // index + value
for _, v := range slice { }      // value only
for i := range slice { }         // index only
for range 10 { }                 // Go 1.22+: repeat N times

// switch — no fallthrough by default, no break needed
switch day {
case "Mon", "Tue", "Wed", "Thu", "Fri":
    fmt.Println("weekday")
case "Sat", "Sun":
    fmt.Println("weekend")
default:
    fmt.Println("unknown")
}

// Switch without condition (cleaner than if/else chains)
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
default:
    grade = "C"
}

// Type switch
switch v := i.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Printf("unknown: %T\n", v)
}

// Labels
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                break outer // breaks outer loop
            }
        }
    }
```

**Gotchas:**
- No ternary operator in Go
- `fallthrough` keyword exists but is rarely used — it falls into the next case unconditionally
- `for range` over a map has random iteration order
- `for range` over a string iterates runes, not bytes

---

## 6. Arrays & Slices

```go
// Array — fixed size, value type (copied on assignment)
var arr [5]int                     // [0 0 0 0 0]
arr := [3]int{1, 2, 3}
arr := [...]int{1, 2, 3}          // size inferred

// Slice — dynamic, reference type (backed by array)
s := []int{1, 2, 3}               // slice literal
s := make([]int, 5)               // len=5, cap=5
s := make([]int, 0, 10)           // len=0, cap=10

// Slicing (half-open interval: includes low, excludes high)
sub := s[1:4]                     // elements at index 1, 2, 3
sub := s[:3]                      // first 3
sub := s[2:]                      // from index 2 to end
sub := s[:]                       // copy of the whole slice header

// Full slice expression (controls capacity)
sub := s[1:3:5]                   // len=2, cap=4

// append — may allocate new backing array
s = append(s, 4, 5)
s = append(s, other...)           // append another slice

// copy
n := copy(dst, src)               // returns number of elements copied

// Delete element at index i (preserves order)
s = append(s[:i], s[i+1:]...)     // Go <1.21
s = slices.Delete(s, i, i+1)     // Go 1.21+

// Insert at index i
s = slices.Insert(s, i, elem)    // Go 1.21+

// Filter in-place
n := 0
for _, v := range s {
    if keep(v) {
        s[n] = v
        n++
    }
}
s = s[:n]
```

### Slice Internals

```
┌─────────┐
│  ptr ────┼──▶ backing array [0][1][2][3][4][5]
│  len: 3  │         ▲
│  cap: 6  │         │
└─────────┘     shared memory!
```

**Gotchas:**
- `nil` slice vs empty slice: `var s []int` (nil) vs `s := []int{}` (empty). Both have len=0, but `nil` slice == nil
- Slices from the same array share memory — mutations are visible
- `append` may or may not create a new backing array — never assume
- Growth strategy roughly doubles capacity (implementation detail, don't depend on it)

---

## 7. Maps

```go
// Create
m := map[string]int{"a": 1, "b": 2}
m := make(map[string]int)          // empty map
m := make(map[string]int, 100)     // hint capacity

// Access with comma-ok idiom
val, ok := m["key"]
if !ok {
    fmt.Println("key not found")
}

// Add / Update
m["key"] = 42

// Delete
delete(m, "key")                   // no-op if key doesn't exist

// Iteration (random order!)
for k, v := range m {
    fmt.Println(k, v)
}

// Check existence only
_, exists := m["key"]

// Set as map (no built-in set type)
set := map[string]struct{}{}
set["item"] = struct{}{}
_, exists := set["item"]
delete(set, "item")

// Map of slices
graph := map[string][]string{}
graph["a"] = append(graph["a"], "b")
```

**Gotchas:**
- `nil` map: reads return zero value, **writes panic**
- Map keys must be comparable (`==`); slices, maps, and functions cannot be keys
- Iteration order is randomized by design
- Maps are not safe for concurrent access — use `sync.Map` or `sync.RWMutex`
- You cannot take the address of a map element (`&m["key"]` is illegal)

---

## 8. Strings

```go
// Strings are immutable UTF-8 byte sequences
s := "Hello, 世界"
len(s)            // 13 (bytes, not characters!)

// Rune iteration
for i, r := range s {
    fmt.Printf("%d: %c (%U)\n", i, r, r)
}

// Byte iteration
for i := 0; i < len(s); i++ {
    fmt.Printf("%d: %x\n", i, s[i])
}

// Rune count
utf8.RuneCountInString(s)   // 9

// Raw strings (no escape processing)
raw := `line1\nstill line1`

// strings package
strings.Contains(s, "sub")
strings.HasPrefix(s, "He")
strings.HasSuffix(s, "lo")
strings.Split("a,b,c", ",")       // []string{"a","b","c"}
strings.Join([]string{"a","b"}, "-") // "a-b"
strings.ToUpper(s)
strings.ToLower(s)
strings.TrimSpace("  hi  ")       // "hi"
strings.Replace(s, "old", "new", -1)
strings.Count(s, "l")
strings.Index(s, "lo")            // -1 if not found
strings.Repeat("ab", 3)           // "ababab"

// Efficient concatenation
var b strings.Builder
for i := 0; i < 1000; i++ {
    b.WriteString("x")
}
result := b.String()

// strconv conversions
i, err := strconv.Atoi("42")       // string → int
s := strconv.Itoa(42)              // int → string
f, err := strconv.ParseFloat("3.14", 64)
s := strconv.FormatFloat(3.14, 'f', 2, 64)
b, err := strconv.ParseBool("true")
```

**Gotchas:**
- `len(s)` returns byte count, not character/rune count
- String indexing `s[i]` returns a byte, not a rune
- Strings are immutable — to modify, convert to `[]byte` or `[]rune`
- `+` concatenation in a loop is O(n²) — use `strings.Builder`

---

## 9. Structs & Methods

```go
// Define
type User struct {
    Name  string
    Email string
    Age   int
}

// Create
u := User{Name: "Alice", Email: "a@b.com", Age: 30}
u := User{"Alice", "a@b.com", 30}  // positional (fragile, avoid)
u := new(User)                       // returns *User (zero values)

// Access
u.Name = "Bob"

// Method with value receiver (operates on a copy)
func (u User) FullName() string {
    return u.Name
}

// Method with pointer receiver (can modify the struct)
func (u *User) SetName(name string) {
    u.Name = name
}

// Embedded struct (composition)
type Admin struct {
    User               // embedded — Admin gets all User fields/methods
    Level int
}
a := Admin{User: User{Name: "Alice"}, Level: 1}
a.Name   // promoted field access

// Anonymous struct
config := struct {
    Host string
    Port int
}{Host: "localhost", Port: 8080}

// Struct tags (metadata for JSON, DB, validation, etc.)
type Person struct {
    Name string `json:"name" validate:"required"`
    Age  int    `json:"age,omitempty"`
}

// Constructor pattern
func NewUser(name, email string) *User {
    return &User{Name: name, Email: email}
}
```

**Gotchas:**
- Value receiver: cannot modify the original struct
- Pointer receiver: can modify the original, and the method set includes both value and pointer methods
- Structs are comparable with `==` only if all fields are comparable
- Embedding is NOT inheritance — there's no polymorphism with embedded types
- Promoted methods can be shadowed by the outer struct's methods

---

## 10. Pointers

```go
x := 42
p := &x          // p is *int, points to x
fmt.Println(*p)  // 42 (dereference)
*p = 100         // x is now 100

// new() allocates zeroed memory, returns pointer
p := new(int)    // *int, value is 0

// nil pointer
var p *int       // nil
// *p would panic: nil pointer dereference

// Pointer to struct — auto-dereference
type Point struct{ X, Y int }
p := &Point{1, 2}
p.X = 10         // same as (*p).X = 10
```

### When to Use Pointer vs Value Receiver

| Use Pointer Receiver | Use Value Receiver |
|----------------------|--------------------|
| Need to mutate the receiver | Read-only access |
| Large struct (avoid copy) | Small struct (int, small structs) |
| Consistency (if any method uses pointer, all should) | Immutable value semantics |

**Gotchas:**
- Go has **no pointer arithmetic** (unlike C)
- Go functions are pass-by-value — pointers let you modify the original
- Returning a pointer to a local variable is safe (Go does escape analysis)
- `nil` pointer dereference causes a runtime panic

---

## 11. Interfaces

```go
// Define — a set of method signatures
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Composition
type ReadWriter interface {
    Reader
    Writer
}

// Implicit implementation — no "implements" keyword
type MyReader struct{}
func (r MyReader) Read(p []byte) (int, error) {
    return 0, nil
}
// MyReader now satisfies the Reader interface

// Empty interface — accepts any type
var x any          // same as interface{}
x = 42
x = "hello"

// Type assertion
s, ok := x.(string)
if ok {
    fmt.Println(s)
}

// Type switch
switch v := x.(type) {
case int:
    fmt.Println("int", v)
case string:
    fmt.Println("string", v)
}

// Common interfaces
// fmt.Stringer      → String() string
// error             → Error() string
// io.Reader         → Read([]byte) (int, error)
// io.Writer         → Write([]byte) (int, error)
// io.Closer         → Close() error
// sort.Interface    → Len(), Less(i,j), Swap(i,j)
```

**Gotchas:**
- Interfaces are satisfied implicitly — great for decoupling, tricky for discoverability
- `nil` interface vs interface holding `nil`:
  ```go
  var e error            // nil interface (type=nil, value=nil) → e == nil ✓
  var p *MyError = nil
  e = p                  // non-nil interface (type=*MyError, value=nil) → e != nil !
  ```
- Accept interfaces, return structs (common Go idiom)
- Keep interfaces small — prefer 1-2 methods

---

## 12. Error Handling

```go
// errors.New
err := errors.New("something failed")

// fmt.Errorf with formatting
err := fmt.Errorf("failed to load %s: %w", name, origErr)

// Custom error type
type NotFoundError struct {
    Name string
}
func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s not found", e.Name)
}

// Wrapping and unwrapping
err := fmt.Errorf("operation failed: %w", origErr)  // wrap
inner := errors.Unwrap(err)                           // unwrap

// errors.Is — check error chain for specific value
if errors.Is(err, os.ErrNotExist) { }

// errors.As — check error chain for specific type
var nfe *NotFoundError
if errors.As(err, &nfe) {
    fmt.Println(nfe.Name)
}

// Sentinel errors (package-level, compare with errors.Is)
var ErrNotFound = errors.New("not found")

// panic / recover — for truly unrecoverable situations
func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    return a / b, nil
}
```

**Gotchas:**
- Always check errors — `val, err := f(); if err != nil { ... }`
- Use `%w` (not `%v`) to wrap errors so `errors.Is`/`errors.As` work
- `panic` should be rare — prefer returning errors
- `recover` only works inside a `defer` function
- Sentinel errors should be `var`, not `const` (errors aren't compile-time constants)

---

## 13. Goroutines

```go
// Start a goroutine
go func() {
    fmt.Println("running concurrently")
}()

// WaitGroup — wait for goroutines to finish
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println(id)
    }(i) // pass i to avoid closure capture bug
}
wg.Wait()

// Mutex — protect shared state
var mu sync.Mutex
var count int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
}

// RWMutex — multiple readers, single writer
var rw sync.RWMutex
rw.RLock()   // shared read lock
rw.RUnlock()
rw.Lock()    // exclusive write lock
rw.Unlock()

// sync.Once — execute exactly once
var once sync.Once
once.Do(func() {
    // initialization
})

// sync/atomic — lock-free operations
var counter int64
atomic.AddInt64(&counter, 1)
val := atomic.LoadInt64(&counter)
```

**Gotchas:**
- Goroutines are cheap (~2KB stack) but not free — can leak if not properly managed
- Closure capture in loops: always pass loop variable as argument or use `i := i`
- Race detector: `go run -race main.go` (use it!)
- `runtime.GOMAXPROCS(n)` — defaults to number of CPUs since Go 1.5

---

## 14. Channels

```go
// Unbuffered (synchronous — sender blocks until receiver is ready)
ch := make(chan int)

// Buffered (async up to capacity)
ch := make(chan int, 10)

// Send and receive
ch <- 42       // send
val := <-ch    // receive

// Channel direction (in function signatures)
func producer(out chan<- int) { out <- 1 }     // send-only
func consumer(in <-chan int)  { val := <-in }  // receive-only

// Close and range
close(ch)
for val := range ch {
    fmt.Println(val)
}

// Check if channel is closed
val, ok := <-ch   // ok is false if closed and empty

// select — multiplex channels
select {
case msg := <-ch1:
    fmt.Println(msg)
case ch2 <- 42:
    fmt.Println("sent")
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no channel ready")
}

// Done channel pattern
done := make(chan struct{})
go func() {
    defer close(done)
    // work...
}()
<-done // wait for completion

// Fan-out / fan-in
// Fan-out: multiple goroutines read from one channel
// Fan-in:  multiple channels merge into one

// Pipeline
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Context cancellation
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
select {
case result := <-doWork(ctx):
    fmt.Println(result)
case <-ctx.Done():
    fmt.Println("cancelled:", ctx.Err())
}
```

**Gotchas:**
- Sending to a closed channel **panics**
- Receiving from a closed channel returns the zero value immediately
- Sending/receiving on a `nil` channel blocks forever
- Always close channels from the sender side
- Buffered channel: sends block only when full; receives block only when empty

---

## 15. Packages & Modules

```sh
go mod init github.com/user/project   # initialize module
go mod tidy                            # sync dependencies
go get github.com/pkg/errors           # add dependency
go list -m all                         # list all dependencies
go mod vendor                          # vendoring
```

```go
// Exported = uppercase first letter
func PublicFunc() {}    // exported (accessible outside package)
func privateFunc() {}   // unexported (internal only)

// Import aliases
import (
    "fmt"
    myfmt "mypackage/fmt"   // alias
    . "math"                 // dot import (avoid — pollutes namespace)
    _ "image/png"            // blank import (side effects only, runs init())
)

// internal/ packages — only accessible by parent and sibling packages
// project/internal/helper   → only project/* can import it
```

**Gotchas:**
- `init()` runs in dependency order, then alphabetical file order within a package
- Circular imports are compile errors
- `go.sum` contains cryptographic hashes for module verification — commit it
- `replace` directive in `go.mod` for local development overrides

---

## 16. Generics (Go 1.18+)

```go
// Generic function
func Map[T, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}

// Type constraint
type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~float32 | ~float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

// Generic type
type Stack[T any] struct {
    items []T
}
func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    v := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return v, true
}

// Built-in constraints
// any        — alias for interface{}, accepts all types
// comparable — types that support == and != (usable as map keys)
```

**Gotchas:**
- `~int` means "any type whose underlying type is int" (includes defined types)
- Generic methods on generic types cannot introduce new type parameters
- Prefer concrete types when generics aren't needed — simpler is better
- No generic methods (only generic functions and generic types)

---

## 17. Testing

```go
// File: xxx_test.go (must end in _test.go)
package mypackage

import "testing"

// Unit test
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2,3) = %d, want %d", got, want)
    }
}

// Table-driven test
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 1, 2, 3},
        {"zero", 0, 0, 0},
        {"negative", -1, -1, -2},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.expected {
                t.Errorf("got %d, want %d", got, tt.expected)
            }
        })
    }
}

// Test helper
func assertEqual(t *testing.T, got, want int) {
    t.Helper() // marks this as helper (error points to caller)
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

// Benchmark
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}

// Example test (appears in godoc)
func ExampleAdd() {
    fmt.Println(Add(1, 2))
    // Output: 3
}

// Fuzz test (Go 1.18+)
func FuzzAdd(f *testing.F) {
    f.Add(1, 2)
    f.Fuzz(func(t *testing.T, a, b int) {
        Add(a, b) // just check it doesn't panic
    })
}

// TestMain — setup/teardown for entire package
func TestMain(m *testing.M) {
    // setup
    code := m.Run()
    // teardown
    os.Exit(code)
}
```

```sh
go test ./...                      # run all tests
go test -v ./...                   # verbose
go test -run TestAdd ./...         # run specific test
go test -count=1 ./...             # bypass cache
go test -cover ./...               # coverage
go test -bench=. ./...             # benchmarks
go test -fuzz=FuzzAdd ./...        # fuzzing
go test -race ./...                # race detector
```

---

## 18. Standard Library Highlights

### fmt

```go
fmt.Println("Hello")                    // + newline
fmt.Printf("Name: %s, Age: %d\n", n, a) // formatted
s := fmt.Sprintf("value: %v", x)        // returns string
fmt.Fprintf(w, "template: %s", data)     // writes to io.Writer
```

### os & io

```go
// Read file
data, err := os.ReadFile("file.txt")

// Write file
os.WriteFile("out.txt", []byte("data"), 0644)

// Environment
val := os.Getenv("HOME")
os.Setenv("KEY", "value")

// Args
args := os.Args[1:] // command-line arguments (skip program name)

// io.Copy
io.Copy(dst, src)    // copies from Reader to Writer

// bufio.Scanner (line-by-line reading)
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    line := scanner.Text()
}
```

### encoding/json

```go
// Marshal (struct → JSON)
data, err := json.Marshal(user)
data, err := json.MarshalIndent(user, "", "  ")

// Unmarshal (JSON → struct)
var user User
err := json.Unmarshal(data, &user)

// Streaming with Encoder/Decoder
json.NewEncoder(w).Encode(user)      // write to io.Writer
json.NewDecoder(r).Decode(&user)     // read from io.Reader

// Struct tags control JSON field names
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
    Age   int    `json:"-"`              // skip this field
}
```

### net/http

```go
// Simple server
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
})
http.ListenAndServe(":8080", nil)

// HTTP client
resp, err := http.Get("https://api.example.com/data")
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
```

### time

```go
now := time.Now()
later := now.Add(2 * time.Hour)
diff := later.Sub(now)             // Duration
time.Sleep(100 * time.Millisecond)

// Formatting — uses reference time: Mon Jan 2 15:04:05 MST 2006
s := now.Format("2006-01-02 15:04:05")
t, err := time.Parse("2006-01-02", "2024-03-15")

// Ticker and Timer
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
for t := range ticker.C {
    fmt.Println("tick", t)
}
```

### context

```go
// WithCancel
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// WithTimeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// WithValue (use sparingly)
ctx = context.WithValue(ctx, "key", "value")
val := ctx.Value("key").(string)

// Check cancellation
select {
case <-ctx.Done():
    return ctx.Err()
default:
    // continue work
}
```

### sort

```go
sort.Ints(nums)
sort.Strings(strs)
sort.Float64s(floats)

// Custom sort
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})

// Check if sorted
sort.IntsAreSorted(nums)

// Binary search (Go 1.19+ slices package)
i, found := slices.BinarySearch(sorted, target)
```

### regexp

```go
re := regexp.MustCompile(`\d+`)
re.MatchString("abc123")          // true
re.FindString("abc123def456")     // "123"
re.FindAllString("a1b2c3", -1)   // ["1","2","3"]
re.ReplaceAllString("a1b2", "X") // "aXbX"
```

---

## 19. Advanced Patterns

### Functional Options

```go
type Server struct {
    port    int
    timeout time.Duration
}

type Option func(*Server)

func WithPort(p int) Option {
    return func(s *Server) { s.port = p }
}

func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.timeout = d }
}

func NewServer(opts ...Option) *Server {
    s := &Server{port: 8080, timeout: 30 * time.Second}
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage: NewServer(WithPort(9090), WithTimeout(10*time.Second))
```

### Worker Pool

```go
func workerPool(jobs <-chan int, results chan<- int, numWorkers int) {
    var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                results <- process(job)
            }
        }()
    }
    wg.Wait()
    close(results)
}
```

### Middleware Pattern (HTTP)

```go
func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
// Usage: http.Handle("/", logging(myHandler))
```

### Builder Pattern

```go
type QueryBuilder struct {
    table  string
    wheres []string
    limit  int
}

func (qb *QueryBuilder) From(table string) *QueryBuilder {
    qb.table = table
    return qb
}

func (qb *QueryBuilder) Where(cond string) *QueryBuilder {
    qb.wheres = append(qb.wheres, cond)
    return qb
}

func (qb *QueryBuilder) Limit(n int) *QueryBuilder {
    qb.limit = n
    return qb
}
```

---

## 20. Reflection & Unsafe

```go
import "reflect"

// Get type and value
t := reflect.TypeOf(x)      // reflect.Type
v := reflect.ValueOf(x)     // reflect.Value

// Inspect struct fields
t := reflect.TypeOf(User{})
for i := 0; i < t.NumField(); i++ {
    f := t.Field(i)
    fmt.Println(f.Name, f.Type, f.Tag.Get("json"))
}

// Modify via reflection (must pass pointer)
v := reflect.ValueOf(&x).Elem()
v.SetInt(42)

// unsafe (skip unless you know what you're doing)
import "unsafe"

size := unsafe.Sizeof(x)       // size in bytes
p := unsafe.Pointer(&x)        // convert any pointer to unsafe.Pointer
```

**Gotchas:**
- Reflection is slow — avoid in hot paths
- `reflect.ValueOf(x)` with non-pointer is read-only; use `reflect.ValueOf(&x).Elem()` to modify
- `unsafe` bypasses Go's type system — can cause memory corruption
- Use `unsafe` only for FFI, performance-critical code, or low-level libraries

---

## Quick Reference: Common Commands

```sh
go run main.go                # compile and run
go build -o app .             # build binary
go test ./...                 # run all tests
go test -v -race ./...        # verbose with race detector
go test -bench=. ./...        # run benchmarks
go test -cover ./...          # test coverage
go vet ./...                  # static analysis
go fmt ./...                  # format code
go mod init module/path       # init module
go mod tidy                   # sync dependencies
go doc fmt.Println            # view documentation
go generate ./...             # run go:generate directives
go install ./...              # compile and install
```

## Quick Reference: Common Idioms

```go
// Error check
if err != nil {
    return fmt.Errorf("context: %w", err)
}

// Comma-ok pattern
val, ok := m[key]
val, ok := x.(Type)
val, ok := <-ch

// Defer close
resp, err := http.Get(url)
if err != nil { return err }
defer resp.Body.Close()

// Goroutine with WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()

// Type assertion with check
if s, ok := val.(string); ok {
    // use s
}

// Ensure interface compliance at compile time
var _ Interface = (*MyType)(nil)
```
