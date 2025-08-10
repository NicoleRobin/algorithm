package main

import "fmt"

/*
思考：
1、压缩字符串

思路：模拟题
1、只是需要注意处理边界和细节就可以了；
*/
func compressedString(word string) string {
	byteWord := []byte(word)
	count := 1

	var res []byte
	for i := 1; i < len(byteWord); i++ {
		if byteWord[i] == byteWord[i-1] {
			count++
			if count >= 9 {
				res = append(res, byte('0'+count), byteWord[i-1])
				count = 0
			}
		} else {
			if count > 0 {
				res = append(res, byte('0'+count), byteWord[i-1])
			}
			count = 1
		}
	}
	if count > 0 {
		res = append(res, byte('0'+count), byteWord[len(byteWord)-1])
	}

	return string(res)
}

func main() {
	testCases := []struct {
		word     string
		expected string
	}{
		{
			word:     "abcde",
			expected: "1a1b1c1d1e",
		},
		{
			word:     "aaaaaaaaaaaaaabb",
			expected: "9a5a2b",
		},
	}
	for i, tc := range testCases {
		res := compressedString(tc.word)
		if res != tc.expected {
			fmt.Printf("Case #%d fail, expected:%s, res:%s\n", i+1, tc.expected, res)
		} else {
			fmt.Printf("Case #%d pass, expected:%s, res:%s\n", i+1, tc.expected, res)
		}
	}
}
