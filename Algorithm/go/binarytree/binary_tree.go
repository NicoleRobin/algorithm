package binarytree

import (
	"fmt"

	"github.com/nicolerobin/algorithm/go/math"
)

func PrePrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	PrePrintTree(root.Left)
	fmt.Println(root.Val)
	PrePrintTree(root.Right)
}

func MidPrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PrePrintTree(root.Left)
	PrePrintTree(root.Right)
}

func LevelPrintTree(root *TreeNode) {
	if root == nil {
		return
	}

	var allLevelPrintingNodes [][]*TreeNode
	printingNodes := []*TreeNode{root}
	for len(printingNodes) > 0 {
		allLevelPrintingNodes = append(allLevelPrintingNodes, printingNodes)
		var tmpNodes []*TreeNode
		for _, node := range printingNodes {
			if node != nil {
				// fmt.Printf("%d", node.Val)
				tmpNodes = append(tmpNodes, node.Left)
				tmpNodes = append(tmpNodes, node.Right)
			}
			// fmt.Printf("\t")
		}
		// fmt.Println()
		printingNodes = tmpNodes
	}

	for i := 0; i < len(allLevelPrintingNodes); i++ {
		for j := i; j < len(allLevelPrintingNodes); j++ {
			fmt.Printf("\t")
		}

		for j := 0; j < len(allLevelPrintingNodes[i]); j++ {
			if allLevelPrintingNodes[i][j] != nil {
				fmt.Printf("%d", allLevelPrintingNodes[i][j].Val)
			} else {
				fmt.Printf(" ")
			}
			fmt.Printf("\t")
		}
		fmt.Println()
	}
}

func LastPrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	PrePrintTree(root.Left)
	PrePrintTree(root.Right)
	fmt.Println(root.Val)
}

func BuildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}
	nums = nums[1:]

	buildingNodes := []*TreeNode{root}
	// for len(buildingNodes) > 0 {
	for len(nums) > 0 {
		currentNode := buildingNodes[0]
		if currentNode.Left == nil {
			if nums[0] != -1 {
				currentNode.Left = &TreeNode{
					Val: nums[0],
				}
				buildingNodes = append(buildingNodes, currentNode.Left)
			}
			nums = nums[1:]
		}
		if len(nums) > 0 {
			if currentNode.Right == nil {
				if nums[0] != -1 {
					currentNode.Right = &TreeNode{
						Val: nums[0],
					}
					buildingNodes = append(buildingNodes, currentNode.Right)
				}
				nums = nums[1:]
			}
		}
		buildingNodes = buildingNodes[1:]
	}

	return root
}

func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := Height(root.Left)
	rightHeight := Height(root.Right)
	return math.Max(leftHeight, rightHeight) + 1
}
