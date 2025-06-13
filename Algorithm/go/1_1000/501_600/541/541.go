package main

/*
字符串操作类的题
这里需要使用的技巧是for循环步进由常用的i++改为i+=2*k
*/
func reverseStr(s string, k int) string {
	bytesS := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(s) {
			reverseBytes(bytesS[i : i+k])
			continue
		}
		reverseBytes(bytesS[i:])
	}

	return string(bytesS)
}

func reverseBytes(str []byte) {
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
}

func main() {
	reverseStr("abcdefg", 2)
}
