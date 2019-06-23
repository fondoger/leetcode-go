package main

import "sort"

type Item [][]int

func (a Item) Len() int           { return len(a) }
func (a Item) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Item) Less(i, j int) bool { return a[i][0] < a[j][0] }

func carPooling(trips [][]int, capacity int) bool {
	items := make([][]int, 0)
	for _, trip := range trips {
		items = append(items, []int{trip[1], trip[0]})
		items = append(items, []int{trip[2], -trip[0]})
	}
	sort.Sort(Item(items))

	sum := 0
	for i := 0; i < len(items); i++ {
		sum += items[i][1]
		for i+1 < len(items) && items[i+1][0] == items[i][0] {
			sum += items[i+1][1]
			i++
		}
		if sum > capacity {
			return false
		}
	}
	return true
}
