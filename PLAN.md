# Go Learning Roadmap

> Track progress by checking off topics as they are implemented.
> Run any example: `go run examples/s01_hello_world.go`

## Workspace Layout

```
learn_go/
├── go.mod                   # Module definition
├── PLAN.md                  # This roadmap
├── README.md                # Quick-start guide
└── examples/
    ├── s01_*.go             # Section 1 topics (Basics)
    ├── s02_*.go             # Section 2 topics (Variables & Constants)
    ├── ...                  # Sections 3-20
    └── s20_*.go
```

## Conventions

- **File naming**: `examples/s{NN}_{topic}.go` — each file is standalone with `package main`
- **Build tag**: every file starts with `//go:build ignore` so multiple `main()` functions
  don't conflict. Files are run individually with `go run`.
- **No consolidation**: every topic = one file, independently runnable via `go run`
- **Detailed comments**: each file has a header explaining the concept, edge cases, and gotchas
- **Comparison notes**: where relevant, compare with Rust/Python equivalents

## Verification (per topic)

```sh
go run examples/<name>.go        # runs and produces expected output
go vet examples/<name>.go        # static analysis
go build examples/<name>.go      # compiles without warnings
```

---

## Section 1: Basics

- [x] 1. `s01_hello_world` — Hello World, package main, func main
- [x] 2. `s01_packages_imports` — Packages, imports, grouped imports
- [x] 3. `s01_fmt_verbs` — fmt package: Println, Printf, Sprintf, format verbs
- [x] 4. `s01_comments` — Single-line, multi-line, and doc comments
- [x] 5. `s01_go_run_build` — go run, go build, go install
- [x] 6. `s01_semicolons_braces` — Implicit semicolons and brace placement rules

## Section 2: Variables and Constants

- [x] 7. `s02_var_declaration` — var keyword, explicit types, zero values
- [x] 8. `s02_short_declaration` — Short variable declaration (:=)
- [x] 9. `s02_zero_values` — Zero values for all types
- [x] 10. `s02_type_inference` — Type inference rules
- [x] 11. `s02_constants` — const keyword, untyped constants
- [x] 12. `s02_iota` — iota enumerator
- [x] 13. `s02_multiple_assignment` — Multiple variable assignment and swapping
- [x] 14. `s02_scope_and_blocks` — Variable scope and block scoping
- [x] 15. `s02_blank_identifier` — The blank identifier (_)

## Section 3: Data Types

- [x] 16. `s03_basic_types` — int, float64, string, bool overview
- [x] 17. `s03_integer_types` — int8/16/32/64, uint8/16/32/64, int, uint, uintptr
- [x] 18. `s03_float_types` — float32, float64, precision issues
- [x] 19. `s03_complex_numbers` — complex64, complex128
- [x] 20. `s03_byte_and_rune` — byte (uint8) vs rune (int32), Unicode
- [x] 21. `s03_type_conversion` — Explicit type conversion (no implicit casting)
- [x] 22. `s03_string_conversions` — strconv: Atoi, Itoa, ParseFloat, FormatFloat
- [x] 23. `s03_overflow` — Integer overflow behavior (wraps silently!)
- [x] 24. `s03_type_aliases_defined` — Type aliases vs defined types
- [x] 25. `s03_comparison_operators` — Comparison and equality

## Section 4: Functions

- [x] 26. `s04_basic_functions` — Function declaration, parameters, return values
- [x] 27. `s04_multiple_returns` — Multiple return values
- [x] 28. `s04_named_returns` — Named return values (naked return)
- [x] 29. `s04_variadic_functions` — Variadic functions (...)
- [x] 30. `s04_first_class_functions` — Functions as values and types
- [x] 31. `s04_anonymous_closures` — Anonymous functions and closures
- [x] 32. `s04_defer` — defer keyword, LIFO order, and gotchas
- [x] 33. `s04_init_function` — init() function special behavior
- [x] 34. `s04_recursion` — Recursion and tail-call (no TCO in Go)

## Section 5: Control Flow

- [x] 35. `s05_if_else` — if/else, initializer statement, no ternary
- [x] 36. `s05_for_loops` — Three forms of for loop (Go has no while/do-while)
- [x] 37. `s05_for_range` — for range over slices, maps, strings, channels
- [x] 38. `s05_switch` — switch statement, no fallthrough by default
- [x] 39. `s05_type_switch` — Type switch for interface values
- [x] 40. `s05_labels_break_continue` — Labels, break, continue, goto
- [x] 41. `s05_infinite_loop` — Infinite loops and breaking out

