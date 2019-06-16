package main

func swap_1090(nums []int, a, b int) {
	tmp := nums[a]
	nums[a] = nums[b]
	nums[b] = tmp
}

func partition_1090_2(values, labels []int, low, high int) int {
	i, j, pivot, pivot_label := low, high, values[low], labels[low]
	for i < j {
		for i < j && values[j] <= pivot {
			j--
		}
		values[i] = values[j]
		labels[i] = labels[j]
		for i < j && values[i] >= pivot {
			i++
		}
		values[j] = values[i]
		labels[j] = labels[i]
	}
	values[i] = pivot
	labels[i] = pivot_label
	return i
}

func partition_1090(values, labels []int, low, high int) int {
	pivot := values[(low+high)/2]
	i, j := low, high
	for i < j {
		for values[i] > pivot {
			i++
		}
		for values[j] < pivot {
			j--
		}
		if i < j {

			swap_1090(values, i, j)
			swap_1090(labels, i, j)
			println(values[i], values[j])
			println(values[i], values[j])
			i++
			j--
		}
	}
	return i
}

func quickSort_1090(values, labels []int, low, high int) {
	if low < high {
		index := partition_1090_2(values, labels, low, high)
		quickSort_1090(values, labels, low, index-1)
		quickSort_1090(values, labels, index+1, high)
	}
}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	quickSort_1090(values, labels, 0, len(values)-1)
	used := make(map[int]int)
	num_used := 0
	res := 0
	for i := 0; i < len(values) && num_used < num_wanted; i++ {
		if used[labels[i]] < use_limit {
			res += values[i]
			used[labels[i]]++
			num_used++
			//println("Use:", values[i])
		}
	}
	return res
}
