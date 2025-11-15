# Advent of Code 2019 Solutions

## Day 19: Tractor Beam

### Part 1
✅ **Status**: Confirmed correct
**Answer**: 160

### Part 2
❌ **Status**: INCORRECT - need to fix

**Attempted answers** (all too high):
1. 21062710 - Too high (was checking bottom-right corner incorrectly)
2. 15362148 - Too high (checking top-right and bottom-left from top-left position)
3. 11491534 - Too high (scanning bottom row, checking top-right corner)

**Issue**: Square detection algorithm is still wrong. The example (10×10 square at position 25,20 = 250020) passes, but the algorithm gives wrong results for the actual input.

The problem asks for the top-left corner (closest to emitter) of a 100×100 square that fits entirely within the tractor beam, encoded as `x*10000 + y`.

**Algorithm currently used:**
- Scan rows as BOTTOM edge of the square (y = square, square+1, ...)
- Find leftmost beam point at that row (bottom-left corner at x,y)
- Check if top-right corner (x+99, y-99) is in beam
- Return top-left corner (x, y-99)

Need to debug why this works for the example but fails for actual input.
