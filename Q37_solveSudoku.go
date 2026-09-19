package leetcode

func solveSudoku(board [][]byte) {
	slots := make([][]int, 0, 81)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				slots = append(slots, []int{i, j})
			}
		}
	}
	_ = canPutSlot(board, slots, 0)
}

func canPutSlot(board [][]byte, slots [][]int, i int) bool {
	if i >= len(slots) {
		return true
	}
	x := slots[i][0]
	y := slots[i][1]
	for _, c := range []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'} {
		if canPut(board, c, x, y) {
			board[x][y] = c
			if canPutSlot(board, slots, i+1) {
				return true
			}
			board[x][y] = '.'
		}
	}
	return false
}

func canPut(board [][]byte, c byte, x, y int) bool {
	if board[x][y] != '.' {
		return false
	}

	I := x / 3
	J := y / 3
	for i := I * 3; i < I*3+3; i++ {
		for j := J * 3; j < J*3+3; j++ {
			if board[i][j] == c {
				return false
			}
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][y] == c {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if board[x][j] == c {
			return false
		}
	}
	return true
}
