package main

/*
思考：
1、一定是一个合法的字符串么？

思路：
1、递归
*/
func isOneBitCharacter(bits []int) bool {
	if len(bits) == 0 {
		return false
	}
	if len(bits) == 1 {
		return true
	}
	if bits[len(bits)-2] == 0 {
		return true
	}

	var oneCount int
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 1 {
			oneCount++
		}
		if bits[i] == 0 {
			break
		}
	}
	if oneCount&1 == 0 {
		return true
	}
	return false
}

func main() {

}
