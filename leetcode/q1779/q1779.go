package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/
// 1779. Find Nearest Point That Has the Same X or Y Coordinate
// easy
func main() {
	fmt.Println(nearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}})) // 2
	fmt.Println(nearestValidPoint(3, 4, [][]int{{3, 4}}))                                 // 0
	fmt.Println(nearestValidPoint(3, 4, [][]int{{2, 3}}))                                 // -1
	fmt.Println(nearestValidPoint(1, 1, [][]int{{1, 1}, {1, 1}}))                         // 0
}

func nearestValidPoint(x int, y int, points [][]int) int {
	m := map[int]int{}
	ans := []int{}
	for i, p := range points {
		if p[0] == x || p[1] == y {
			t := abs(p[0]-x) + abs(p[1]-y)
			_, ok := m[t]
			if !ok {
				m[t] = i
				ans = append(ans, t)
			}
		}
	}
	if len(ans) == 0 {
		return -1
	}
	sort.Ints(ans)
	return m[ans[0]]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
