package main

/*
思考：
1、该题是考察模运算的
2、我们可以使用前缀和来解决这个问题
*/
func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	var preSum int
	var ans int64
	cnt := map[int]int{}
	cnt[0] = 1
	for _, num := range nums {
		if num%modulo == k {
			preSum++
		}
		ans += int64(cnt[(preSum-k+modulo)%modulo])
		cnt[preSum%modulo]++
	}
	return int64(ans)
}

func main() {

}
