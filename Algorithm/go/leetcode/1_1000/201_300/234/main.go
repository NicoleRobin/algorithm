package main

import (
	"fmt"
	"github.com/nicolerobin/algorithm/go/linklist"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *linklist.ListNode) bool {
	var count int
	for p := head; p != nil; p = p.Next {
		count++
	}

	tail := head
	lastHalfHead := head
	for i := 0; i < count/2; i++ {
		tail = lastHalfHead
		lastHalfHead = lastHalfHead.Next
	}
	fmt.Println("last half head", lastHalfHead)

	lastHalfNewHead := reverseLinkList(lastHalfHead)
	fmt.Println("last reverse head", lastHalfNewHead)

	p := head
	newP := lastHalfNewHead
	isReverseList := true
	for i := 0; i < count/2; i++ {
		if p.Val != newP.Val {
			isReverseList = false
			break
		}
		p = p.Next
		newP = newP.Next
	}
	lastHalfHead = reverseLinkList(lastHalfNewHead)
	tail.Next = lastHalfHead

	return isReverseList
}

func reverseLinkList(head *linklist.ListNode) *linklist.ListNode {
	newHead := &linklist.ListNode{}
	for p := head; p != nil; {
		newHeadNext := newHead.Next
		pNext := p.Next

		newHead.Next = p
		p.Next = newHeadNext

		p = pNext
	}
	return newHead.Next
}

func main() {
	nums := []int{1, 1, 2, 1}
	head := linklist.NewList(nums)
	fmt.Println("init", head)
	ok := isPalindrome(head)
	fmt.Println(ok)
	fmt.Println(head)
}
