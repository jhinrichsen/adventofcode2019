# Advent of Code 2019 Solutions

## Day 19: Tractor Beam

### Part 1
✅ **Status**: Confirmed correct
**Answer**: 160

### Part 2
✅ **Status**: Confirmed correct
**Answer**: 9441282 (x=944, y=1282)

**Algorithm:**
- Scan rows as BOTTOM edge of the square (y = square, square+1, ...)
- Find leftmost beam point at that row (bottom-left corner at x,y)
- Check if top-right corner (x+99, y-99) is in beam
- Return top-left corner (x, y-99)

**Previous failed attempts** (all too high):
1. 21062710 - Was checking bottom-right corner
2. 15362148 - Checking top-right and bottom-left from top-left
3. 11491534 - Had optimization that skipped beam points incorrectly
