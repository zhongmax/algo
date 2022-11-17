package main

import "fmt"

func main() {
	// fmt.Println(subarrayLCM([]int{3, 6, 2, 7, 1}, 6))
	fmt.Println(subarrayLCM([]int{2, 1, 1, 5}, 5))
}

func subarrayLCM(nums []int, k int) int {
	ans := 0
	m := map[int]bool{}
	for _, num := range nums {
		if num == k {
			m[num] = true
		}
		if k > num {
			if k%num == 0 {
				m[num] = true
			}
		} else {
			if num%k == 0 {
				m[num] = true
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			a := nums[i:j]
			flag := true
			for _, num := range a {
				if !m[num] {
					flag = false
				}
			}
			if flag {
				fmt.Println(a)
				ans++
			}
		}
	}
	return ans
}
