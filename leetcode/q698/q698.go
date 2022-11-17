/*
 * algo
 * Package: q698
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/20
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % k != 0 {
		return false
	}
	sort.Ints(nums)
	cur := make([]int, k)
	n := len(nums)
	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == n {
			return true
		}
		for j := 0; j < k; j++ {
			if j > 0 && cur[j] == cur[j-1] {
				continue
			}
			cur[j] += nums[i]
			if cur[j] <= sum && dfs(i+1) {
				return true
			}
			cur[j] -= nums[i]
		}
		return false
	}
	return dfs(0)
}