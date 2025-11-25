package main

/*
余数的等价交换
*/
func smallestRepunitDivByK(k int) int {
	if k%10 == 2 || k%10 == 5 {
		return -1
	}

	ans := 1
	resid := 1 % k
	residSet := map[int]bool{
		resid: true,
	}
	for resid != 0 {
		resid = (resid*10 + 1) % k
		ans++
		if _, ok := residSet[resid]; ok {
			return -1
		}
		residSet[resid] = true
	}
	return ans
}

func main() {

}
