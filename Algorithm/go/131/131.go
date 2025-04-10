package main

import "fmt"

/*
思路：
1、回溯法+动态规划
2、判断是否是回文串
3、预处理，先使用动态规划生成一个二维数组，用于判断是否是回文串
4、回溯法，递归遍历字符串，判断是否是回文串
*/
func partition(s string) (ans [][]string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}

func main() {
	strs := []string{"aab", "a"}
	for _, str := range strs {
		fmt.Println(partition(str))
	}
}
