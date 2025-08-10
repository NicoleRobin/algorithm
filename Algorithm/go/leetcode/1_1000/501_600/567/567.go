package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	charCount := [26]int{}
	for i, ch := range s1 {
		charCount[ch-'a']++
		charCount[s2[i]-'a']--
	}

	differ := 0
	for _, count := range charCount {
		if count != 0 {
			differ++
		}
	}

	if differ == 0 {
		return true
	}

	for i, ch := range s2[:len(s2)-len(s1)] {
		fmt.Printf("i:%d, ch:%s, charCount:%+v\n", i, string(ch), charCount)
		if charCount[ch-'a'] == -1 {
			differ--
		} else if charCount[ch-'a'] == 0 {
			differ++
		}
		charCount[ch-'a']++

		if charCount[s2[i+len(s1)]-'a'] == 1 {
			differ--
		} else if charCount[s2[i+len(s1)]-'a'] == 0 {
			differ++
		}
		charCount[s2[i+len(s1)]-'a']--

		if differ == 0 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("leetcode-567")
	s1, s2 := "ab", "eidbaooo"
	fmt.Printf("s1:%s, s2:%s\n", s1, s2)
	result := checkInclusion(s1, s2)
	fmt.Println(result)
}
