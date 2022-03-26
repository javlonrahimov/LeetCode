package main

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	N := len(grid)
	if N == 1 && grid[0][0] == 0 {
		return 1
	}
	directions := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	type Node struct {
		col      int
		row      int
		distance int
	}
	var queue []Node
	queue = append(queue, Node{0, 0, 1})
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]
		grid[n.row][n.col] = 1
		for _, dir := range directions {
			row := n.row + dir[0]
			col := n.col + dir[1]
			if row < 0 || row >= N {
				continue
			}
			if col < 0 || col >= N {
				continue
			}
			if grid[row][col] != 0 {
				continue
			}
			if row == N-1 && col == N-1 {
				return n.distance + 1
			}
			queue = append(queue, Node{col, row, n.distance + 1})
			grid[row][col] = 1
		}
	}
	return -1
}
