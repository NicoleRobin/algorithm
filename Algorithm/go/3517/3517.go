package main

import "fmt"

func smallestPalindrome(s string) string {
	var charCount [26]int
	for i := 0; i < len(s); i++ {
		charCount[s[i]-'a']++
	}

	var lHalf []byte
	var oddChar byte
	for i := 0; i < 26; i++ {
		if charCount[i]%2 == 1 {
			oddChar = byte(i + 'a')
		}
		for j := 0; j < charCount[i]/2; j++ {
			lHalf = append(lHalf, byte(i+'a'))
		}
	}
	result := make([]byte, len(s))
	for i := 0; i < len(lHalf); i++ {
		result[i] = lHalf[i]
		result[len(result)-1-i] = lHalf[i]
	}
	if len(s)%2 == 1 {
		result[len(result)/2] = oddChar
	}

	return string(result)
}

func main() {
	fmt.Println(smallestPalindrome("babab"))
}
