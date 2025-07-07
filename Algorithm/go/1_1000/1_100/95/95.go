package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：动态规划
1、定义状态：dp[i] 表示长度为i的BST的个数
2、状态转移方程：

	dp[i] = dp[i-1] + dp[i-2] + ... + dp[0]

3、初始化：dp[0] = 1
4、返回结果：dp[n]
*/
func generateTrees(n int) []*TreeNode {
	var ans []*TreeNode
	for i := 1; i <= n; i++ {
		// 以i为根节点，左子树为[1, i-1]，右子树为[i+1, n]

	}
	return ans
}

func main() {
}
