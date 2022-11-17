/*
 * algo
 * Package: q658
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/25
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findClosestElements1([]int{1, 2, 4, 5, 6}, 4, 3))
	fmt.Println(findClosestElements1([]int{1, 2, 3, 4, 5}, 4, -1))
}

func findClosestElements(arr []int, k int, x int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		return abs(arr[i]-x) < abs(arr[j]-x)
	})
	sort.Ints(arr)
	return arr[:k]
}

func findClosestElements1(arr []int, k int, x int) []int {
	right := sort.SearchInts(arr, x)
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= len(arr) || x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}


func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
