package main

import "fmt"

func decodeString(s string) string {
	var result []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			// 出栈，直到遇见左括号
			subStrStart := len(result) - 1
			for subStrStart >= 0 && result[subStrStart] != '[' {
				subStrStart--
			}
			subStr := make([]byte, len(result[subStrStart+1:]))
			copy(subStr, result[subStrStart+1:])
			result = result[:subStrStart]

			// 继续出栈找到数字
			count := 1
			var subStrCount int
			for len(result) > 0 && (result[len(result)-1] >= '0' && result[len(result)-1] <= '9') {
				subStrCount += count * int(result[len(result)-1]-'0')
				result = result[:len(result)-1]
				count *= 10
			}
			for j := 0; j < subStrCount; j++ {
				result = append(result, subStr...)
			}
		} else {
			// 直接入栈
			result = append(result, s[i])
		}
	}
	return string(result)
}

func main() {
	s := "3[a]2[bc]"
	result := decodeString(s)
	fmt.Println(result)
}
