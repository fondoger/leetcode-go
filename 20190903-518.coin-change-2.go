/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change 2
 */
func change(amount int, coins []int) int {
	memo := make([][]int, 501)
	for i := 0; i <= 500; i++ {
		memo[i] = make([]int, 5001)
		for j := 0; j <= 5000; j++ {
			memo[i][j] = -1
		}
	}
	return helper518(memo, coins, amount, 0)
}
func helper518(memo [][]int, coins []int, amount int, begin int) int {
	if amount == 0 {
		return 1
	} else if amount < 0 {
		return 0
	}
	if memo[begin][amount] != -1 {
		return memo[begin][amount]
	}
	res := 0
	for i := begin; i < len(coins); i++ {
		res += helper518(memo, coins, amount-coins[i], i)
	}
	memo[begin][amount] = res
	return res
}