## Section 6: Arrays and Slices

- [x] 42. `s06_arrays` — Array declaration, fixed size, value type
- [x] 43. `s06_slice_basics` — Slice creation, len, cap
- [x] 44. `s06_make_slice` — make() for slices
- [x] 45. `s06_append` — append() behavior and growth strategy
- [x] 46. `s06_copy` — copy() function
- [x] 47. `s06_nil_vs_empty_slice` — nil slice vs empty slice vs zero-length slice
- [x] 48. `s06_slice_internals` — Slice header (ptr, len, cap), shared backing array
- [x] 49. `s06_slice_tricks` — Common slice tricks (delete, insert, filter)
- [x] 50. `s06_multidimensional` — Multi-dimensional slices

## Section 7: Maps

- [x] 51. `s07_create_map` — Creating maps with make and literals
- [x] 52. `s07_access_comma_ok` — Accessing values, comma-ok idiom
- [x] 53. `s07_add_delete` — Adding and deleting entries
- [x] 54. `s07_iteration` — Iterating maps (random order!)
- [x] 55. `s07_nil_map` — nil map gotcha (reads ok, writes panic)
- [x] 56. `s07_map_of_slices` — Maps with slice values
- [x] 57. `s07_sets_with_maps` — Implementing sets with maps

## Section 8: Strings

- [x] 58. `s08_string_basics` — Strings are immutable byte slices (UTF-8)
- [x] 59. `s08_runes_vs_bytes` — Rune iteration vs byte iteration
- [x] 60. `s08_strings_package` — strings package methods (Contains, Split, Join, etc.)
- [x] 61. `s08_string_builder` — strings.Builder for efficient concatenation
- [x] 62. `s08_strconv` — strconv package conversions
- [x] 63. `s08_raw_strings` — Raw string literals (backtick)
- [x] 64. `s08_string_mutability` — Strings are immutable, []byte for mutation

## Section 9: Structs

- [x] 65. `s09_defining_structs` — Defining and creating structs
- [x] 66. `s09_accessing_fields` — Accessing and modifying fields
- [x] 67. `s09_methods_value_receiver` — Methods with value receiver
- [x] 68. `s09_methods_pointer_receiver` — Methods with pointer receiver
- [x] 69. `s09_embedded_structs` — Embedded structs (composition over inheritance)
- [x] 70. `s09_struct_tags` — Struct tags (JSON, DB, validation)
- [x] 71. `s09_anonymous_structs` — Anonymous structs
- [x] 72. `s09_struct_comparison` — Struct comparison and equality
- [x] 73. `s09_constructor_pattern` — Constructor function pattern (NewXxx)

## Section 10: Pointers

- [x] 74. `s10_pointer_basics` — & and * operators, pointer declaration
- [x] 75. `s10_new_keyword` — new() function
- [x] 76. `s10_nil_pointers` — nil pointers, zero value of pointers
- [x] 77. `s10_no_pointer_arithmetic` — No pointer arithmetic in Go
- [x] 78. `s10_value_vs_pointer_receiver` — When to use value vs pointer receiver
- [x] 79. `s10_pointer_to_struct` — Automatic dereferencing with struct pointers

## Section 11: Interfaces

- [x] 80. `s11_defining_interfaces` — Defining interfaces
- [x] 81. `s11_implicit_implementation` — Implicit interface satisfaction
- [x] 82. `s11_empty_interface` — Empty interface (any / interface{})
- [x] 83. `s11_type_assertions` — Type assertions
- [x] 84. `s11_type_switch` — Type switch
- [x] 85. `s11_interface_composition` — Interface embedding/composition
- [x] 86. `s11_stringer_interface` — fmt.Stringer interface
- [x] 87. `s11_error_interface` — error interface
- [x] 88. `s11_io_reader_writer` — io.Reader and io.Writer
- [x] 89. `s11_nil_interface_gotcha` — nil interface vs interface holding nil

## Section 12: Error Handling

- [x] 90. `s12_error_basics` — error interface, errors.New
- [x] 91. `s12_fmt_errorf` — fmt.Errorf and error formatting
- [x] 92. `s12_custom_error_types` — Custom error types
- [x] 93. `s12_wrapping_errors` — Error wrapping with %w
- [x] 94. `s12_errors_is_as` — errors.Is and errors.As
- [x] 95. `s12_panic_recover` — panic, recover, and defer
- [x] 96. `s12_when_to_panic` — When to panic vs return error

