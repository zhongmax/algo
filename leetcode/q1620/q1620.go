/*
 * algo
 * Package: q1620
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/11/2
 */

package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/coordinate-with-maximum-network-quality/
// 1620. 网络信号最好的坐标
// medium
func main() {
	fmt.Println(bestCoordinate([][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, 2)) // [2, 1]
	fmt.Println(bestCoordinate([][]int{{23, 11, 21}}, 9)) // [23, 11]
	fmt.Println(bestCoordinate([][]int{{1, 2, 13}, {2, 1, 7}, {0, 1, 9}}, 9)) // [1, 2]
}

func bestCoordinate(towers [][]int, radius int) []int {
	var g [110][110]int
	x, y, val := 0, 0, 0
	for _, tower := range towers {
		a, b, k := tower[0], tower[1], tower[2]
		si, sj := 0, 0
		if a - radius > 0 {
			si = a - radius
		}
		if b - radius > 0 {
			sj = b - radius
		}
		for i := si; i <= a + radius; i++ {
			for j := sj; j <= b + radius; j++ {
				d := math.Sqrt(float64((a-i)*(a-i)+(b-j)*(b-j)))
				if d > float64(radius) {
					continue
				}
				g[i][j] += int(float64(k) / (1 + d))
				if g[i][j] > val {
					x = i
					y = j
					val = g[i][j]
				} else if g[i][j] == val && (i < x || (i == x && j < y)) {
					x = i
					y = j
				}
			}
		}
	}
	return []int{x, y}
}