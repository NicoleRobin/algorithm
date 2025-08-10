package main

/*
思路：动态规划
1、这道题是经典的卡特兰数问题，表示有多少种不同的二叉搜索树可以用1到n的数字构建。
2、将问题拆解为子问题：对于每个数字i作为根节点，左子树可以由1到i-1的数字构成，右子树可以由i+1到n的数字构成。
*/
func numTrees(n int) int {
	g := make([]int, n+1)
	g[0], g[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}

func main() {

}
