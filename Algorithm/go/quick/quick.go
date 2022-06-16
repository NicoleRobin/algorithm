package quick

/*
快速乘法的原理是将其中一个数拆分为二进制，根据某一位是否是1决定是否需要加上乘数
*/
func QuickMul(x, y int) int {
	ans := 0
	for ; y > 0; y >>= 1 {
		if y&1 > 0 {
			ans += x
		}
		x <<= 1
	}

	return ans
}

/*
x的10次方，转换为
x的1010次方，转换为
x的1*2^3 + 1*2^1次方，转换为
x的2^3次方乘以x的2^1次方
*/
func QuickPow(x, y int) int {
	ans := 1
	for ; y > 0; y >>= 1 {
		if y&1 > 0 {
			ans *= x
		}
		x *= x
	}
	return ans
}
