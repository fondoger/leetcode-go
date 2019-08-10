package main

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// 方法：回溯
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, pos, i, j int) (found bool) {
	if pos == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] == 0 || board[i][j] != word[pos] {
		return false
	}
	board[i][j] = 0
	ret := dfs(board, word, pos+1, i-1, j) ||
		dfs(board, word, pos+1, i, j-1) ||
		dfs(board, word, pos+1, i+1, j) ||
		dfs(board, word, pos+1, i, j+1)
	board[i][j] = word[pos]
	return ret
}
