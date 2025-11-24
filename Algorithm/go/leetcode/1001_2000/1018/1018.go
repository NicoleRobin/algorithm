package main

/*
注意：
1、需要考虑before会溢出的情况，由于是否被5整除只会跟后面的几位相关，可以通过求余舍弃高位避免溢出；
*/
func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	var before int
	for i := 0; i < len(nums); i++ {
		before = before*2 + nums[i]
		if before%5 == 0 {
			ans[i] = true
		}

		before = before % 1000
	}

	return ans
}

func main() {

}
