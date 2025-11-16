# Claude Code Guidelines for Advent of Code 2019

## Critical Rules

### Function Signatures (PRIMARY RULE)
- **MUST** implement: `func DayXX(<input>) uint`
- **SHOULD** use: `func DayXX(<input>, part1 bool) uint` unless alternatives are more elegant
- **IF** parser required (input cannot be directly processed by `input_test.go` functions):
  - Parser: `func NewDayXX(<input>) DayXXPuzzle` (return by value)
  - Combined: `func DayXX(puzzle DayXXPuzzle) uint`
- **NEVER** use methods: `func (p *DayXXPuzzle) DayXX() uint`

### File Access Prohibition
- Puzzles must not perform I/O
- **ONLY** tests may read files using `input_test.go` functions

### uint Pattern (MANDATORY)
- **ALL** puzzle return types that are counts/sums/totals/amounts must be `uint`
- Push `uint` contract up the entire call chain
- Area, perimeter, distances, prices are inherently non-negative
- Example: `func exploreRegion(...) (area, perimeter uint)`

### Coordinate System
- **ALWAYS** use `x`/`y` throughout (never row/col)
- `dimX` (width), `dimY` (height) for dimensions
- `grid[y][x]` indexing pattern
- `startY, startX int` parameter order

### Data Types
- **ALWAYS** use `byte` for ASCII characters (A-Z, 0-9, symbols)
- **NEVER** use `rune` - unnecessary UTF-8 overhead for AoC
- Use `[]byte(string)` for conversion, not manual loops

### Algorithm Requirements
- **NEVER** use recursion
- **ALWAYS** use iterative with explicit stacks: `[]image.Point`
- Use `image.Point{X: x, Y: y}` for coordinates

### Modern Go Patterns (MANDATORY)
- **ALWAYS** use latest Go 1.24+ features where applicable
- Use `for range N` instead of `for i := 0; i < N; i++` (range over integers)
- Use `slices` package: `slices.Equal`, `slices.Contains`, `slices.Sort`
- Use `maps` package: `maps.Equal`, `maps.Clone` when needed
- Use `clear(map)` and `clear(slice)` for efficient clearing
- Use `min()` and `max()` built-in functions

### Test Structure
- Table-driven tests with external files
- `testdata/dayXX_example1.txt` not inline strings
- **NEVER** use multiline string literals in tests - always use external testdata files
- `lines := linesFromFilename(t, filename)` in tests only
- Multiple examples: use `example1Filename(day)`, `example2Filename(day)`, etc.
- Available filename functions: `exampleFilename()`, `exampleNFilename()`, `example1Filename()`, `example2Filename()`, `example3Filename()`, `filename()`

### Input Parsing (Flexible)
- **Parser is optional** - only use if beneficial for complexity
- `func DayXX(input []byte)` - fine if puzzle can parse bytes directly
- `func DayXX(lines []string)` - fine if puzzle needs line-based input
- `func NewDayXX()` + `DayXX(puzzle)` - use for complex data structures
- Choose the most appropriate input format for each puzzle's needs
- Use appropriate `input_test.go` helper functions

## Commit Message Convention

Use conventional commits with day number as scope:

**Format:** `<type>(<scope>): <description>`

**Examples:**
- `feat(day14): add part 1 solution`
- `feat(day14): add part 2 with binary search`
- `test(day14): add example tests for part 1`
- `refactor(day14): convert to table-driven tests`
- `fix(day14): correct ORE calculation for surplus`
- `docs(day14): add algorithm explanation`

**Types:**
- `feat`: New feature/solution
- `fix`: Bug fix
- `refactor`: Code refactoring
- `test`: Test additions/changes
- `docs`: Documentation
- `chore`: Build/tooling changes

**Scope:**
- Use `dayXX` (lowercase, zero-padded) for day-specific commits
- Omit scope for repository-wide changes (e.g., `chore: update CLAUDE.md`)

## Benchmark Optimization Workflow

When optimizing puzzle solutions, follow this workflow to measure and document performance improvements:

### Process

1. **Create baseline benchmark (b0)**
   ```bash
   go test -run=^$ -bench=DayXXPart.$ -count=8 -benchmem > b0
   ```
   Remove the last two lines (PASS and ok lines) from b0

2. **Optimize the code**
   Apply performance improvements following the guidelines above

3. **Run benchmark again (b1)**
   ```bash
   go test -run=^$ -bench=DayXXPart.$ -count=8 -benchmem > b1
   ```
   Remove the last two lines from b1

4. **Compare with benchstat**
   ```bash
   benchstat b0 b1
   ```

5. **Document if worthwhile (>5% improvement)**
   - Add a "Day XX" section to README.adoc if not present
   - Include the benchstat output
   - Add a short, concise explanation of the optimization
   - NEVER use bold (**) in README.adoc - use proper AsciiDoc formatting

6. **Repeat** until no further significant improvements

### Example README.adoc Entry

```asciidoc
== Day 01: The Tyranny of the Rocket Equation

Optimized fuel calculation by eliminating redundant allocations.

----
name       old time/op    new time/op    delta
Day01-16     1.23µs ± 2%    0.95µs ± 1%  -22.76%

name       old alloc/op   new alloc/op   delta
Day01-16      512B ± 0%       0B ± 0%  -100.00%
----
```
