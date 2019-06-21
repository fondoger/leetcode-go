package main

/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

/* 解法2：滑动窗口：进入窗口计数递减，出窗口计数递增 + 判断是否全为零 */
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	counter := [26]int{}
	for i := 0; i < len(s1); i++ {
		counter[s1[i]-'a']++
		counter[s2[i]-'a']--
	}
	if allZero(counter) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		counter[s2[i]-'a']--         // 进入窗口递减
		counter[s2[i-len(s1)]-'a']++ // 出窗口递增
		if allZero(counter) {
			return true
		}
	}

	return false
}

func allZero(nums [26]int) bool {
	for _, val := range nums {
		if val != 0 {
			return false
		}
	}
	return true
}

/* 解法1：我的笨办法：不停的判断字串是否可由s1中的字符组成 */
func checkInclusion_1(s1 string, s2 string) bool {
	alphabet := countOccurance(s1)
	for i, j := 0, len(s1); j <= len(s2); i, j = i+1, j+1 {
		if judgeCombination(alphabet, s2[i:j]) {
			return true
		}
	}
	return false
}

// countOccurance: 返字符串各个字符的出现次数, 及字符种类数目
func countOccurance(s1 string) []int {
	alphabet := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		alphabet[s1[i]-'a']++
	}
	return alphabet
}

func judgeCombination(alphabet []int, str string) bool {
	copied := make([]int, len(alphabet))
	copy(copied, alphabet)
	for i := 0; i < len(str); i++ {
		if copied[str[i]-'a'] > 0 {
			copied[str[i]-'a']--
		} else {
			return false
		}
	}
	return true
}
