package main

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
func findInMountainArray(target int, mountainArr *MountainArray) int {
	// find Mountain Top
	low, high := 1, mountainArr.length()-2
	top := 0
	for low <= high {
		mid := low + (high-low)/2
		midVal := mountainArr.get(mid)
		left := mountainArr.get(mid - 1)
		right := mountainArr.get(mid + 1)
		if midVal > left && midVal > right {
			top = mid
			break
		} else if midVal > left {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if mountainArr.get(top) >= target {
		low, high = top, mountainArr.length()-1
	} else {
		low, high = 0, top
	}
	low, high = 0, top
	for low < high {
		mid := low + (high-low)/2
		midVal := mountainArr.get(mid)
		if midVal < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if mountainArr.get(low) == target {
		return low
	}
	low, high = top, mountainArr.length()-1
	for low < high {
		mid := low + (high-low)/2
		midVal := mountainArr.get(mid)
		if midVal <= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if mountainArr.get(low) == target {
		return low
	}
	return -1
}
