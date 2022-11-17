/*
 * algo
 * Package: q901
 * Author: maxwell
 * Email: peng_zhong@droidhang.com
 * Date: 2022/10/21
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	stock := Constructor()
	fmt.Println(stock.Next(100))
	fmt.Println(stock.Next(80))
	fmt.Println(stock.Next(60))
	fmt.Println(stock.Next(70))
	fmt.Println(stock.Next(60))
	fmt.Println(stock.Next(75))
	fmt.Println(stock.Next(85))
}

type StockSpanner struct {
	stack [][2]int
	idx int
}


func Constructor() StockSpanner {
	return StockSpanner{
		stack: [][2]int{{-1, math.MaxInt32}},
		idx:   -1,
	}
}


func (s *StockSpanner) Next(price int) int {
	s.idx++
	for price >= s.stack[len(s.stack)-1][1] {
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, [2]int{s.idx, price})
	return s.idx - s.stack[len(s.stack)-2][0]
}
