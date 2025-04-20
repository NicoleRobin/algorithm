package main

func numRabbits(answers []int) int {

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
	}

	for i, testCase := range testCases {
		result := numRabbits(testCase.answers)
		println("Test Case", i+1, "Result:", result, "Expected:", testCase.expected)
	}
}
