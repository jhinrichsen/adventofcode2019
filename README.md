# Advent of Code 2019 Solutions

## Day 19: Tractor Beam

### Part 1
✅ **Status**: Confirmed correct
**Answer**: 160

### Part 2
❌ **Status**: INCORRECT - need to fix

**Attempted answers** (all too high):
1. 21062710 - Too high (was checking bottom-right corner incorrectly)
2. 15362148 - Too high (checking top-right and bottom-left)

**Issue**: Square detection algorithm is still wrong. Need to rethink how to verify a 100×100 square fits in the beam.

The problem asks for the top-left corner (closest to emitter) of a 100×100 square that fits entirely within the tractor beam, encoded as `x*10000 + y`.
