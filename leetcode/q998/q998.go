/*
 * algo
 * Package: q998
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/30
 */

package main

import "fmt"

func main() {
	fmt.Println(insertIntoMaxTree(&TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}, 5))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	for cur := root; cur != nil; cur = cur.Right {
		if val > cur.Val {
			if parent == nil {
				return &TreeNode{val, root, nil}
			}
			parent.Right = &TreeNode{val, cur, nil}
			return root
		}
		parent = cur
	}
	parent.Right = &TreeNode{Val: val}
	return root
}
