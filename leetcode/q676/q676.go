/*
 * algo
 * Package: q676
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/7/11
 */

package main

import (
	"fmt"
)

func main() {
	c := Constructor()
	c.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(c.Search("hello"))
	fmt.Println(c.Search("hhllo"))
}

type MagicDictionary []string


func Constructor() MagicDictionary {
	return []string{}
}


func (this *MagicDictionary) BuildDict(dictionary []string)  {
	*this = dictionary
}


func (this MagicDictionary) Search(searchWord string) bool {
	for _, item := range this {
		if compare(item, searchWord) {
			return true
		}
	}
	return false
}

func compare(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if s1[:i] == s2[:i] && s1[i+1:] == s2[i+1:] {
				return true
			}
		}
	}
	return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
