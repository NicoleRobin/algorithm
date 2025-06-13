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
/*
思路：
1、遍历整棵树
2、当遇到和key相同的节点时，就考虑删除该节点
3、删除该节点后需要合并左右子树，当左右子树其一为空时则直接返回另一半做为新的子树即可
4、当左右子树均不为空时则有两种处理方式：将右子树挂在左子树的最右节点的右子树上；将左子树挂在右子树的最左节点的左子树上。

时间复杂度：O(n)
*/
func deleteNode1(root *binarytree.TreeNode, key int) *binarytree.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		newRoot := root.Left
		pNode := newRoot
		for pNode.Right != nil {
			pNode = pNode.Right
		}
		pNode.Right = root.Right
		return newRoot
	}

	root.Left = deleteNode1(root.Left, key)
	root.Right = deleteNode1(root.Right, key)
	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
思路：
1、遍历二叉树
2、当前节点小于key时走右子树
3、当前节点大于key时走左子树
4、当前节点等于key时找到节点进行处理
*/
func deleteNode2(root *binarytree.TreeNode, key int) *binarytree.TreeNode {
	var pre, newRoot *binarytree.TreeNode
	left := false
	pNode := root
	for pNode != nil {
		if pNode.Val == key {
			if pNode.Left == nil {
				newRoot = pNode.Right
			} else if pNode.Right == nil {
				newRoot = pNode.Left
			} else {
				newRoot = pNode.Left
				pTmp := newRoot
				for pTmp.Right != nil {
					pTmp = pTmp.Right
				}
				pTmp.Right = pNode.Right
			}

			if pre == nil {
				return newRoot
			} else {
				if left == true {
					pre.Left = newRoot
				} else {
					pre.Right = newRoot
				}
			}
			break
		} else if pNode.Val > key {
			pre = pNode
			left = true
			pNode = pNode.Left
		} else {
			pre = pNode
			left = false
			pNode = pNode.Right
		}
	}
	return root
}

func main() {
	fmt.Println("vim-go")
	nums := []int{5, 3, 6, 2, 4, -1, 7}
	root := binarytree.BuildTree(nums)
	binarytree.LevelPrintTree(root)
	newRoot := deleteNode2(root, 5)
	binarytree.LevelPrintTree(newRoot)
}
