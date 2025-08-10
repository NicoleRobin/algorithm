package main

/*
思路：
1、回溯算法
2、递归函数 backtrack(open int, close int, path string)

注意：
1、open < n 时，可以添加左括号
2、close < open 时，可以添加右括号
*/
func generateParenthesis(n int) []string {
	// open: 左括号个数 close: 又括号个数
	var backtrack func(open, close int, path string)

	var result []string
	backtrack = func(open, close int, path string) {
		if len(path) == 2*n {
			result = append(result, path)
			return
		}
		if open < n {
			backtrack(open+1, close, path+"(")
		}
		if close < open {
			backtrack(open, close+1, path+")")
		}
	}
	backtrack(0, 0, "")
	return result
}

func main() {
	for i := 1; i < 4; i++ {
		result := generateParenthesis(i)
		for _, row := range result {
			println("row:", row)
		}
	}
}
