package main

import (
	"fmt"
)

// https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/
// 1742. 盒子中小球的最大数量
// easy
func main() {
	fmt.Println(countBalls(1, 10))  // 2
	fmt.Println(countBalls(5, 15))  // 2
	fmt.Println(countBalls(19, 28)) // 2
}

func countBalls(lowLimit int, highLimit int) int {
	m := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		m[getSum(i)]++
		// fmt.Println(getSum(i))
	}
	max := -1
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		if n >= 10 {
			sum += n % 10
			n = n / 10
		} else {
			sum += n
			n = 0
		}
	}
	return sum
}
