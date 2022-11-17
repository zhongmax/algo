/*
 * algo
 * Package: q1450
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/19
 */

package main

import "fmt"

func main() {
	fmt.Println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
	fmt.Println(busyStudent([]int{4}, []int{4}, 5))
	fmt.Println(busyStudent([]int{9,8,7,6,5,4,3,2,1}, []int{10,10,10,10,10,10,10,10,10}, 5))
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ans := 0
	n := len(startTime)
	for i := 0; i < n; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ans++
		}
	}
	return ans
}
