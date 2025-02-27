package main

import (
	"fmt"
	"github.com/nicolerobin/algorithm/go/linklist"
)

func merge(head1, head2 *linklist.ListNode) *linklist.ListNode {
	newHead := &linklist.ListNode{}

	return newHead.Next
}

func main() {
	head1 := linklist.NewList([]int{1, 2, 4})
	head2 := linklist.NewList([]int{1, 3, 4})

	newHead := merge(head1, head2)
	fmt.Println("head1: ", head1)
	fmt.Println("head2: ", head2)
	fmt.Println("newHead: ", newHead)
}
