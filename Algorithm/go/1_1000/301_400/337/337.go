package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：
1、选当前节点：左右儿子都不能选；不选当前节点：左右儿子可选可不选
*/
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		leftRob, leftNotRob := dfs(root.Left)
		rightRob, rightNotRob := dfs(root.Right)

		return root.Val + leftNotRob + rightNotRob, max(leftRob, leftNotRob) + max(rightRob, rightNotRob)
	}
	rootRob, rootNotRob := dfs(root)

	return max(rootRob, rootNotRob)
}

func main() {
}
