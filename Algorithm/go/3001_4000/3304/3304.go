package main

import "math/bits"

func kthCharacter(k int) byte {
	var ans int
	for k != 1 {
		t := bits.Len(uint(k)) - 1
		if 1<<t == k {
			t--
		}
		k -= 1 << t
		ans++
	}
	return byte('a' + ans)
}

func main() {

}
