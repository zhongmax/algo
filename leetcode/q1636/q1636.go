/*
 * algo
 * Package: q1636
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/19
 */

package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/sort-array-by-increasing-frequency/
// 1636. 按照频率将数组升序排序
// easy
func main() {
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(frequencySort([]int{2, 3, 1, 3, 2}))
	fmt.Println(frequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}

func frequencySort(nums []int) []int {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if mp[nums[i]] == mp[nums[j]] {
			return nums[i] > nums[j]
		}
		return mp[nums[i]] < mp[nums[j]]
	})
	return nums
}