package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/find-the-highest-altitude/
// 1732. 找到最高海拔
// easy
func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))         // 1
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2})) // 0
}

func largestAltitude(gain []int) int {
	origin := make([]int, len(gain)+1)
	origin[0] = 0
	for i, g := range gain {
		origin[i+1] = origin[i] + g
	}
	sort.Ints(origin)
	return origin[len(origin)-1]
}
