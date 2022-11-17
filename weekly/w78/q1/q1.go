/*
 * algo
 * Package: q1
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/14
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(divisorSubstrings(430043, 2))
}

func divisorSubstrings(num int, k int) int {
	ans := 0
	snum := strconv.Itoa(num)
	for i := 0; i <= len(snum)-k; i++ {
		n, _ := strconv.Atoi(snum[i:i+k])
		if n == 0 {
			continue
		}
		if num % n == 0 {
			ans++
		}
	}
	return ans
}
