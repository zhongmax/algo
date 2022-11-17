/*
 * algo
 * Package: q1700
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/10/19
 */

package main

import "fmt"

// https://leetcode.cn/problems/number-of-students-unable-to-eat-lunch/
// 1700. 无法吃午餐的学生数量
// easy
func main() {
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}

func countStudents(students []int, sandwiches []int) int {
	s1 := 0
	for _, s := range students {
		s1 += s
	}
	s0 := len(students) - s1
	for _, s := range sandwiches {
		if s == 0 && s0 > 0 {
			s0--
		} else if s == 1 && s1 > 0 {
			s1--
		} else {
			break
		}
	}
	return s0 + s1
}