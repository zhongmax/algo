/*
 * algo
 * Package: q2
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/8
 */

package main

import "fmt"

func main() {
	root := TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				1, nil, nil,
			},
		},
		Right: &TreeNode{
			5,
			nil,
			&TreeNode{6, nil, nil},
		},
	}
	fmt.Println(averageOfSubtree(&root))
	// root := &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}}
	// fmt.Println(averageOfSubtree(root))
	// root := &TreeNode{
	// 	Val:   1,
	// 	Left:  nil,
	// 	Right: nil,
	// }
	// fmt.Println(averageOfSubtree(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	ans := 0
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var bfs func(node *TreeNode, count int, origin int) int
	bfs = func(node *TreeNode, count int, origin int) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return 1
		}
		count += node.Val
		bfs(node.Left, count, node.Val)
		bfs(node.Right, count, node.Val)
		if origin == count / n {
			ans++
		}
	}
	bfs(root, 0, root.Val)
	return ans
}

func inorder(node *TreeNode, count, n, origin int) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	if node.Left != nil {
		count += node.Val
		n++
		inorder(node.Left, count, n, origin)
	}
	return inorder(node.Left, count, n, origin) +
}
