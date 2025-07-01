package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：
1、
*/
func rob(root *TreeNode) int {
	var ans int
	if root == nil {
		return 0
	}
	left := rob(root.Left)
	right := rob(root.Right)

	return max(left+right, root.Val+rob(root.Left.Left)+rob(root.Left.Right)+rob(root.Right.Left)+rob(root.Right.Right))
}

func main() {

}
