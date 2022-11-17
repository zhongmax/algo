/*
 * algo
 * Package: q944
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/12
 */

package main

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"cba","daf","ghi"}))
}

func minDeletionSize(strs []string) int {
	ans := 0
	for i := 0; i < len(strs); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] > strs[j][i] {
				ans++
				break
			}
		}
	}
	return ans
}
