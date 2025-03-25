package main

import (
	"fmt"
	"github.com/nicolerobin/algorithm/go/linklist"
)

func reverseKGroup(head *linklist.ListNode, k int) *linklist.ListNode {
	newHead := &linklist.ListNode{}
	newLastTail := newHead

	var count int
	lastHead := head
	p := head
	for p != nil {
		count++
		pNext := p.Next

		if count == k {
			count = 0

			// 翻转lastHead到p
			newSubHead := &linklist.ListNode{}
			subP := lastHead
			for subP != nil {
				subPNext := subP.Next
				subP.Next = newSubHead.Next
				newSubHead.Next = subP

				if subP == p {
					break
				}

				subP = subPNext
			}
			fmt.Println(newSubHead)
			newLastTail.Next = newSubHead.Next
			// 链上原来的剩余部分
			lastHead.Next = pNext

			newLastTail = lastHead

			lastHead = pNext
		}

		p = pNext
	}

	return newHead.Next
}

func main() {
	head := linklist.NewList([]int{1, 2, 3, 4, 5})
	fmt.Println(head)
	newHead := reverseKGroup(head, 2)

	fmt.Println(newHead)
}
