/*
 * algo
 * Package: q719
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/6/15
 */

package main

import "sort"

func main() {

}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
		cnt := 0
		for j, num := range nums {
			i := sort.SearchInts(nums[:j], num-mid)
			cnt += j - i
		}
		return cnt >= k
	})
}
