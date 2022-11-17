/*
 * algo
 * Package: q513
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/6/22
 */

package main

import "fmt"

func main() {
	test1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	test2 := &TreeNode{
		1,
		&TreeNode{2, &TreeNode{4, nil, nil}, nil},
		&TreeNode{3, &TreeNode{
			Val:   5,
			Left:  &TreeNode{7, nil, nil},
			Right: nil,
		}, &TreeNode{
			Val:   6,
			Left:  nil,
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		}},
	}
	fmt.Println("dfs")
	fmt.Println(findBottomLeftValue(test1))
	fmt.Println(findBottomLeftValue(test2))
	fmt.Println("bfs")
	fmt.Println(findBottomLeftValue2(test1))
	fmt.Println(findBottomLeftValue2(test2))
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

// dfs
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	currentHeight := 0
	currentValue := 0
	var dfs func(node *TreeNode, height int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
		if height > currentHeight {
			currentHeight = height
			currentValue = node.Val
		}
	}
	dfs(root, 0)
	return currentValue
}

// bfs
func findBottomLeftValue2(root *TreeNode) int {
	ans := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		ans = node.Val
	}
	return ans
}
