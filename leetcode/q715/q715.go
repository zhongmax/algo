/*
 * algo
 * Package: q715
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/6/20
 */

package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	r := Constructor()
	r.AddRange(10, 20)
	r.RemoveRange(14, 16)
	fmt.Println(r.QueryRange(10, 14))
	fmt.Println(r.QueryRange(13, 15))
	fmt.Println(r.QueryRange(16, 17))
}

type RangeModule struct {
	*redblacktree.Tree
}


func Constructor() RangeModule {
	return RangeModule{redblacktree.NewWithIntComparator()}
}


func (t *RangeModule) AddRange(left int, right int)  {
	if node, ok := t.Floor(left); ok {
		r := node.Value.(int)
		if r >= right {
			return
		}
		if r >= left {
			left = node.Key.(int)
			t.Remove(left)
		}
	}
	for node, ok := t.Ceiling(left); ok && node.Key.(int) <= right; node, ok = t.Ceiling(left) {
		right = max(right, node.Value.(int))
		t.Remove(node.Key)
	}
	t.Put(left, right)
}


func (t *RangeModule) QueryRange(left int, right int) bool {
	node, ok := t.Floor(left)
	return ok && node.Value.(int) >= right
}


func (t *RangeModule) RemoveRange(left int, right int)  {
	if node, ok := t.Floor(left); ok {
		l, r := node.Key.(int), node.Value.(int)
		if r >= right {
			if l == left {
				t.Remove(l)
			} else {
				node.Value = left
			}
			if right != r {
				t.Put(right, r)
			}
			return
 		}
 		if r > left {
 			node.Value = left
		}
	}
	for node, ok := t.Ceiling(left); ok && node.Key.(int) < right; node, ok = t.Ceiling(left) {
		r := node.Value.(int)
		t.Remove(node.Key)
		if r > right {
			t.Put(right, r)
			break
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
