/*
 * algo
 * Package: q1
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/1
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDigit("123", '3'))
}

func removeDigit(number string, digit byte) string {
	ans := []string{}
	for i, ch := range number {
		if byte(ch) == digit {
			newStr := number[:i] + number[i+1:]
			ans = append(ans, newStr)
		}
	}
	max := ans[0]
	for _, item := range ans[1:] {
		if item > max {
			max = item
		}
	}
	return max
}
