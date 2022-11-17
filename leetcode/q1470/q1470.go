/*
 * algo
 * Package: q1470
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/29
 */

package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	fmt.Println(shuffle([]int{1, 1, 2, 2}, 2))
}

func shuffle(nums []int, n int) []int {
	ans := make([]int, len(nums))
	index := 0
	for i := 0; i < n; i++ {
		ans[index] = nums[i]
		index++
		ans[index] = nums[n+i]
	}
	return ans
}