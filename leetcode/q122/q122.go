/*
 * algo
 * Package: q122
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/4/26
 */

package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] -= prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	hold := -prices[0]
	noHold := 0
	for i := 1; i < len(prices); i++ {
		hold = max(hold, noHold-prices[i])
		noHold = max(noHold, hold+prices[i])
	}
	return noHold
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}