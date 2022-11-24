package main

import (
	"fmt"
)

// https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/
// 795. 区间子数组个数
// medium
func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3)) // 3
	// fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8)) // 7
	// [2] [5] [6] [2 5] [5 6] [2 5 6] [2]
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	ans := 0
	last1, last2 := -1, -1
	for i, num := range nums {
		if num >= left && num <= right {
			last1 = i
		} else if num > right {
			last2 = i
			last1 = -1
		}
		if last1 != -1 {
			ans += last1 - last2
		}
	}
	return ans
}

func numSubarrayBoundedMax1(nums []int, left int, right int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			tmp := nums[i:j]
			valid, idx := checkValid(tmp, left, right)
			if valid {
				ans++
			} else {
				i = idx + 1
			}
		}
	}
	return ans
}

func checkValid(arr []int, l, r int) (bool, int) {
	max := -1
	maxIndex := -1
	for i, item := range arr {
		if item > max {
			max = item
			maxIndex = i
		}
	}
	if max < l || max > r {
		return false, maxIndex
	}
	return true, maxIndex
}
