/*
欢迎各位勇者来到力扣城，本次试炼主题为「二叉搜索树染色」。

每位勇士面前设有一个二叉搜索树的模型，模型的根节点为 root，树上的各个节点值均不重复。初始时，所有节点均为蓝色。现在按顺序对这棵二叉树进行若干次操作， ops[i] = [type, x, y] 表示第 i 次操作为：

type 等于 0 时，将节点值范围在 [x, y] 的节点均染蓝
type 等于 1 时，将节点值范围在 [x, y] 的节点均染红
请返回完成所有染色后，该二叉树中红色节点的数量。

注意：

题目保证对于每个操作的 x、y 值定出现在二叉搜索树节点中

 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
/*
操作个数的范围为10^9，因此该算法时间复杂度必须小于等于O(n)
求区间的问题，和之前遇到的给出一段账号上下线时间计算最大在线用户个数很像
*/
package main

const (
	OP_TYPE_BLUE = 0
	OP_TYPE_RED  = 1
)

func getNumber(root *TreeNode, ops [][]int) int {
	nums := []int{}
	var preTraverse func(root *TreeNode)
	preTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		preTraverse(root.Left)
		nums = append(nums, root.Val)
		preTraverse(root.Right)
	}

	redCount := 0
	for _, op := range ops {
		opType := op[0]
		opMin := op[1]
		opMax := op[2]
		if opType == OP_TYPE_BLUE {

		} else if opOptye == OP_TYPE_RED {

		} else {

		}
	}
}
