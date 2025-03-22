package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	numStrMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var backtrack func(start int, path string)

	var result []string
	backtrack = func(start int, path string) {
		if start == len(digits) {
			result = append(result, path)

			return
		}

		str := numStrMap[digits[start]]
		for j := 0; j < len(str); j++ {
			backtrack(start+1, path+string(str[j]))
		}
	}

	backtrack(0, "")
	return result
}

func main() {
	TestCases := []string{
		"23",
	}
	for _, testCase := range TestCases {
		result := letterCombinations(testCase)
		for _, row := range result {
			fmt.Println("row:", row)
		}
	}
}
