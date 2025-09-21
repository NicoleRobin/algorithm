package main

func minOperations(s string) int {
	ans := 0
	for _, char := range s {
		if char == 'a' {
			continue
		}
		if int('z'-char)+1 > ans {
			ans = int('z'-char) + 1
		}
	}
	return ans
}

func main() {
}
