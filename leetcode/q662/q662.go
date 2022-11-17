/*
 * algo
 * Package: q662
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/27
 */

package main

import "fmt"

func main() {
	test1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(widthOfBinaryTree(test1))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type pair struct {
	node *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	ans := 1
	q := []pair{{root, 1}}
	for q != nil {
		ans = max(ans, q[len(q)-1].index-q[0].index+1)
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.node.Left != nil {
				q = append(q, pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				q = append(q, pair{p.node.Right, p.index * 2 + 1})
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}