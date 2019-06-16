package main

func duplicateZeros(arr []int) {
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
