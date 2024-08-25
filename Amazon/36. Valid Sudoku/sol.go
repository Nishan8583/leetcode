func isValidSudoku(board [][]byte) bool {

	rows := map[int]map[byte]struct{}{} // map of row number to values
	col := map[int]map[byte]struct{}{}  // map of column number to values
	boxes := map[[2]int]map[byte]struct{}{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if curr_row, ok := rows[i]; ok {
				if _, ok := curr_row[board[i][j]]; ok {
					return false
				} else {
					rows[i][board[i][j]] = struct{}{}
				}
			} else {
				rows[i] = map[byte]struct{}{
					board[i][j]: {},
				}
			}

			if curr_col, ok := col[j]; ok {
				if _, ok := curr_col[board[i][j]]; ok {
					return false
				} else {
					col[j][board[i][j]] = struct{}{}
				}
			} else {
				col[j] = map[byte]struct{}{
					board[i][j]: {},
				}
			}

			pos := [2]int{
				int(i / 3), int(j / 3),
			}
			if curr_box, ok := boxes[pos]; ok {
				if _, ok := curr_box[board[i][j]]; ok {
					return false
				} else {
					boxes[pos][board[i][j]] = struct{}{}
				}
			} else {
				boxes[pos] = map[byte]struct{}{
					board[i][j]: {},
				}
			}
		}
	}

	return true

}
