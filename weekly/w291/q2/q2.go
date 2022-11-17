/*
 * algo
 * Package: q2
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/5/1
 */

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimumCardPickup1([]int{3, 4, 2, 3, 4, 7}))
	fmt.Println(minimumCardPickup1([]int{1, 0 , 5, 3}))
	fmt.Println(minimumCardPickup1([]int{70,37,70,41,1,62,71,49,38,50,29,76,29,41,22,66,88,18,85,53}))
}

func minimumCardPickup(cards []int) int {
	min := math.MaxInt
	for i := 0; i < len(cards); i++ {
		e := i
		for e < len(cards)-1 {
			e++
			if e-i+1 > min {
				break
			}
			if cards[i] == cards[e] && e-i+1 < min {
				min = e-i+1
			}
		}
	}
	if min == math.MaxInt {
		return -1
	}
 	return min
}

func minimumCardPickup1(cards []int) int {
	min := math.MaxInt
	mp := map[int][]int{}
	for i := 0; i < len(cards); i++ {
		mp[cards[i]] = append(mp[cards[i]], i)
	}
	for _, v := range mp {
		if len(v) < 2 {
			continue
		}
		sort.Ints(v)
		if v[1]-v[0] < min {
			min = v[1]-v[0]
		}
	}
	if min == math.MaxInt {
		return -1
	}
	return min+1
}
