/*
 * algo
 * Package: q3
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/4/30
 */

package main

import "fmt"

func main() {
	fmt.Println(countUnguarded(3, 3, [][]int{{1, 1}}, [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}))
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
}

type point struct {
	x, y int
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	visited := make(map[point]int, m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited[point{i, j}] = 0
		}
	}
	flag := false
	if m == 1 {
		flag = true
	}
	for _, wall := range walls {
		x, y := wall[0], wall[1]
		visited[point{x, y}] = 1
	}
	for z := len(guards)-1; z >= 0; z-- {
		x, y := guards[z][0], guards[z][1]
		visited[point{x, y}] = 1 // 将当前坐标标记为浏览过
		// 处理上下左右的情况，需要注意墙
		// 上
		for i := x-1; i >= 0; i-- {
			if visited[point{i, y}] == 1 {
				break
			}
			if flag {
				visited[point{i, y}] = 1
			} else {
				visited[point{i, y}] = 2
			}

		}
		// 下
		for i := x+1; i < m; i++ {
			if visited[point{i, y}] == 1 {
				break
			}
			if flag {
				visited[point{i, y}] = 1
			} else {
				visited[point{i, y}] = 2
			}
		}
		// 左
		for i := y-1; i >= 0; i-- {
			if visited[point{x, i}] == 1 {
				break
			}
			if flag {
				visited[point{x, i}] = 1
			} else {
				visited[point{x, i}] = 2
			}
		}
		// 右
		for i := y+1; i < n; i++ {
			if visited[point{x, i}] == 1 {
				break
			}
			if flag {
				visited[point{x, i}] = 1
			} else {
				visited[point{x, i}] = 2
			}
		}
	}
	ans := 0
	for _, v := range visited {
		if v == 0 {
			ans++
		}
	}
	return ans
}
