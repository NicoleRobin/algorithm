package main

<<<<<<< Updated upstream
import "github.com/nicolerobin/algorithm/go/binarytree"
=======
import (
	"fmt"
	"github.com/nicolerobin/algorithm/go/binarytree"
)
>>>>>>> Stashed changes

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *binarytree.TreeNode, targetSum int) int {
<<<<<<< Updated upstream
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func rootSum(root *binarytree.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var res int
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return res
}

func pathSum1(root *binarytree.TreeNode, targetSum int) (ans int) {
	preSum := map[int64]int{0: 1}
	var dfs func(*binarytree.TreeNode, int64)
	dfs = func(node *binarytree.TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr-int64(targetSum)]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
		return
	}
	dfs(root, 0)
	return
}

func main() {
=======
	var result int
	var traverse func(*binarytree.TreeNode, int, bool)
	traverse = func(nRoot *binarytree.TreeNode, beforeSum int, withBefore bool) {
		if nRoot == nil {
			return
		}
		if beforeSum+nRoot.Val == targetSum {
			result++
		}
		traverse(nRoot.Left, beforeSum+nRoot.Val, true)
		traverse(nRoot.Left, 0, true)

		traverse(nRoot.Right, beforeSum+nRoot.Val, true)
		traverse(nRoot.Right, 0, true)
	}
	traverse(root, 0, false)
	return result
}

func main() {
	root := binarytree.BuildTree([]int{1, -1, 2, -1, 3, -1, 4, -1, 5})
	binarytree.LevelPrintTree(root)
	result := pathSum(root, 3)
	fmt.Println(result)
>>>>>>> Stashed changes
}
