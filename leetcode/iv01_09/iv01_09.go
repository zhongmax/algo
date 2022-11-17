/*
 * algo
 * Package: iv01_09
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/29
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isFlipedString("waterbottlewaterbottle", "lewaterbott"))
}

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	return strings.Contains(s1+s1, s2)
}