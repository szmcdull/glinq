# Copilot Instructions for glinq

This repository implements a LINQ-style library for Go, leveraging Go 1.18+ generics and Go 1.23+ iterators. It is organized into modular packages for different data operations, with a focus on composable, lazy-evaluated pipelines.

## Architecture & Key Components
- **giter/**: Core LINQ-like iterator operations using `iter.Seq[T]` (Go 1.23+). Includes filtering, mapping, grouping, chunking, set ops, etc.
- **garray/**: Slice-based utilities for direct operations on Go slices, including sorting, mapping, grouping, and more. Use for performance-critical or non-iterator scenarios.
- **gset/**: Set operations, supporting both in-place and chainable (fluent) styles. See `gset/README.md` for usage patterns.
- **gmap/**: Generic map utilities, including sync map wrappers.
- **partial/**: Partial function application (argument binders) for Go functions.
- **unsafe/**: Unsafe, non-error-checking versions of core operations for performance or simplicity.
- **adapters.go**: Adapters to convert between slices, maps, channels, and iterators.

## Patterns & Conventions
- All iterators are **lazily evaluated**. Most functions return `iter.Seq[T]` or similar, which are only executed when ranged.
- Operations like `GroupBy` and `OrderBy` require full traversal and will buffer all input; others (e.g., `Chunk`, `Where`, `Select`) are streaming.
- Prefer using `garray` for slice-heavy workflows; use `giter` for streaming/iterator pipelines.
- Set operations in `gset` can be in-place (mutate first arg) or chainable (return new set). See `gset/README.md` for idioms.
- Partial application in `partial/` uses a naming convention: `Bind[x_y][r[z]]` (see `partial/README.md`).

## Developer Workflows
- **Build**: Standard Go build (`go build ./...`).
- **Test**: Standard Go test (`go test ./...`). Each package has its own tests (e.g., `giter/giter_test.go`).
- **Debug**: Use Go's built-in tooling. No custom scripts required.
- **Docs/Examples**: See each package's `README.md` for usage and idioms. Main `README.md` gives high-level overview.

## Integration & External Dependencies
- Relies on Go 1.23+ for iterator support (`iter.Seq[T]`).
- No external dependencies except standard Go modules.
- Some adapters for `sqlx` and `bufio.Scanner` are provided.

## Examples
- See `giter/README.md` for iterator chaining, grouping, chunking, etc.
- See `garray/README.md` for slice utilities.
- See `gset/README.md` for set idioms (in-place vs. chain).
- See `partial/README.md` for function binding patterns.

## Notes
- Do not assume all operations are streaming; check if buffering is required (see docs).
- When in doubt, prefer explicit conversion (e.g., `ToSlice`, `FromSlice`).
- For performance-sensitive code, benchmark both `giter` and `garray` approaches.
