package main

import (
	"fmt"

	"github.com/nicolerobin/algorithm/go/binarytree"
	"github.com/nicolerobin/algorithm/go/math"
)

func maxDepth(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return math.Max(left, right) + 1
}

func main() {
	fmt.Println("vim-go")
	nusm := []int{3, 9, 20, -1, -1, 15, 7}
	root := binarytree.BuildTree(nusm)
	maxDepth := maxDepth(root)
	fmt.Printf("maxDepth:%d\n", maxDepth)
}