## Section 13: Goroutines

- [x] 97. `s13_goroutine_basics` — Creating goroutines with go keyword
- [x] 98. `s13_waitgroup` — sync.WaitGroup for synchronization
- [x] 99. `s13_race_conditions` — Race conditions and -race flag
- [x] 100. `s13_mutex` — sync.Mutex and sync.RWMutex
- [x] 101. `s13_sync_once` — sync.Once for one-time initialization
- [x] 102. `s13_atomic` — sync/atomic operations

## Section 14: Channels

- [x] 103. `s14_unbuffered_channels` — Unbuffered channels (synchronous)
- [x] 104. `s14_buffered_channels` — Buffered channels
- [x] 105. `s14_channel_direction` — Channel direction (send-only, receive-only)
- [x] 106. `s14_select` — select statement
- [x] 107. `s14_timeouts` — Timeouts with select and time.After
- [x] 108. `s14_closing_channels` — Closing channels, range over channel
- [x] 109. `s14_fan_in_fan_out` — Fan-in and fan-out patterns
- [x] 110. `s14_done_channel` — Done channel pattern for cancellation

## Section 15: Packages and Modules

- [x] 111. `s15_go_mod_init` — go mod init, module path
- [x] 112. `s15_importing` — Importing packages, aliases, dot import
- [x] 113. `s15_exported_names` — Exported vs unexported (capitalization)
- [x] 114. `s15_init_function` — init() function ordering and side effects
- [x] 115. `s15_internal_packages` — internal/ package restriction
- [x] 116. `s15_go_sum` — go.sum file and module verification

## Section 16: Generics (Go 1.18+)

- [x] 117. `s16_type_parameters` — Type parameters on functions
- [x] 118. `s16_constraints` — Type constraints and interfaces as constraints
- [x] 119. `s16_any_comparable` — any and comparable built-in constraints
- [x] 120. `s16_generic_types` — Generic structs and types
- [x] 121. `s16_type_sets` — Type sets and union elements (~int | ~float64)

## Section 17: Testing

- [x] 122. `s17_basic_tests` — Test functions and go test
- [x] 123. `s17_table_driven_tests` — Table-driven tests pattern
- [x] 124. `s17_subtests` — Subtests with t.Run
- [x] 125. `s17_test_helpers` — Test helpers and t.Helper()
- [x] 126. `s17_benchmarks` — Benchmarks with testing.B
- [x] 127. `s17_example_tests` — Example tests (testable documentation)
- [x] 128. `s17_testmain` — TestMain for setup/teardown
- [x] 129. `s17_tdd` — TDD (Test-Driven Development) — RED→GREEN→REFACTOR cycle

## Section 18: Standard Library Highlights

- [x] 129. `s18_fmt_deep_dive` — fmt package: verbs, width, precision, flags
- [x] 130. `s18_os_package` — os package: args, env, file operations
- [x] 131. `s18_io_package` — io package: Reader, Writer, Copy, TeeReader
- [x] 132. `s18_sort_package` — sort package: Sort, Slice, Search
- [x] 133. `s18_time_package` — time package: Now, Duration, Ticker, formatting
- [x] 134. `s18_json_package` — encoding/json: Marshal, Unmarshal, tags
- [x] 135. `s18_http_basics` — net/http: simple server and client
- [x] 136. `s18_context_package` — context: WithCancel, WithTimeout, WithValue
- [x] 137. `s18_log_package` — log and slog packages

## Section 19: Advanced Patterns

- [x] 138. `s19_functional_options` — Functional options pattern
- [x] 139. `s19_builder_pattern` — Builder pattern
- [x] 140. `s19_singleton` — Singleton with sync.Once
- [x] 141. `s19_worker_pool` — Worker pool pattern
- [x] 142. `s19_pipeline` — Pipeline pattern with channels
- [x] 143. `s19_method_sets` — Method sets and interface satisfaction
- [x] 144. `s19_embedding_vs_inheritance` — Embedding is not inheritance

## Section 20: Reflection and Unsafe

- [x] 145. `s20_reflect_basics` — reflect.TypeOf, reflect.ValueOf
- [x] 146. `s20_reflect_struct_tags` — Reading struct tags via reflection
- [x] 147. `s20_unsafe_overview` — unsafe.Pointer, unsafe.Sizeof, uintptr
