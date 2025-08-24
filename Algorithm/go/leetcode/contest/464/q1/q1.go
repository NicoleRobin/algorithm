package main

func gcdOfOddEvenSums(n int) int {
	var sumOdd, sumEven int

	for i := 1; i <= 2*n; i++ {
		if i%2 == 1 {
			sumOdd += i
		} else {
			sumEven += i
		}
	}

	oddGcdList := gcd(sumOdd)
	evenGcdList := gcd(sumEven)
	var i, j int
	var ans int
	for i < len(oddGcdList) && j < len(evenGcdList) {
		if oddGcdList[i] == evenGcdList[j] {
			ans = oddGcdList[i]
			i++
			j++
		} else if oddGcdList[i] < evenGcdList[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func gcd(n int) []int {
	var ans []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {

}
