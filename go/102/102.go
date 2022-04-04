package main

import (
	"fmt"

	"github.com/nicolerobin/algorithm/go/binarytree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *binarytree.TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	traversingNodes := []*binarytree.TreeNode{root}
	for len(traversingNodes) > 0 {
		currentLevelResult := []int{}
		nextLevelNodes := []*binarytree.TreeNode{}
		for _, node := range traversingNodes {
			currentLevelResult = append(currentLevelResult, node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		result = append(result, currentLevelResult)
		traversingNodes = nextLevelNodes
	}
	return result
}

func main() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := binarytree.BuildTree(nums)
	binarytree.LevelPrintTree(root)

	result := levelOrder(root)
	fmt.Println(result)
}
