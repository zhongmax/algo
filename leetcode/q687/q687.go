/*
 * algo
 * Package: q687
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/2
 */

// https://leetcode.cn/problems/longest-univalue-path/
// 687. 最长同值路径
package main

import "fmt"

func main() {
	test1 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(longestUnivaluePath(test1))
	test2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(longestUnivaluePath(test2))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var dfs func(*TreeNode) int
	ans := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		left1, right1 := 0, 0
		if node.Left != nil && node.Val == node.Left.Val {
			left1 = left + 1
		}
		if node.Right != nil && node.Val == node.Right.Val {
			right1 = right + 1
		}
		ans = max(ans, left1+right1)
		return max(left1, right1)
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}