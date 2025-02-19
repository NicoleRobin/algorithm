package linklist

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

func NewList(arr []int) *ListNode {
	head := &ListNode{}
	p := head
	for i := 0; i < len(arr); i++ {
		p.Next = &ListNode{Val: arr[i]}
		p = p.Next
	}
	return head.Next
}

func (l *ListNode) String() string {
	var nodeList []string
	for p := l; p != nil; p = p.Next {
		nodeList = append(nodeList, fmt.Sprintf("%d", p.Val))
	}

	return strings.Join(nodeList, "->")
}
