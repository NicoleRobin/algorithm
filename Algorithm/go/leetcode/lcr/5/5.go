package main

import "fmt"

func maxProduct(words []string) int {
	var ans int

	return ans
}

func main() {
	testCases := []struct {
		words  []string
		expect int
	}{
		{
			words:  []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"},
			expect: 16,
		},
	}

	for i, testCase := range testCases {
		ans := maxProduct(testCase.words)
		if ans != testCase.expect {
			fmt.Printf("case %d fail, words:%v, expect:%d, ans:%d\n", i, testCase.words, testCase.expect, ans)
		} else {
			fmt.Printf("case %d pass\n", i)
		}
	}
}
