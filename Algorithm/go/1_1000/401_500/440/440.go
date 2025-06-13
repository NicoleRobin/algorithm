package main

import (
	"fmt"
)

/*
思路：字典树
1、将所有数字转换为字符串形式，插入到字典树中。
2、使用字典树的前序遍历来获取第k小的数字。
*/
func findKthNumber(n int, k int) int {
	var ans int
	dictTree := NewDictTreeNode(0)
	for i := 1; i <= n; i++ {
		dictTree.Insert(i)
	}
	dictTree.PreOrderPrint()

	var preTraverse func(root *DictTreeNode, k *int)
	preTraverse = func(root *DictTreeNode, k *int) {
		if root == nil || *k <= 0 {
			return
		}
		if root.exist {
			*k--
			if *k == 0 {
				ans = root.value
				return
			}
		}
		for _, child := range root.children {
			preTraverse(child, k)
			if ans != 0 { // 如果已经找到第k小的数字，直接返回
				return
			}
		}
	}
	preTraverse(dictTree, &k)

	return ans
}

type DictTreeNode struct {
	char     rune
	exist    bool
	value    int
	children [10]*DictTreeNode
}

func NewDictTreeNode(char rune) *DictTreeNode {
	return &DictTreeNode{
		char:     char,
		children: [10]*DictTreeNode{},
	}
}

func (node *DictTreeNode) Insert(value int) {
	strValue := fmt.Sprintf("%d", value)
	current := node
	for _, char := range strValue {
		charIndex := char - '0'
		if current.children[charIndex] == nil {
			current.children[charIndex] = NewDictTreeNode(char)
		}
		current = current.children[charIndex]
	}
	current.exist = true
	current.value = value
}

func (node *DictTreeNode) PreOrderPrint() {
	if node == nil {
		return
	}
	for _, child := range node.children {
		if child == nil {
			continue
		}
		if child.exist {
			fmt.Printf("%d ", child.value)
		}
		child.PreOrderPrint()
	}
}

func findKthNumber1(n int, k int) int {
	// 计算以 prefix 为根的子树中有多少个合法数字（不超过 n）
	getCount := func(prefix int, n int) int {
		count := 0
		curr := prefix
		next := prefix + 1
		for curr <= n {
			count += min(n+1, next) - curr
			curr *= 10
			next *= 10
		}
		return count
	}

	curr := 1
	k -= 1 // 从 1 开始已经是第一个，所以 k--

	for k > 0 {
		count := getCount(curr, n)
		if count <= k {
			// 跳过当前前缀子树，去下一个兄弟
			curr += 1
			k -= count
		} else {
			// 进入子树，往下一层
			curr *= 10
			k -= 1
		}
	}

	return curr
}

func main() {
	testCases := []struct {
		n, k, expected int
	}{
		{
			n:        13,
			k:        2,
			expected: 10,
		},
		{
			n:        2,
			k:        2,
			expected: 2,
		},
	}

	for i, tc := range testCases {
		ans := findKthNumber(tc.n, tc.k)
		fmt.Println("i:", i, ans, "expected:", tc.expected)
	}
}
