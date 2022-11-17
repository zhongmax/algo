/*
 * algo
 * Package: q1224
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/18
 */

// https://leetcode.cn/problems/maximum-equal-frequency/
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{2, 2, 1, 1, 5, 3, 3, 5}
	fmt.Println(maxEqualFreq(nums1)) // 7
	nums2 := []int{1,1,1,2,2,2,3,3,3,4,4,4,5}
	fmt.Println(nums2) // 13
}

func maxEqualFreq(nums []int) (ans int) {
	freq := map[int]int{}
	count := map[int]int{}
	maxFreq := 0
	for i, num := range nums {
		if count[num] > 0 {
			freq[count[num]]--
		}
		count[num]++
		maxFreq = max(maxFreq, count[num])
		freq[count[num]]++
		if maxFreq == 1 ||
			freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 ||
			freq[maxFreq]*maxFreq+1 == i+1 && freq[1] == 1 {
			ans = max(ans, i+1)
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}