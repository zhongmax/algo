/*
 * algo
 * Package: q14
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/1
 */

package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0][0]
	n := 1
	for _, str := range strs {
		if len(str) > n && str[:n]
	}
}
