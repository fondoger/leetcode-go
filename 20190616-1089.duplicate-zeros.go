package main

/*
大佬们都是用两个数组完成的
*/
func duplicateZeros(nums []int) {
	copied := make([]int, len(nums))
	copy(copied, nums)
	m, n := 0, 0
	for n < len(nums) {
		nums[n] = copied[m]
		n++
		m++
		if copied[m] == 0 && n < len(copied) {
			nums[n] = 0
			n++
		}
	}
}

/* 下面是我用队列实现的蠢办法 */

func duplicateZeros_1(arr []int) {
	queue := make(Queue, 0, 10)
	for i := 0; i < len(arr); i++ {
		if len(queue) != 0 {
			queue.offer(arr[i])
			arr[i] = queue.poll()
		}
		if arr[i] == 0 {
			if i+1 < len(arr) {
				queue.offer(arr[i+1])
				arr[i+1] = 0
				i += 1
			}
		}
	}
}

type Queue []int

func (queue *Queue) offer(val int) {
	*queue = append(*queue, val)
}
func (queue *Queue) poll() int {
	res := (*queue)[0]
	*queue = (*queue)[1:]
	return res
}
