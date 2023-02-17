package main

import "fmt"

// https://leetcode.cn/problems/minimum-changes-to-make-alternating-binary-string/
// 1758. Minimum Changes To Make Alternating Binary String
// easy
func main() {
	fmt.Println(minOperations("0100")) // 1
	fmt.Println(minOperations("10"))   // 0
	fmt.Println(minOperations("1111")) // 2

}

func minOperations(s string) int {
	return -1
}
