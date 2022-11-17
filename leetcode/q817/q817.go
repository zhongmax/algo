/*
 * algo
 * Package: q817
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/10/12
 */

package main

import "fmt"

// https://leetcode.cn/problems/linked-list-components/
// 817. 链表组件
// meddle
func main() {
	case1 := &ListNode{0,  &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}
	case2 := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}}
	fmt.Println(numComponents(case1, []int{0, 1, 3}))
	fmt.Println(numComponents(case2, []int{0, 3, 1, 4}))
}

type ListNode struct {
	Val int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	ans := 0
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}
	for inSet := false; head != nil; head = head.Next {
		if _, ok := set[head.Val]; !ok {
			inSet = false
		} else if !inSet {
			inSet = true
			ans++
		}
	}
	return ans
}