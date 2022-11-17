/*
 * algo
 * Package: t2
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/10/16
 */

package main

import "fmt"

func main() {
	var m map[string]int

	delete(m, "1")
	fmt.Println(m)
}
