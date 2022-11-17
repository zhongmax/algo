package main

import "fmt"

// 764. 最大加号标志
// https://leetcode.cn/problems/largest-plus-sign/
// medium
func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
	fmt.Println(orderOfLargestPlusSign(1, [][]int{{0, 0}}))
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	grids := make([][]int, n)
	for i := 0; i < n; i++ {
		grids[i] = make([]int, n)
		for j := 0; j < n; j++ {
			grids[i][j] = 1
		}
	}
	for _, mine := range mines {
		x, y := mine[0], mine[1]
		grids[x][y] = 0
	}
	maxPlus := 0
	for i, grid := range grids {
		for j := 0; j < len(grid); j++ {
			tmpPlus := 0
			if grids[i][j] == 1 {
				tmpPlus = 1
				for z := 1; z <= n/2; z++ {
					if i-z >= 0 && grids[i-z][j] == 1 &&
						i+z < n && grids[i+z][j] == 1 &&
						j-z >= 0 && grids[i][j-z] == 1 &&
						j+z < n && grids[i][j+z] == 1 {
						tmpPlus++
					} else {
						break
					}
				}
			}
			if tmpPlus > maxPlus {
				maxPlus = tmpPlus
			}
		}
	}
	return maxPlus
}
