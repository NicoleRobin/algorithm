package main

/*
思考：
完全没有思路，好像不是常见的题型
能想到的是hash，Key为颜色数量，Value为回答相同颜色兔子数量为Key值的个数

还学到一个技巧：
a/b的向上取整可以通过：(a + b -1)/b来实现；
*/
func numRabbits(answers []int) int {
	sameCountMap := map[int]int{}
	for _, answer := range answers {
		sameCountMap[answer]++
	}

	var ans int
	for answer, count := range sameCountMap {
		ans += (answer + count) / (answer + 1) * (answer + 1)
	}
	return ans
}

func main() {
	testCases := []struct {
		answers  []int
		expected int
	}{
		{
			answers:  []int{1, 1, 2},
			expected: 5,
		},
		{
			answers:  []int{10, 10, 10},
			expected: 11,
		},
	}

	for i, testCase := range testCases {
		result := numRabbits(testCase.answers)
		println("Test Case", i+1, "Result:", result, "Expected:", testCase.expected)
	}
}
