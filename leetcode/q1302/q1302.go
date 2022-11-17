/*
 * algo
 * Package: q1302
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/17
 */

package main

import "fmt"

func main() {
	tn := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: nil,
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	fmt.Println(deepestLeavesSum(tn))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	levelNum := 0
	node := []*TreeNode{root}
	for len(node) > 0 {
		n := len(node)
		sum := 0
		for i := 0; i < n; i++ {
			h := node[0]
			node = node[1:]
			sum += h.Val
			if h.Left != nil {
				node = append(node, h.Left)
			}
			if h.Right != nil {
				node = append(node, h.Right)
			}
		}
		levelNum = sum
	}
	return levelNum
}
