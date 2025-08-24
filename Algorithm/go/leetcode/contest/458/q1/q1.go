package main

func processStr(s string) string {
	var ans []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		} else if s[i] == '#' {
			ans = append(ans, ans[:]...)
		} else if s[i] == '%' {
			reverseBytes(ans)
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

func reverseBytes(ans []byte) {
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
}

func main() {

}
