/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	popLeft, popRight := 0, 0
	i, j := 0, len(nums1)-1
	m, n := 0, len(nums2)-1
	for i <= j || m <= n {
		// pop from left
		if m > n || i <= j && nums1[i] < nums2[m] {
			popLeft = nums1[i]
			i++
		} else {
			popLeft = nums2[m]
			m++
		}
		if i > j && m > n {
			return float64(popLeft)
		}
		// pop from right
		if m > n || i <= j && nums1[j] > nums2[n] {
			popRight = nums1[j]
			j--
		} else {
			popRight = nums2[n]
			n--
		}
	}
	return (float64(popLeft) + float64(popRight)) / 2
}

/* Solution 2: 归并排序中的合并步骤 O(n)*/
func findMedianSortedArrays_SOLUTION2(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	sorted := make([]int, n)
	count, i, j := 0, 0, 0
	for i < len(nums1) || j < len(nums2) {
		if j >= len(nums2) || i < len(nums1) && nums1[i] < nums2[j] {
			sorted[count] = nums1[i]
			count++
			i++
		} else {
			sorted[count] = nums2[j]
			count++
			j++
		}
	}
	return float64(sorted[(n-1)/2]+sorted[n/2]) / 2
}
