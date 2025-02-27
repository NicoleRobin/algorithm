package union_find

// UnionFind 并查集结构体
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind 初始化并查集
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)

	// 初始化，每个元素的父节点指向自己，秩为0
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &UnionFind{parent: parent, rank: rank}
}

// Find 查找操作，带路径压缩
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// Union 合并操作，按秩合并
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// 按秩合并
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++ // 如果秩相同，合并后增加根的秩
		}
	}
}

// Connected 判断两个元素是否属于同一个集合
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
