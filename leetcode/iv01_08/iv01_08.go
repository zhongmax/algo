/*
 * algo
 * Package: iv01_08
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/30
 */

package main

import "fmt"

func main() {
	case1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(case1)
	fmt.Println(case1)
	case2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(case2)
	fmt.Println(case2)
}

type point struct {
	x, y int
}

func setZeroes(matrix [][]int)  {
	m, n := len(matrix), len(matrix[0])
	points := []point{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				points = append(points, point{i, j})
			}
		}
	}
	for _, p := range points {
		for i := 0; i < m; i++ {
			matrix[i][p.y] = 0
		}
		for j := 0; j < n; j++ {
			matrix[p.x][j] = 0
		}
	}
}