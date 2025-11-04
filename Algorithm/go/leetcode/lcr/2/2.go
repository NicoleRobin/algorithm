package main

func addBinary(a string, b string) string {
	var remainder int
	reverseA := reverse(a)
	reverseB := reverse(b)

	var ans []byte
	for i := 0; i < len(reverseA) || i < len(reverseB); i++ {
		var charA, charB int
		if i < len(reverseA) {
			charA = int(reverseA[i] - '0')
		}
		if i < len(reverseB) {
			charB = int(reverseB[i] - '0')
		}

		sum := charA + charB + remainder
		charSum := sum % 2
		remainder = sum / 2
		ans = append(ans, byte(charSum+'0'))
	}
	if remainder != 0 {
		ans = append(ans, byte(remainder+'0'))
	}
	return reverse(string(ans))
}

func reverse(s string) string {
	var result []byte
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return string(result)
}

func main() {}
