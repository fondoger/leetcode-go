package main

import "sort"

func indexOfString(text string, sub string) [][]int {
	pairs := make([][]int, 0, 3)
	for i := 0; i < len(text)-len(sub)+1; i++ {
		match := true
		for j := 0; j < len(sub) && i+j < len(text); j++ {
			if text[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			pair := []int{i, i + len(sub) - 1}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func indexPairs(text string, words []string) [][]int {
	res := make([][]int, 0, 10)
	for _, word := range words {
		pairs := indexOfString(text, word)
		res = append(res, pairs...)
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] == res[j][0] {
			return res[i][1] < res[j][1]
		} else {
			return res[i][0] < res[j][0]
		}
	})
	return res
}
