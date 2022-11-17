/*
 * algo
 * Package: q1640
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/22
 */

package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.cn/problems/check-array-formation-through-concatenation/
// 1640. 能否连接形成数组
// easy
func main() {
	fmt.Println(canFormArray1([]int{15, 88}, [][]int{{88}, {15}}))
	fmt.Println(canFormArray1([]int{49, 18, 16}, [][]int{{16, 18, 49}}))
	fmt.Println(canFormArray1([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}))
}

func canFormArray(arr []int, pieces [][]int) bool {
	n, m := len(arr), len(pieces)
	for i := 0; i < n; {
		flag := false
		for j := 0; j < m; j++ {
			z := len(pieces[j])
			if arr[i] == pieces[j][0] {
				flag = true
				if reflect.DeepEqual(arr[i:i+z], pieces[j]) {
					i += z
					break
				} else {
					return false
				}
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func canFormArray1(arr []int, pieces [][]int) bool {
	mp := make(map[int]int, len(arr))
	for i, p := range pieces {
		mp[p[0]] = i
	}

	for i := 0; i < len(arr); {
		j, ok := mp[arr[i]]
		if !ok {
			return false
		}
		for _, x := range pieces[j] {
			if arr[i] != x {
				return false
			}
			i++
		}
	}
	return true
}