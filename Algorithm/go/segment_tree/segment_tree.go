package segment_tree

type SegmentTree struct {
	data []int
	tree []int
	n    int
}

// NewSegmentTree 构建线段树
func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	tree := make([]int, 4*n) // 线段树大小一般取 4 倍空间
	st := &SegmentTree{nums, tree, n}
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
	st.tree[node] = st.tree[node*2] + st.tree[node*2+1]
}

// Query 区间和查询
func (st *SegmentTree) Query(qL, qR int) int {
	return st.query(1, 0, st.n-1, qL, qR)
}

func (st *SegmentTree) query(node, l, r, qL, qR int) int {
	if qL > r || qR < l {
		return 0 // 不相交
	}
	if qL <= l && r <= qR {
		return st.tree[node] // 完全包含
	}
	mid := (l + r) / 2
	left := st.query(node*2, l, mid, qL, qR)
	right := st.query(node*2+1, mid+1, r, qL, qR)
	return left + right
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
