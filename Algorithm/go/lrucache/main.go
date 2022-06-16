package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type LRUCache struct {
	kv         map[int]*ListNode
	head, tail *ListNode
	capacity   int
}

type ListNode struct {
	Pre, Next *ListNode
	Key       int
	Val       int
}

func New(capacity int) *LRUCache {
	head := &ListNode{}
	tail := &ListNode{}
	head.Next = tail
	tail.Pre = head
	return &LRUCache{
		kv:       make(map[int]*ListNode),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (lc *LRUCache) Print() {
	pNode := lc.head.Next
	for pNode.Next != nil {
		fmt.Printf("%d  ", pNode.Val)
		pNode = pNode.Next
	}
	fmt.Println()
}

func (lc *LRUCache) refreshNode(node *ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	node.Next = lc.head.Next
	lc.head.Next.Pre = node
	node.Pre = lc.head
	lc.head.Next = node
}

func (lc *LRUCache) Get(key int) (int, error) {
	if pNode, ok := lc.kv[key]; ok {
		lc.refreshNode(pNode)
		return pNode.Val, nil
	}
	return 0, errors.New("key not found")
}

func (lc *LRUCache) Put(key int, value int) error {
	if pNode, ok := lc.kv[key]; ok {
		pNode.Val = value
		lc.refreshNode(pNode)
	} else {
		if len(lc.kv) == lc.capacity {
			lastNode := lc.tail.Pre
			lastNode.Pre.Next = lastNode.Next
			lastNode.Next.Pre = lastNode.Pre
			delete(lc.kv, lastNode.Key)
		}
		pNode := &ListNode{
			Key: key,
			Val: value,
		}
		pNode.Next = lc.head.Next
		lc.head.Next.Pre = pNode
		lc.head.Next = pNode
		pNode.Pre = lc.head
		lc.kv[key] = pNode
	}
	return nil
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	fmt.Println("vim-go")
	lc := New(10)
	for i := 0; i < 10; i++ {
		lc.Put(i, i)
	}
	lc.Print()

	count := 0
	for count < 10 {
		key := rand.Intn(10)
		val, err := lc.Get(key)
		if err != nil {
			fmt.Printf("get key:%d, err:%s\n", key, err)
		} else {
			fmt.Printf("get key:%d, val:%d\n", key, val)
		}
		lc.Print()
		count++
	}

	count = 0
	for count < 10 {
		key := 100 + rand.Intn(10)
		err := lc.Put(key, key)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("put key:%d\n", key)
		}
		lc.Print()
		count++
	}
}
