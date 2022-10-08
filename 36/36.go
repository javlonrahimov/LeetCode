package main

func isValidSudoku(board [][]byte) bool {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if hasDuplicates(column(x, board)) {
				return false
			}
			if hasDuplicates(row(y, board)) {
				return false
			}
			if hasDuplicates(square(x, y, board)) {
				return false
			}
		}
	}
	return true
}

func row(y int, board [][]byte) []byte {
	return board[y]
}

func column(x int, board [][]byte) []byte {
	xs := make([]byte, 9)
	for y := 0; y < len(board); y++ {
		xs[y] = board[y][x]
	}
	return xs
}

func square(x, y int, board [][]byte) []byte {
	xs := make([]byte, 9)
	x1 := (x / 3) * 3
	y1 := (y / 3) * 3
	idx := 0
	for y2 := y1; y2 < y1+3; y2++ {
		for x2 := x1; x2 < x1+3; x2++ {
			xs[idx] = board[y2][x2]
			idx += 1
		}
	}
	return xs
}

func hasDuplicates(xs []byte) bool {
	m := map[byte]struct{}{}
	for _, x := range xs {
		if x > '0' {
			if _, ok := m[x]; ok {
				return true
			} else {
				m[x] = struct{}{}
			}
		}
	}
	return false
}
