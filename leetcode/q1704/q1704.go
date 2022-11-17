package main

import "fmt"

func main() {
	fmt.Println(halvesAreAlike("book"))
	fmt.Println(halvesAreAlike("textbook"))
}

func halvesAreAlike(s string) bool {
	n := len(s)
	s1 := s[:n/2]
	s2 := s[n/2:]
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	m1, m2 := map[byte]int{}, map[byte]int{}
	for i := 0; i < n/2; i++ {
		m2[s2[i]]++
		m1[s1[i]]++
	}
	sum1, sum2 := 0, 0
	for _, vowel := range vowels {
		sum1 += m1[vowel]
		sum2 += m2[vowel]
	}
	return sum1 == sum2
}
