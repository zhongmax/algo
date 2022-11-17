/*
 * algo
 * Package: q827
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/18
 */

package main

import (
	"fmt"
	"time"
)

// https://leetcode.cn/problems/making-a-large-island/
// 827. 最大人工岛
// hard
func main() {
	// fmt.Println(largestIsland([][]int{{1, 0}, {0, 1}}))
	// fmt.Println(largestIsland1([][]int{{1, 1}, {1, 0}}))
	// fmt.Println(largestIsland([][]int{{1, 1}, {1, 0}}))
	fmt.Println(largestIsland1([][]int{{1, 0, 1}, {0, 1, 1}, {0, 1, 1}}))
	fmt.Println(largestIsland([][]int{{1, 0, 1}, {0, 1, 1}, {0, 1, 1}}))
}

func largestIsland(grid [][]int) int {
	startTime := time.Now()
	defer fmt.Println(time.Since(startTime))
	max := getIslandSize(0, 0, grid)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			} else {
				size := getIslandSize(i, j, grid)
				if size > max {
					max = size
				}
			}
		}
	}
	return max
}

type CustomIndex struct {
	x int
	y int
}

func getIslandSize(i int, j int, grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := map[CustomIndex]bool{}
	visited[CustomIndex{i, j}] = true
	count := 1
	q := []CustomIndex{{i, j}}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		if t.x - 1 >= 0 && grid[t.x-1][t.y] == 1 && !visited[CustomIndex{t.x-1, t.y}] {
			visited[CustomIndex{t.x-1, t.y}] = true
			q = append(q, CustomIndex{t.x-1, t.y})
			count++
		}
		if t.x + 1 < m && grid[t.x+1][t.y] == 1 && !visited[CustomIndex{t.x+1, t.y}] {
			visited[CustomIndex{t.x+1, t.y}] = true
			q = append(q, CustomIndex{t.x+1, t.y})
			count++
		}
		if t.y - 1 >= 0 && grid[t.x][t.y-1] == 1 && !visited[CustomIndex{t.x, t.y-1}] {
			visited[CustomIndex{t.x, t.y-1}] = true
			q = append(q, CustomIndex{t.x, t.y-1})
			count++
		}
		if t.y + 1 < n && grid[t.x][t.y+1] == 1 && !visited[CustomIndex{t.x, t.y+1}] {
			visited[CustomIndex{t.x, t.y+1}] = true
			q = append(q, CustomIndex{t.x, t.y+1})
			count++
		}
	}
	return count
}

func largestIsland1(grid [][]int) (ans int) {
	startTime := time.Now()
	defer fmt.Println(time.Since(startTime))
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n, t := len(grid), 0
	tag := make([][]int, n)
	for i := range tag {
		tag[i] = make([]int, n)
	}
	area := map[int]int{}
	var dfs func(int, int)
	dfs = func(i, j int) {
		tag[i][j] = t
		area[t]++
		for _, d := range dir4 {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] > 0 && tag[x][y] == 0 {
				dfs(x, y)
			}
		}
	}
	for i, row := range grid {
		for j, x := range row {
			if x > 0 && tag[i][j] == 0 { // 枚举没有访问过的陆地
				t = i*n + j + 1
				dfs(i, j)
				ans = max(ans, area[t])
			}
		}
	}

	for i, row := range grid {
		for j, x := range row {
			if x == 0 { // 枚举可以添加陆地的位置
				newArea := 1
				conn := map[int]bool{0: true}
				for _, d := range dir4 {
					x, y := i+d.x, j+d.y
					if 0 <= x && x < n && 0 <= y && y < n && !conn[tag[x][y]] {
						newArea += area[tag[x][y]]
						conn[tag[x][y]] = true
					}
				}
				ans = max(ans, newArea)
			}
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}