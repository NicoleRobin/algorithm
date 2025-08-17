package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth != 0 && rightDepth != 0 {
		if leftDepth < rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	} else if leftDepth == 0 {
		return rightDepth + 1
	} else {
		return leftDepth + 1
	}
}

func main() {}
