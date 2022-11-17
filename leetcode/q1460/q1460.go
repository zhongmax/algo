/*
 * algo
 * Package: q1460
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/24
 */

package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
	fmt.Println(canBeEqual([]int{7}, []int{7}))
	fmt.Println(canBeEqual([]int{3, 7, 9}, []int{3, 7, 11}))
}

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	return reflect.DeepEqual(target, arr)
}
