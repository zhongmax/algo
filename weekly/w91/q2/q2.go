package main

func main() {

}

func countGoodStrings(low int, high int, zero int, one int) int {
	var mod int = 1e9 + 7
	ans := 0
	zeroStr := ""
	oneStr := ""
	for i := 0; i < zero; i++ {
		zeroStr += "0"
	}
	for i := 0; i < one; i++ {
		oneStr += "1"
	}

	return ans % mod
}
