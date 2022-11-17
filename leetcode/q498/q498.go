/*
 * algo
 * Package: q498
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/6/14
 */

package main

import "fmt"

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2}, {3, 4}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func findDiagonalOrder(mat [][]int) []int {
	ans := []int{}
	flag := true
	m, n := len(mat), len(mat[0])
	for i, j := 0, 0; i < m && j < n; {
		ans = append(ans, mat[i][j])
		if flag {
			if i == 0 && j+1 != n {
				flag = !flag
				j++
			} else if j+1 == n {
				flag = !flag
				i++
			} else {
				i--
				j++
			}
		} else {
			if i+1 == m {
				flag = !flag
				j++
			} else if j == 0 {
				flag = !flag
				i++
			} else {
				i++
				j--
			}
		}
	}
	return ans
}