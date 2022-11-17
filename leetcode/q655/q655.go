/*
 * algo
 * Package: q655
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/22
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(calDepth(t1))
	// fmt.Println(printTree(t1))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func calDepth(node *TreeNode) int {
	h := 0
	if node.Left != nil {
		h = calDepth(node.Left) + 1
	}
	if node.Right != nil {
		h = max(h, calDepth(node.Right)+1)
	}
	return h
}

func printTree(root *TreeNode) [][]string {
	height := calDepth(root)
	m := height + 1
	n := 1<<m - 1
	ans := make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, r, c int) {
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, r+1, c-1<<(height-r-1))
		}
		if node.Right != nil {
			dfs(node.Right, r+1, c+1<<(height-r-1))
		}
	}
	dfs(root, 0, (n-1)/2)
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

