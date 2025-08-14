package main

import "fmt"

/*
思路：暴力
*/
func numOfUnplacedFruits1(fruits []int, baskets []int) int {
	res := len(fruits)
	for _, fruit := range fruits {
		for j, basket := range baskets {
			if fruit <= basket {
				baskets[j] = 0
				res--
				break
			}
		}
	}

	return res
}

/*
思路：线段树
*/

// SegmentTree segment tree
type SegmentTree struct {
	data []int
	tree []int
	n    int
}

// NewSegmentTree 构建线段树
func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	tree := make([]int, 4*n) // 线段树大小一般取 4 倍空间
	st := &SegmentTree{
		data: nums,
		tree: tree,
		n:    n,
	}
	st.build(1, 0, n-1)
	return st
}

// build 构建函数（递归）
func (st *SegmentTree) build(node, l, r int) {
	if l == r {
		st.tree[node] = st.data[l]
		return
	}
	mid := (l + r) / 2
	st.build(node*2, l, mid)
	st.build(node*2+1, mid+1, r)
	st.tree[node] = max(st.tree[node*2], st.tree[node*2+1])
}

// Query 区间和查询
func (st *SegmentTree) Query(max int) int {
	return st.query(1, 0, st.n-1, max)
}

func (st *SegmentTree) query(node, l, r, value int) int {
	if st.tree[node] < value {
		return -1
	}
	if l == r {
		return l
	}

	mid := (l + r) / 2
	left := st.query(node*2, l, mid, value)
	if left != -1 {
		return left
	}
	return st.query(node*2+1, mid+1, r, value)
}

// Update 更新某个位置的值
func (st *SegmentTree) Update(index, val int) {
	st.update(1, 0, st.n-1, index, val)
}

func (st *SegmentTree) update(node, l, r, index, val int) {
	if l == r {
		st.tree[node] = val
		return
	}
	mid := (l + r) / 2
	if index <= mid {
		st.update(node*2, l, mid, index, val)
	} else {
		st.update(node*2+1, mid+1, r, index, val)
	}
	st.tree[node] = st.tree[node*2] + st.tree[node*2+1]
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	stTree := NewSegmentTree(baskets)

	var res int
	for _, fruit := range fruits {
		index := stTree.Query(fruit)
		if index == -1 {
			res++
		} else {
			// update
			stTree.Update(index, 0)
		}
	}

	return res
}

func main() {
	testCases := []struct {
		fruits   []int
		baskets  []int
		expected int
	}{
		{
			fruits:   []int{4, 2, 5},
			baskets:  []int{3, 5, 4},
			expected: 1,
		},
	}
	for i, tc := range testCases {
		res := numOfUnplacedFruits(tc.fruits, tc.baskets)
		if res != tc.expected {
			fmt.Printf("Case #%d fail, expected %d, got %d\n", i, tc.expected, res)
		} else {
			fmt.Printf("Case #%d pass, expected %d, got %d\n", i, tc.expected, res)
		}
	}
}
