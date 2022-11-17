/*
 * algo
 * Package: iv17_09
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/28
 */

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/get-kth-magic-number-lcci/
// 面试题 17.09. 第 k 个数
// medium
func main() {
	fmt.Println(getKthMagicNumber(5))
}

var factors = []int{3, 5, 7}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func getKthMagicNumber(k int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}
	magic := 0
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		magic = x
		if i == k {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
	return magic
}