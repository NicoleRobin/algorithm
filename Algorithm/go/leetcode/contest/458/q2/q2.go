package main

import "sort"

/*
思路：并查集
*/
func minCost(n int, edges [][]int, k int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] > edges[j][2]
	})

	var ans int

	return ans
}

func main() {

}
