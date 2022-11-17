/*
 * algo
 * Package: q398
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/4/25
 */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := Constructor([]int{1,2,3,3,3})
	fmt.Println(s.Pick(3))
	fmt.Println(s.Pick(2))
	fmt.Println(s.Pick(1))
}

type Solution []int


func Constructor(nums []int) Solution {
	return nums
}


func (nums Solution) Pick(target int) int {
	ans, cnt := 0, 0
	for i, num := range nums {
		if num == target {
			cnt++
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return ans
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
