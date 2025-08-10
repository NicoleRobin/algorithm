package main

import (
	"fmt"
	"sort"
)

/*
回溯法
*/
func permutation(s string) []string {
	sBytes := []byte(s)
	sort.Slice(sBytes, func(i, j int) bool {
		return sBytes[i] < sBytes[j]
	})
	result := []string{}
	existS := map[string]interface{}{}
	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(s) {
			if _, ok := existS[string(sBytes)]; ok {
				fmt.Printf("duplicate str:%s\n", sBytes)
			} else {
				existS[string(sBytes)] = struct{}{}
			}
			result = append(result, string(sBytes))
			return
		}

		existChar := map[string]interface{}{}
		for i := start; i < len(s); i++ {
			if _, ok := existChar[string(sBytes[i])]; ok {
				continue
			} else {
				existChar[string(sBytes[i])] = struct{}{}
			}
			fmt.Printf("start:%d, i:%d, sBytes:%c, existChar:%+v\n", start, i, sBytes[i], existChar)
			sBytes[start], sBytes[i] = sBytes[i], sBytes[start]
			fmt.Printf("-> already:%s, remain:%s\n", sBytes[:start+1], sBytes[start+1:])
			backtrack(start + 1)
			sBytes[i], sBytes[start] = sBytes[start], sBytes[i]
			fmt.Printf("<- already:%s, remain:%s\n", sBytes[:start+1], sBytes[start+1:])
		}
	}
	backtrack(0)
	return result
}

func main() {
	fmt.Println("leetcode-38")
	// s := "suvyls"
	s := "suvs"
	result := permutation(s)
	fmt.Println(result)
	/*
		existS := map[string]interface{}{}
		for _, item := range result {
			if _, ok := existS[item]; ok {
				fmt.Printf("duplicate str:%s\n", item)
			} else {
				existS[item] = struct{}{}
			}
		}
	*/
}
