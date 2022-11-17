package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/maximum-units-on-a-truck/
// 1710.卡车上的最大单元数
// easy
func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))                                                          // 8
	fmt.Println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))                                                // 91
	fmt.Println(maximumUnits([][]int{{1, 3}, {5, 5}, {2, 5}, {4, 2}, {4, 1}, {3, 1}, {2, 2}, {1, 3}, {2, 5}, {3, 2}}, 35)) // 91
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes[:], func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	// fmt.Println(boxTypes)
	ans := 0
	for _, box := range boxTypes {
		num, size := box[0], box[1]
		if truckSize >= num {
			truckSize -= num
			ans += size * num
		} else {
			ans += size * truckSize
			truckSize = 0
		}
		if truckSize <= 0 {
			break
		}
	}
	return ans
}
