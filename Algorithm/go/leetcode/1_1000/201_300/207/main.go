package main

import "fmt"

/*
拓扑排序
深度优先搜索
对于每一个节点有三种状态：初始状态、遍历中、已遍历
这个思路是按照节点的顺序关系一直遍历到底，所以先遍历到的节点它在拓扑排序中的顺序越靠后
*/
func canFinish1(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	// 遍历所有节点的顺序数组
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

/*
拓扑排序
广度优先搜索
这个比较好理解，一层层的消除入度为0的节点
最后判断得到的节点个数和总个数是否相等就可以知道是否有环
*/
func canFinish2(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result []int
	)

	// 首先创建节点的顺序关系
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		// 计算每个节点的入度
		indeg[info[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}

type TestCase struct {
	numCourses    int
	prerequisites [][]int
}

func main() {
	fmt.Println("vim-go")
	testCases := []TestCase{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{0, 1}, {1, 0}},
		},
		{
			numCourses:    6,
			prerequisites: [][]int{{3, 0}, {3, 1}, {4, 1}, {4, 2}, {5, 3}, {5, 4}},
		},
	}
	for _, testCase := range testCases {
		result := canFinish1(testCase.numCourses, testCase.prerequisites)
		fmt.Printf("prerequistes:%+v, result:%t\n", testCase.prerequisites, result)
	}
}
