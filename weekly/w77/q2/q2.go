/*
 * algo
 * Package: q2
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/4/30
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumAverageDifference([]int{2, 5, 3, 9, 5, 3}))
	fmt.Println(minimumAverageDifference([]int{0}))
}

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	arr := [][]int{}
	rsum := 0
	for j := 1; j < n; j++ {
		rsum += nums[j]
	}
	arr0 := []int{nums[0], rsum}
	arr = append(arr, arr0)
	for i := 1; i < n; i++ {
		item := []int{}
		lsum := arr[i-1][0] + nums[i]
		rsum := arr[i-1][1] - nums[i]
		item = append(item, lsum, rsum)
		arr = append(arr, item)
	}
	results := make([]int, n)
	for i := range nums {
		rn := n - (i+1)
		cnt := 0
		if arr[i][1] == 0 {
			cnt = arr[i][0] / (i+1)
		} else {
			cnt = arr[i][0] / (i+1) - arr[i][1] / rn
		}
		results[i] = abs(cnt)
	}
	min := math.MaxInt
	index := -1
	for i, item := range results {
		if item < min {
			min = item
			index = i
		}
	}
	return index
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
