package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	return dfs64(grid, 0, 0, make(map[string]int))
}

func dfs64(grid [][]int, row, col int, cache map[string]int) int {
	if (row < 0 || row >= len(grid)) || (col < 0 || col >= len(grid[row])) {
		return math.MaxInt32
	}
	if row == len(grid)-1 && col == len(grid[0])-1 {
		return grid[row][col]
	}
	key := fmt.Sprintf("%d:%d", row, col)
	if val, ok := cache[key]; ok {
		return val
	}

	down := grid[row][col] + dfs64(grid, row+1, col, cache)
	right := grid[row][col] + dfs64(grid, row, col+1, cache)
	cache[key] = min(down, right)
	return cache[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
