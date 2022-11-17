package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(distinctAverages([]int{6, 83, 67, 2, 41, 82, 26, 1, 76, 62, 41, 69, 47, 56, 79, 55, 54, 94, 56, 42, 64, 31, 36, 100, 54, 86, 97, 55, 89, 88, 0, 40, 59, 15, 64, 18, 31, 0, 25, 22, 3, 35, 27, 11}))
}

func distinctAverages(nums []int) int {
	n := len(nums)
	m := map[float64]int{}
	sort.Ints(nums)
	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[len(nums)-(i+1)]
		// fmt.Println(float64(a+b) / 2)
		m[float64(a+b)/2]++
	}
	return len(m)
}
