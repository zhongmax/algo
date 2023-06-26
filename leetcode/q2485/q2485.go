package main

import "fmt"

func main() {
	case1 := 8
	// ans1 := 6
	case2 := 1
	// ans2 := 1
	case3 := 4
	// ans3 := -1
	fmt.Println(pivotInteger(case1))
	fmt.Println(pivotInteger(case2))
	fmt.Println(pivotInteger(case3))
}

// https://leetcode.cn/problems/find-the-pivot-integer/submissions/
// 2485. 找出中枢整数
// easy
func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}
	m := map[int]int{}
	for i := 1; i <= n; i++ {
		m[i] = m[i-1] + i
	}
	// fmt.Println(m)
	for i := 1; i <= n; i++ {
		if m[n]-m[i] == m[i]-i {
			return i
		}
	}
	return -1
}
