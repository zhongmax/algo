/*
 * algo
 * Package: q38
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/4/30
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay1(4))
}

func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j-start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}

func countAndSay1(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay1(n-1)

	builder := &strings.Builder{}
	s, e := 0, 0
	for e < len(str) {
		for e < len(str) && str[s] == str[e] {
			e++
		}
		builder.WriteString(strconv.Itoa(e-s))
		builder.WriteByte(str[s])
		s = e
	}
	return builder.String()
}

