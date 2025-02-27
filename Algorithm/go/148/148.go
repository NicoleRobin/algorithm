package main

import (
	"fmt"
	"github.com/nicolerobin/algorithm/go/linklist"
)

func merge(head1, head2 *linklist.ListNode) *linklist.ListNode {
	newHead := &linklist.ListNode{}
	for p1, p2, p := head1, head2, newHead; p1 != nil || p2 != nil; {
		if p1 == nil {
			p.Next = p2
			break
		}

		if p2 == nil {
			p.Next = p1
			break
		}

		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}

		p = p.Next
	}

	return newHead.Next
}

func sort(head, tail *linklist.ListNode) *linklist.ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *linklist.ListNode) *linklist.ListNode {
	return sort(head, nil)
}

func main() {
	head1 := linklist.NewList([]int{4, 2, 1, 3})
	fmt.Println("origin list:", head1)
	newHead := sortList(head1)
	fmt.Println("sorted list:", newHead)
}
