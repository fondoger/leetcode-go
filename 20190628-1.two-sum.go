package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i, val := range nums {
		if idx, ok := mapping[target-val]; ok {
			return []int{idx, i}
		} else {
			mapping[val] = i
		}
	}
	return nil
}
