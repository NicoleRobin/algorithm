package main

import (
	"fmt"
)

/*
该题是字节的面试题：
微服务架构下存在多种调用关系，但是当存在循环调用时是比较危险的，因此需要对这些调用进行检测
例如：
[('A', 'B'),('B', 'D'),('A', 'C'),('D', 'A')]
存在循环调用，因此返回true
*/

type CallRelation struct {
	Caller string
	Callee string
}

func callRelation(callRelations []CallRelation) bool {
	callerMap := map[string][]string{}
	for _, callRelation := range callRelations {
		if callees, ok := callerMap[callRelation.Caller]; ok {
			callees = append(callees, callRelation.Callee)
			callerMap[callRelation.Caller] = callees
		} else {
			callerMap[callRelation.Caller] = []string{callRelation.Callee}
		}
	}
	fmt.Println(callerMap)

	var checkRelation func(caller, root string) bool
	checkRelation = func(caller, root string) bool {
		fmt.Printf("caller:%s, root:%s\n", caller, root)
		if callees, ok := callerMap[caller]; !ok {
			return true
		} else {
			for _, callee := range callees {
				if callee == root {
					return false
				}

				if !checkRelation(callee, root) {
					return false
				}
			}
		}

		return true
	}

	for key, _ := range callerMap {
		if !checkRelation(key, key) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("字节面试")
	callRelations := []CallRelation{
		{
			Caller: "A",
			Callee: "B",
		},
		{
			Caller: "A",
			Callee: "C",
		},
		{
			Caller: "B",
			Callee: "D",
		},
		{
			Caller: "D",
			Callee: "E",
		},
	}

	hasCycle := callRelation(callRelations)
	fmt.Println(hasCycle)
}
