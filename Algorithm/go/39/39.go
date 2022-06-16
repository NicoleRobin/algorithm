package main

import (
	"fmt"
	"sort"
)

/*
注意：
1、需要判断path是否重复，例如数字个数相同，顺序不一样是被认为相同的，不能重复添加

目前只有一个办法是在结束时去判断
并且，由于存储结果集是使用slice来存的，要去重就需要和已有的所有结果集进行匹配，成本有点高
*/
func combinationSum(candidates []int, target int) [][]int {
	mapResult := map[string][]int{}
	path := []int{}
	backtrack(candidates, path, target, mapResult)

	sliceResult := [][]int{}
	for _, value := range mapResult {
		sliceResult = append(sliceResult, value)
	}
	return sliceResult
}

func strPath(path []int) string {
	sort.Ints(path)
	strPath := ""
	for _, num := range path {
		strPath += fmt.Sprintf("%d_", num)
	}
	return strPath
}

func backtrack(candidates, path []int, target int, result map[string][]int) {
	fmt.Printf("candidates:%+v, path:%+v, target:%d\n", candidates, path, target)
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		strPath := strPath(tmp)
		fmt.Println(strPath)
		if _, ok := result[strPath]; !ok {

			result[strPath] = tmp
		}
		return
	} else if target < 0 {
		return
	}

	for _, candidate := range candidates {
		fmt.Printf("candidate:%d, path:%+v\n", candidate, path)
		path = append(path, candidate)
		fmt.Printf("candidate:%d, path:%+v\n", candidate, path)
		backtrack(candidates, path, target-candidate, result)
		fmt.Printf("candidate:%d, path:%+v\n", candidate, path)
		path = path[:len(path)-1]
		fmt.Printf("candidate:%d, path:%+v\n", candidate, path)
	}
}

func main() {
	fmt.Println("leetcode-39")
	candidates := []int{2, 7, 6, 3, 5, 1}
	target := 9
	result := combinationSum(candidates, target)
	fmt.Println(result)
}
