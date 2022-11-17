/*
 * algo
 * Package: q735
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/7/13
 */

package main

import "fmt"

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
	fmt.Println(asteroidCollision_lc([]int{10, 2, -5}))
}

func asteroidCollision_lc(asteroids []int) []int {
	ans := []int{}
	for _, at := range asteroids {
		alive := true
		for alive && at < 0 && len(ans) > 0 && ans[len(ans)-1] > 0 {
			alive = ans[len(ans)-1] < -at
			if ans[len(ans)-1] <= -at {
				ans = ans[:len(ans)-1]
			}
		}
		if alive {
			ans = append(ans, at)
		}
	}
	return ans
}

func asteroidCollision(asteroids []int) []int {
	allDirection := false
	for !allDirection {
		for i := 0; i < len(asteroids)-1; i++ {
			if asteroids[i] > 0 && asteroids[i+1] < 0 {
				if abs(asteroids[i]) > abs(asteroids[i+1]) {
					asteroids = append(asteroids[:i+1], asteroids[i+2:]...)
				} else if abs(asteroids[i]) < abs(asteroids[i+1]) {
					asteroids = append(asteroids[:i], asteroids[i+1:]...)
				} else {
					asteroids = append(asteroids[:i], asteroids[i+2:]...)
				}
			}
			if len(asteroids) == 0 {
				return []int{}
			}
			if checkArrIsSame(asteroids) {
				allDirection = true
			}
		}
	}
	return asteroids
}

func checkArrIsSame(arr []int) bool {
	if len(arr) == 0 {
		return true
	}
	mp := map[bool]struct{}{}
	for _, item := range arr {
		if item > 0 {
			mp[true] = struct{}{}
		} else {
			mp[false] = struct{}{}
		}
	}
	if len(mp) == 1 {
		return true
	} else {
		return  false
	}
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}