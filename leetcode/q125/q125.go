/*
 * algo
 * Package: q125
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/4/30
 */

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for !(unicode.IsLetter(rune(s[l])) || unicode.IsDigit(rune(s[l]))) {
			l++
		}
		for !(unicode.IsLetter(rune(s[r])) || unicode.IsDigit(rune(s[r]))) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
