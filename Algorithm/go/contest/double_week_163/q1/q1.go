package main

func minSensors(n int, m int, k int) int {
	sideLen := 2*k + 1
	mCount := (m + sideLen - 1) / sideLen
	nCount := (n + sideLen - 1) / sideLen

	return mCount * nCount
}

func main() {

}
