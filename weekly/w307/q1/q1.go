/*
 * algo
 * Package: q1
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/8/21
 */

package main

import "fmt"

func main() {
	fmt.Println(minNumberOfHours(5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1}))
	fmt.Println(minNumberOfHours(5, 3, []int{1, 4}, []int{2, 5}))
	fmt.Println(minNumberOfHours(1, 1, []int{1, 1, 1, 1}, []int{1, 1, 1, 50}))
}

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	n := len(energy)
	ans := 0
	totalEnergy := 0
	for i := 0; i < n; i++ {
		totalEnergy += energy[i]
		if initialExperience  < experience[i] + 1  {
			ans += (experience[i] - initialExperience + 1)
			initialExperience += (experience[i] - initialExperience + 1)
		}
		initialExperience += experience[i]
	}
	if totalEnergy + 1 > initialEnergy {
		ans += (totalEnergy - initialEnergy + 1)
	}
	return ans
}
/*
1+1
*/
