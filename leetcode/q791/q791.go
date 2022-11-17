package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/custom-sort-string/
// 791. 自定义字符串排序
// medium
func main() {
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println(customSortString("cbafg", "abcd"))
}

func customSortString(order string, s string) string {
	m := map[rune]int{}
	for i, o := range order {
		m[o] = i
	}
	indexList := []int{}
	notMatchList := []byte{}
	for _, item := range s {
		if data, ok := m[item]; ok {
			indexList = append(indexList, data)
		} else {
			notMatchList = append(notMatchList, byte(item))
		}
	}
	ansList := []byte{}
	sort.Ints(indexList)
	for _, idx := range indexList {
		ansList = append(ansList, order[idx])
	}
	ansList = append(ansList, notMatchList...)
	return string(ansList)
}
