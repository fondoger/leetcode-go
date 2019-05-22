/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package main

func lengthOfLongestSubstring(s string) int {
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
