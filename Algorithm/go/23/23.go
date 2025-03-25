package main

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	newHead := &ListNode{}
	newP := newHead

	var emptyListCount int
	for emptyListCount < len(lists) {
		emptyListCount = 0

		var min int = math.MaxInt32
		var minI int
		var p *ListNode
		for i, listNode := range lists {
			if listNode == nil {
				emptyListCount++
				continue
			}
			if listNode != nil && listNode.Val < min {
				p = listNode
				min = listNode.Val
				minI = i
			}
		}

		if p != nil {
			newP.Next = p
			lists[minI] = p.Next
			p.Next = nil
			newP.Next = p
			newP = newP.Next
		}
	}

	return newHead.Next
}

func main() {

}
