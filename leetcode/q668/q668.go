/*
 * algo
 * Package: q668
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/18
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findKthNumber(3, 3, 5))
}

func findKthNumber(m int, n int, k int) int {
	// arr := make([][]int, m)
	// for i := range arr {
	// 	arr[i] = make([]int, n)
	// }
	arr := make([]int, m * n)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			arr[(i-1)*n+j-1] = i * j
		}
	}
	sort.Ints(arr)
	return arr[k-1]
}
