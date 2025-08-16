package main

type NumArray struct {
	nums []int
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, 4*n) // 线段树大小一般取 4 倍空间
	st := NumArray{
		nums: nums,
		tree: tree,
		n:    n,
	}
	st.build(1, 0, n-1)
	return st
}

func (this *NumArray) build(node, l, r int) {
	if l == r {
		this.tree[node] = this.nums[l]
	}
	mid := (l + r) / 2
	this.build(node*2, l, mid)
	this.build(node*2+1, mid+1, r)
	this.tree[node] = this.tree[node*2] + this.tree[node*2+1]
}

func (this *NumArray) Update(index int, val int) {
	this.update(1, 0, this.n-1, index, val)
}

func (this *NumArray) update(node, l, r, index, val int) {
	if l == r {
		this.tree[node] = val
		return
	}
	mid := (l + r) / 2
	if index <= mid {
		this.update(node*2, l, mid, index, val)
	} else {
		this.update(node*2+1, mid+1, r, index, val)
	}
	this.tree[node] = this.tree[node*2] + this.tree[node*2+1]
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.query(1, 0, this.n-1, left, right)
}

func (this *NumArray) query(node, l, r, qL, qR int) int {
	if qL > r || qR < l {
		return 0 // 不相交
	}
	if qL <= l && r <= qR {
		return this.tree[node] // 完全包含
	}
	mid := (l + r) / 2
	left := this.query(node*2, l, mid, qL, qR)
	right := this.query(node*2+1, mid+1, r, qL, qR)
	return left + right
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func main() {}
