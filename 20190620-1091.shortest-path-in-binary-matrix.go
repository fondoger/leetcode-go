package main

/**
 *	LeetCode 第141次周赛 第三题
 *
 * 对于单源最短路径问题：
 * 容易理解的算法有：BFS算法、Dijkstra算法
 * 其他算法有：自己百度去，反正我也学不会
 *
 * 一般在做算法题是，BFS算法是比较容易实现的
 *
 * 需要注意的是：当点与点之间的权值全部相等时，直接使用队列进行BFS即可，
 * 特别简单，类似二叉树的层序遍历。此时普通BFS与Dijkstra效率一样。(注意1. 结点不要重复加入 2. 队列是类似带层次层序遍历分批读取)
 * 如果权值不相等，也可以用队列，不过一个结点可能被加入队列多次，
 * 即当发现到达这个结点的长度能更优时，此时效率逊于Dijkstra
 */

var Dir = [8][2]int{
	{1, -1}, {1, 0}, {1, 1},
	{0, 1}, {0, -1},
	{-1, 1}, {-1, 0}, {-1, 1},
}

// 方法二：可重复的队列BFS解法，适用于所有有权的单源最短路径问题
// 这个方法比方法一通用一点，写起来考虑的东西也少一点
// 主要就是利用队列，并且一个点可以重复进入队列，即当发现有到达该点的更优路径时
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	visit := make([][]int, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]int, n)
	}
	queue := make(Queue47, 0)
	queue.offer(Point47{0, 0})
	visit[0][0] = 1
	for !queue.isEmpty() {
		cur := queue.poll()
		for i := 0; i < 8; i++ {
			next := Point47{cur.x + Dir[i][0], cur.y + Dir[i][1]}
			if next.x >= 0 && next.x < m && next.y >= 0 && next.y < n && grid[next.x][next.y] == 0 {
				if visit[next.x][next.y] == 0 || visit[cur.x][cur.y]+1 < visit[next.x][next.y] {
					// 要么是之前没走过，要么是发现了到达该点的更短路径
					visit[next.x][next.y] = visit[cur.x][cur.y] + 1
					queue.offer(next)
				}
			}
		}
	}
	if visit[m-1][n-1] == 0 {
		return -1
	}
	return visit[m-1][n-1]
}

// 方法一：不重复加入的队列BFS解法：适用于权值相等的最短路径问题，用类似带层次的层序遍历分批处理
func shortestPathBinaryMatrix_1(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	queue := make(Queue47, 0, m+n)
	queue.offer(Point47{0, 0})
	res := 0
	for !queue.isEmpty() {
		res++ // res依批次递增
		size := len(queue)
		for i := 0; i < size; i++ { // 注意点1：分批读取
			cur := queue.poll()
			if cur.x == m-1 && cur.y == n-1 {
				return res
			}
			for i := 0; i < 8; i++ {
				next := Point47{cur.x + Dir[i][0], cur.y + Dir[i][1]}
				if next.x >= 0 && next.x < m && next.y >= 0 && next.y < n && grid[next.x][next.y] == 0 {
					println("offer", next.x, next.y)
					queue.offer(next)
					grid[next.x][next.y] = 1 // 注意点2：
				}
			}
		}
	}
	return -1
}

/* A simple Queue implementation */
type Point47 struct {
	x, y int
}
type Queue47 []Point47

func (q *Queue47) offer(val Point47) {
	*q = append(*q, val)
}
func (q *Queue47) poll() Point47 {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}
func (q *Queue47) isEmpty() bool {
	return len(*q) == 0
}
