package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	s = reverseStr(s)
	begin := 0
	newS := ""
	for i := 0; i < len(s); i++ {
		if s[i] == byte(' ') {
			if i > begin {
				word := string(s[begin:i])
				word = reverseStr(word)
				if newS == "" {
					newS = word
				} else {
					newS = newS + " " + word
				}
			}
			begin = i + 1
		}
	}
	if len(s) > begin {
		lastWord := string(s[begin:len(s)])
		lastWord = reverseStr(lastWord)
		if newS == "" {
			newS = lastWord
		} else {
			newS = newS + " " + lastWord
		}
	}
	return newS
}

func reverseStr(s string) string {
	newS := ""
	for i := len(s) - 1; i >= 0; i-- {
		newS = newS + string(s[i])
	}
	return newS
}

func main() {
	str := "the sky is blue"
	reverseStr := reverseWords(str)
	fmt.Printf("str:%s, reverseStr:%s\n", str, reverseStr)
}
