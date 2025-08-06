package segment_tree

import (
	"reflect"
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *SegmentTree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSegmentTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSegmentTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegmentTree_Query(t *testing.T) {
	type fields struct {
		data []int
		tree []int
		n    int
	}
	type args struct {
		qL int
		qR int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				data: tt.fields.data,
				tree: tt.fields.tree,
				n:    tt.fields.n,
			}
			if got := st.Query(tt.args.qL, tt.args.qR); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegmentTree_Update(t *testing.T) {
	type fields struct {
		data []int
		tree []int
		n    int
	}
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				data: tt.fields.data,
				tree: tt.fields.tree,
				n:    tt.fields.n,
			}
			st.Update(tt.args.index, tt.args.val)
		})
	}
}

func TestSegmentTree_build(t *testing.T) {
	type fields struct {
		data []int
		tree []int
		n    int
	}
	type args struct {
		node int
		l    int
		r    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				data: tt.fields.data,
				tree: tt.fields.tree,
				n:    tt.fields.n,
			}
			st.build(tt.args.node, tt.args.l, tt.args.r)
		})
	}
}

func TestSegmentTree_query(t *testing.T) {
	type fields struct {
		data []int
		tree []int
		n    int
	}
	type args struct {
		node int
		l    int
		r    int
		qL   int
		qR   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				data: tt.fields.data,
				tree: tt.fields.tree,
				n:    tt.fields.n,
			}
			if got := st.query(tt.args.node, tt.args.l, tt.args.r, tt.args.qL, tt.args.qR); got != tt.want {
				t.Errorf("query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegmentTree_update(t *testing.T) {
	type fields struct {
		data []int
		tree []int
		n    int
	}
	type args struct {
		node  int
		l     int
		r     int
		index int
		val   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				data: tt.fields.data,
				tree: tt.fields.tree,
				n:    tt.fields.n,
			}
			st.update(tt.args.node, tt.args.l, tt.args.r, tt.args.index, tt.args.val)
		})
	}
}
