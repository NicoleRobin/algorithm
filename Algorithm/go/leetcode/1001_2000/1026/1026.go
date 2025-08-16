package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：遍历二叉树，并记录先前访问过的节点

优化思路：
1、不需要记录之前的所有节点，只需要记录之前的最大值和最小值即可；
*/
func maxAncestorDiff(root *TreeNode) int {
	var traverse func(*TreeNode, []int)

	var res int
	traverse = func(root *TreeNode, parents []int) {
		if root == nil {
			return
		}
		for _, parent := range parents {
			if int(math.Abs(float64(parent-root.Val))) > res {
				res = int(math.Abs(float64(parent - root.Val)))
			}
		}
		parents = append(parents, root.Val)
		traverse(root.Left, parents)
		traverse(root.Right, parents)
		parents = parents[:len(parents)-1]
	}
	traverse(root, []int{})

	return res
}

/*
思路：二叉树
优化思路：
1、不需要记录之前的所有节点，只需要记录之前的最大值和最小值即可；
*/
func maxAncestorDiff1(root *TreeNode) int {
	var traverse func(*TreeNode, int, int)

	var res int
	traverse = func(newRoot *TreeNode, minVal, maxVal int) {
		if newRoot == nil {
			return
		}
		if newRoot != root {
			if int(math.Abs(float64(minVal-newRoot.Val))) > res {
				res = int(math.Abs(float64(minVal - newRoot.Val)))
			}
			if int(math.Abs(float64(maxVal-newRoot.Val))) > res {
				res = int(math.Abs(float64(maxVal - newRoot.Val)))
			}
		}
		// update minVal and maxVal
		if newRoot.Val > maxVal {
			maxVal = newRoot.Val
		}
		if newRoot.Val < minVal {
			minVal = newRoot.Val
		}
		traverse(newRoot.Left, minVal, maxVal)
		traverse(newRoot.Right, minVal, maxVal)
	}
	if root == nil {
		return 0
	}
	traverse(root, root.Val, root.Val)

	return res
}

func main() {
}
