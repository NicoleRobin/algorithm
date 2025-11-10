package main

func countBits(n int) []int {
	var ans []int
	for i := 0; i <= n; i++ {
		ans = append(ans, countBit(i))
	}
	return ans
}

func countBit(n int) int {
	var ans int
	for i := 0; (2 << i) <= n; i++ {
		if n&(2<<i) == 1 {
			ans++
		}
	}
	return ans
}

func main() {

}
