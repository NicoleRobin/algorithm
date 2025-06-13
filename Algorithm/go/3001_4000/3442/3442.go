package main

import "math"

func maxDifference(s string) int {
	charCountMap := map[rune]int{}
	for _, char := range s {
		charCountMap[char]++
	}

	var maxJi, maxOu int
	minJi, minOu := math.MaxInt32, math.MaxInt32
	for _, count := range charCountMap {
		if count%2 == 0 {
			if count < minOu {
				minOu = count
			}
			if count > maxOu {
				maxOu = count
			}
		} else {
			if count < minJi {
				minJi = count
			}
			if count > maxJi {
				maxJi = count
			}
		}
	}
	diff1 := maxJi - minOu
	diff2 := maxJi - maxOu
	diff3 := minJi - minOu
	diff4 := minJi - maxOu

	return max(diff1, max(diff2, max(diff3, diff4)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
