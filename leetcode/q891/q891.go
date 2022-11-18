package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/sum-of-subsequence-widths/
// 891. 子序列宽度之和
// hard
func main() {
	fmt.Println(sumSubseqWidths([]int{2, 1, 3})) // 6
	// fmt.Println(sumSubseqWidths([]int{2}))       // 0
}

func sumSubseqWidths(nums []int) int {
	const mod int = 1e9 + 7
	ans := 0
	// 将数组按从小到大排序, 这里因为题目要求子序列可以不连续, 我们直接采用求每个元素为最大/最小的子序列的个数
	// 我们将排序后的数组中, 每个数定为x, 下标i从0开始
	// 其中x作为最大值的子序列的个数为 2^i
	// x作为最小值的子序列的个数为 2^(n-1-i)
	// 每个x对答案的影响为 (2^i - 2^(n-1-i)) * x
	sort.Ints(nums)
	n := len(nums)
	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = pow2[i-1] * 2 % mod // 预处理 2 的幂次
	}
	for i, x := range nums {
		ans += (pow2[i] - pow2[n-1-i]) * x // 在题目的数据范围下，这不会溢出
	}
	return (ans%mod + mod) % mod // 注意上面有减法，ans 可能为负数
}
