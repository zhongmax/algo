/*
 * algo
 * Package: q462
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/19
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMoves2([]int{1, 2, 3}))
	fmt.Println(minMoves2([]int{1, 10, 2, 9}))
}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	mid := nums[n/2]
	ans := 0
	for i := 0; i < n; i++ {
		ans += abs(nums[i]-mid)
	}
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}