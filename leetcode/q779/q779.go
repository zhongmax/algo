/*
 * algo
 * Package: q779
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/10/20
 */

package main

import "fmt"

// https://leetcode.cn/problems/k-th-symbol-in-grammar/
// 779. 第K个语法符号
// medium
func main() {
	fmt.Println(kthGrammar2(1, 1))
	fmt.Println(kthGrammar2(2, 1))
	fmt.Println(kthGrammar2(2, 2))
}

func kthGrammar1(n int, k int) int {
	if n == 1 {
		return 0
	}
	return k & 1 ^ 1 ^ kthGrammar1(n-1, (k+1)/2)
}

func kthGrammar2(n int, k int) int {
	var dfs func(r, c, cur int) int
	dfs = func(r, c, cur int) int {
		if r == 1 {
			return cur
		}
		if (c%2 == 0 && cur == 0) || (c%2 == 1 && cur == 1) {
			return dfs(r-1, (c-1)/2+1, 1)
		} else {
			return dfs(r-1, (c-1)/2+1, 0)
		}
	}
	if dfs(n, k, 1) == 0 {
		return 1
	} else {
		return 0
	}
}

func kthGrammar(n int, k int) int {
	cnt := [][]int{}
	cnt = append(cnt, []int{0})
	for i := 1; i < n; i++ {
		last := cnt[i-1]
		c := make([]int, len(last)*2)
		for j := 0; j < len(last); j++ {
			if last[j] == 0 {
				c[j * 2] = 0
				c[j * 2 + 1] = 1
			} else {
				c[j * 2] = 1
				c[j * 2 + 1] = 0
			}
		}
		cnt = append(cnt, c)
	}
	// fmt.Println(cnt)
	return cnt[n-1][k-1]
}
