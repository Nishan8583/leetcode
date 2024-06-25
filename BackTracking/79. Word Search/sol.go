/*
79. Word Search
Solved
Medium
Topics
Companies

Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.


https://leetcode.com/problems/word-search/description/

Solution:
Same as other grid problem, check up, down, left, right
Here added one trick from neetcode where we replace already included value with '*', so that we don't visit that value again
We check if it is *, if it is we return false
This way we do not have to create seperate map for each dfs
*/

package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			//fmt.Println("Checking", row, col)
			if helper(row, col, board, 0, word) {
				return true
			}
		}
	}
	return false
}

func helper(row, col int, board [][]byte, wordPos int, word string) bool {
	//fmt.Printf("Checking %d,%d and wordPos %d string %s\n", row, col, wordPos, string(word[wordPos]))
	// return false if out of bound
	if row >= len(board) || row < 0 {
		return false
	}

	// out of bound again
	if col >= len(board[0]) || col < 0 {
		return false
	}

	// if the word does not match or it is * i.e. already included, return false
	if board[row][col] != word[wordPos] || board[row][col] == '*' {
		return false
	}

	// IF we reached here, means the value is matched and it is not *
	// And if below condition matched means it is the end of the word
	// SO all letters matched, return true
	if wordPos == len(word)-1 {
		return true
	}

	tmp := board[row][col]
	board[row][col] = '*' // include the value so that we don't visit it again
	// If it does match, check if we've reached the end of the word

	if wordPos == len(word)-1 && board[row][col] == word[wordPos] {
		//fmt.Println("Final match", word[wordPos], string(word[wordPos]))
		return true
	}

	//fmt.Println("Matched", word[wordPos], string(word[wordPos]), row, col)
	// if we are not in the end, then lets check for the next word in adjacent cells
	nextRow := row + 1
	nextCol := col + 1
	prevRow := row - 1
	prevCol := col - 1

	// check left
	wasBehind := helper(row, prevCol, board, wordPos+1, word)

	// check right
	wasAhead := helper(row, nextCol, board, wordPos+1, word)

	// check up
	wasUp := helper(prevRow, col, board, wordPos+1, word)

	// check down
	wasBelow := helper(nextRow, col, board, wordPos+1, word)

	board[row][col] = tmp
	//fmt.Printf("Returning %t %t %t %t\n", wasAhead, wasBehind, wasUp, wasBelow)
	return wasAhead || wasBehind || wasUp || wasBelow

}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))

	fmt.Println("checking for SEE")
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))

	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))

	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'}}, "ABCESEEEFS"))

}
