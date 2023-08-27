package main

/*
	3 maps
	one contains row number to values in that row, on next check if that row already has the value
	one contains column number to values in that column
	final has index [2]int, if u divide each row and column and get the int, u get the box
	i.e.
	[["5","3",".",".","7",".",".",".","."]
	["6",".",".","1","9","5",".",".","."]
	,[".","9","8",".",".",".",".","6","."]
	,["8",".",".",".","6",".",".",".","3"]
	,["4",".",".","8",".","3",".",".","1"]
	,["7",".",".",".","2",".",".",".","6"]
	,[".","6",".",".",".",".","2","8","."]
	,[".",".",".","4","1","9",".",".","5"]
	,[".",".",".",".","8",".",".","7","9"]]

	for
	["5","3",".",]
	["6",".",".",]
	[".","9","8",]
	if u divide the row and column for these places, and get int, and put in the map
	u get the box area, we can append the values there and check if the values have been repeated

	My solution works ok too, if this is confusing, just use grid traversal

*/
func isValidSudoku(board [][]byte) bool {

	rows := map[int]map[byte]struct{}{}     // map of row number to values, each index reperents row number, and values are all the values in that row
	col := map[int]map[byte]struct{}{}      // map of column number to values, each index represents column number, and values are all the values in that row
	boxes := map[[2]int]map[byte]struct{}{} // the box, diviing by 3 represents the same box

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
