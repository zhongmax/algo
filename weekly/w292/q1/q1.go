/*
 * algo
 * Package: q1
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/8
 */

package main

import "fmt"

func main() {
	fmt.Println(largestGoodInteger("1221000"))
}

func largestGoodInteger(num string) string {
	if len(num) == 3 && num[0] == num[1] && num[1] == num[2] {
		return num
	}
	max := ""
	for i := 0; i < len(num) - 2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			if max == "" {
				max = num[i:i+3]
			} else if max[0] < num[i] {
				max = num[i:i+3]
			}
		}
	}
	return max
}

