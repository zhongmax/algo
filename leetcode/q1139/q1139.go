package main

import "fmt"

// 1139. Largest 1-Bordered Square
// https://leetcode.cn/problems/largest-1-bordered-square/
// medium
func main() {
	case1 := [][]int{
		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
	}
	case2 := [][]int{
		{1, 1, 0, 0},
	}
	fmt.Println(largest1BorderedSquare(case1)) // 9
	fmt.Println(largest1BorderedSquare(case2)) // 1
}

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rows := make([][]int, m)
	for i := range rows {
		rows[i] = make([]int, n+1)
	}
	cols := make([][]int, n)
	for i := range cols {
		cols[i] = make([]int, m+1)
	}
	for i, row := range grid {
		for j, x := range row {
			rows[i][j+1] = rows[i][j] + x
			cols[j][i+1] = cols[j][i] + x
		}
	}
	fmt.Println(rows)
	fmt.Println(cols)
	for d := min(m, n); d > 0; d-- {
		for i := 0; i <= m-d; i++ {
			for j := 0; j <= n-d; j++ {
				if rows[i][j+d]-rows[i][j] == d && cols[j][i+d]-cols[j][i] == d && rows[i+d-1][j+d]-rows[i+d-1][j] == d && cols[j+d-1][i+d]-cols[j+d-1][i] == d {
					return d * d
				}
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
