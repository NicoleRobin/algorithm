package main

import "fmt"

func maxProduct(words []string) int {
	var ans int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			duplicate := isDuplicate(words[i], words[j])
			if duplicate {
				continue
			}

			if ans < len(words[i])*len(words[j]) {
				ans = len(words[i]) * len(words[j])
			}
		}
	}

	return ans
}

func isDuplicate(a, b string) bool {
	var ans bool
	charAMap := map[byte]bool{}
	for i := 0; i < len(a); i++ {
		charAMap[a[i]] = true
	}
	for i := 0; i < len(b); i++ {
		if _, ok := charAMap[b[i]]; ok {
			ans = true
			break
		}
	}
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
