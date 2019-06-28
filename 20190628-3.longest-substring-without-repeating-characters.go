/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package main

// 20190628：二刷
func lengthOfLongestSubstring(s string) int {
	res := 0
	mapping := [128]int{}
	for i := 0; i < len(mapping); i++ {
		mapping[i] = -1
	}
	slow, fast := 0, 0
	for fast < len(s) {
		if val := mapping[s[fast]]; val >= slow {
			if fast-slow > res {
				res = fast - slow
			}
			slow = val + 1
		}
		mapping[s[fast]] = fast
		fast++
	}
	if (fast - slow) > res {
		res = fast - slow
	}
	return res
}

// 20190522 首刷
func lengthOfLongestSubstring_20190522(s string) int {
	best := 0
	prev := -1
	mapping := [128]int{}
	for i := range mapping {
		mapping[i] = -1
	}
	for i := 0; i < len(s); i++ {
		if mapping[s[i]] > prev {
			if (i - prev - 1) > best {
				best = i - prev - 1
			}
			prev = mapping[s[i]]
		}
		mapping[s[i]] = i
	}
	if len(s)-prev-1 > best {
		best = len(s) - prev - 1
	}

	return best
}
