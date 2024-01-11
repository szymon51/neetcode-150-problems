package validsudoku

func makeTwoDimentionalArray() [][]bool {
	var twoD [][]bool
	for i := 0; i < 9; i++ {
		oneD := make([]bool, 9)
		twoD = append(twoD, oneD)
	}

	return twoD
}

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
//	Each row must contain the digits 1-9 without repetition.
//	Each column must contain the digits 1-9 without repetition.
//	Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//
// Note:
//
//	A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//	Only the filled cells need to be validated according to the mentioned rules.
func isValidSudoku(board [][]byte) bool {
	col := makeTwoDimentionalArray()
	row := makeTwoDimentionalArray()
	square := makeTwoDimentionalArray()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			n := board[i][j] - '1'
			if col[j][n] {
				return false
			}

			if row[i][n] {
				return false
			}

			b := i/3*3 + j/3
			if square[b][n] {
				return false
			}

			col[j][n] = true
			row[i][n] = true
			square[b][n] = true
		}
	}

	return true
}
