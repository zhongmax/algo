/*
 * algo
 * Package: q646
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/3
 */

// https://leetcode.cn/problems/maximum-length-of-pair-chain/
// 646. 最长数对链
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
	fmt.Println(findLongestChain([][]int{{2, 3}, {1, 2}, {3, 4}, {2, 10}}))
}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][0]
	})
	fmt.Println(pairs)
	return -1
}
