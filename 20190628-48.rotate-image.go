package main

/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */
func rotate(matrix [][]int) {
	m, n := len(matrix)-1, len(matrix[0])-1
	for i := 0; i < (m+1)/2; i++ {
		for j := i; j < n-i; j++ {
			swap48(&matrix[i][j], &matrix[j][n-i])
			swap48(&matrix[i][j], &matrix[m-j][i])
			swap48(&matrix[m-j][i], &matrix[m-i][n-j])
		}
	}
}
func swap48(a *int, b *int) {
	*a, *b = *b, *a
}
