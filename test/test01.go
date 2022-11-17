/*
 * algo
 * Package: test
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/9/19
 */

package main

import "fmt"

type temp struct {

}

func (t *temp) Add(item int) *temp {
	fmt.Println(item)
	return &temp{}
}

func main() {
	tt := &temp{}
	defer tt.Add(1).Add(2)
	tt.Add(3)
}
