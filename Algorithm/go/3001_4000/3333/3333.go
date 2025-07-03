package main

/*
思考：
1、该题是3330的升级版，添加了至少输入字符长度的限制；
2、是不是仍然可以按照3330的思路来做，只是减去长度小于指定最小长度k的方案即可；

思路：
1、
*/
func possibleStringCount(word string, k int) int {
	n := len(word)
	if n < k {
		return 0
	}

	var dupCount int
	for i := 1; i < n; i++ {
		if word[i] == word[i-1] {
			dupCount++
		}
	}

	var ans int

	return ans % (1e9 + 7)
}
func main() {

}
