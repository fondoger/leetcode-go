/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package main

/* Solution 1: 从左右两边不停的pop元素 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, len(nums1)-1
	m, n := 0, len(nums2)-1
	popLeft := 0
	popRight := 0
	for i <= j || m <= n {
		// pop 1 item from left
		if m > n || i <= j && nums1[i] < nums2[m] {
			popLeft = nums1[i]
			i++
		} else {
			popLeft = nums2[m]
			m++
		}
		// pop 1 item from right
		if !(i <= j || m <= n) {
			return float64(popLeft)
		}
		if m > n || i <= j && nums1[j] > nums2[n] {
			popRight = nums1[j]
			j--
		} else {
			popRight = nums2[n]
			n--
		}
	}
	return float64(popLeft+popRight) / 2
}
