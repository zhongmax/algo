/*
 * algo
 * Package: q1
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/20
 */

package main

import "fmt"

func main() {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("WBWBBBW", 2))
	fmt.Println(minimumRecolors("BWWWBB", 6))
}

func minimumRecolors(blocks string, k int) int {
	mp := map[int]int{}
	for i := 0; i <= len(blocks) - k; i++ {
		count := 0
		for j := i; j < len(blocks) && j - i < k; j++ {
			if blocks[j] != 'B' {
				count++
			}
		}
		mp[i] = count
	}
	min := 105
	for _, v := range mp {
		if v < min {
			min = v
		}
	}
	return min
}
