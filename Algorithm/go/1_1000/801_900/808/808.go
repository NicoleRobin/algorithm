package main

/*
思路：动态规划
1、状态定义：
dp[i][j]:
2、
*/
func soupServings(n int) float64 {
	var res float64
	if n > 4800 {
		return 1.0
	}
	m := (n + 24) / 25 // 单位化
	memo := make(map[[2]int]float64)

	var dfs func(a, b int) float64
	dfs = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 && b > 0 {
			return 1.0
		}
		key := [2]int{a, b}
		if v, ok := memo[key]; ok {
			return v
		}
		res := 0.25 * (dfs(a-4, b) + dfs(a-3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3))
		memo[key] = res
		return res
	}
	return dfs(m, m)
}

func main() {

}
