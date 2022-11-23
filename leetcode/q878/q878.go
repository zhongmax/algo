package main

import "fmt"

// https://leetcode.cn/problems/nth-magical-number/
// 878. 第 N 个神奇数字
// hard
func main() {
	fmt.Println(nthMagicalNumber(1, 2, 3)) // 2
	fmt.Println(nthMagicalNumber(4, 2, 3)) // 6
}

func nthMagicalNumber(n int, a int, b int) int {
	var mod int = 1e9 + 7
	// 获取最小公倍数
	lcm := a / gcd(a, b) * b
	l, r := 0, min(a, b)*n
	// 二分
	for l+1 < r {
		m := l + (r-l)/2
		if m/a+m/b-m/lcm >= n {
			r = m
		} else {
			l = m
		}
	}
	return r % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
