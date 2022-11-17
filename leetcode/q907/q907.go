/*
 * algo
 * Package: q907
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/16
 */

package main

import "fmt"

func main() {
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}

func sumSubarrayMins(arr []int) int {
	// 单调栈+dp
	const mod int = 1e9+7
	stack := []int{}
	dp := make([]int, len(arr)+1)
	res := 0
	stack = append(stack, -1)

	for i := 0; i < len(arr); i++ {
		for stack[len(stack)-1] != -1 && arr[i] <= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		dp[i+1] = (dp[stack[len(stack)-1]+1] + (i-stack[len(stack)-1])*arr[i]) % mod
		stack = append(stack, i)
		res += dp[i+1]
		res %= mod
	}
	return res
}
