package main

type NumArray struct {
	nums []int
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, 4*n) // 线段树大小一般取 4 倍空间
	st := &NumArray{
		nums: nums,
		tree: tree,
		n:    n,
	}
	st.build(1, 0, n-1)
	return st
}

func (this *NumArray) build(node, l, r int) int {
	if l == r {
		st.tree[node] = st.nums[l]
		return
	}
	mid := (l + r) / 2
	left := st.build(node*2, l, mid)
	right := st.build(node*2+1, mid+1, r)
	st.tree[node] = left + right
	return st.tree[node]
}

func (this *NumArray) Update(index int, val int) {

}

func (this *NumArray) update(index int, val int) {

}

func (this *NumArray) SumRange(left int, right int) int {

}

func (this *NumArray) query(node, l, r int) int {
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

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func main() {}
