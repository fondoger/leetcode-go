package main

/*
 * @lc app=leetcode id=621 lang=golang
 *
 * [621] Task Scheduler
 */
func leastInterval(tasks []byte, n int) int {
	hashMap := make([]int, 26)

	var maxVal, maxCount int
	for _, t := range tasks {
		hashMap[t-'A']++
		if hashMap[t-'A'] > maxVal {
			maxVal = hashMap[t-'A']
			maxCount = 1
		} else if hashMap[t-'A'] == maxVal {
			maxCount++
		}
	}

	return max(len(tasks), (n+1)*(maxVal-1)+maxCount)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
