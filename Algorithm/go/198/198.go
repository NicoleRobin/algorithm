package main

/*
思路：二维动态规划
抢了当前节点的结果：robDp[i] = max(robDp[i-2] + nums[i], notRobDp[i-1]+nums[i])
不抢当前节点的结果：notRobDp[i] = max(robDp[i-1], notRobDp[i-1])
*/
func rob(nums []int) int {

}

func main() {
}
