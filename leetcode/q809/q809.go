package main

import "fmt"

// https://leetcode.cn/problems/expressive-words/
// 809. 情感丰富的文字
// medium
func main() {
	fmt.Println(expressiveWords("heeellooo", []string{"hello", "hi", "helo"})) // 1
}

func expressiveWords(s string, words []string) int {
	// 双指针, si wj判断当前指向字符是否相同, 当字符不一致时
	// 判断前面一致的字符数, 如果si == sj则继续
	// 如果si sj不一致, 需要si > sj 并且 si >= 3
	ans := 0
	for _, word := range words {
		if isExpressive(s, word) {
			ans++
		}
	}
	return ans
}

func isExpressive(s, w string) bool {
	n, m := len(s), len(w)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] != w[j] {
			return false
		}
		ch := s[i]
		cnti := 0
		for i < n && s[i] == ch {
			cnti++
			i++
		}
		cntj := 0
		for j < m && w[j] == ch {
			cntj++
			j++
		}
		if cntj > cnti || cntj < cnti && cnti < 3 {
			return false
		}
	}
	return i == n && j == m
}
