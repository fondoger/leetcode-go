package main

func main() {
	{
		list1 := []int{1, 3, 5, 7, 9}
		list2 := []int{3, 4, 5, 6, 7}
		list3 := []int{2, 4, 6, 7, 8, 10}
		l1 := buildLinkedList(list1)
		l2 := buildLinkedList(list2)
		l3 := buildLinkedList(list3)
		printLinkedList(l1)
		printLinkedList(l2)
		printLinkedList(l3)
		listHead := mergeKLists([]*ListNode{l1, l2, l3})
		printLinkedList(listHead)
	}
	{
		nums := []int{13, 123, 12, 342, 12, 123, 432, 23, 55, 123, 3234}
		printIntArray(nums)
		qsort(nums, 0, len(nums)-1)
		printIntArray(nums)
	}
}

func qsort(nums []int, low, high int) {
	if low >= high {
		return
	}
	i, j, pivot := low, high, nums[low]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	qsort(nums, low, i-1)
	qsort(nums, i+1, high)
}
