package main

import (
	"fmt"
	"sync"
)

/*
思路：
1、通过两个条件变量实现队列满时的写阻塞和队列空时的读阻塞
*/

type BlockQueue struct {
	mutex     *sync.Mutex
	emptyCond *sync.Cond
	fullCond  *sync.Cond
	queue     []int
	writePos  int
	readPos   int
	capacity  int
}

func NewBlockQueue(capacity int) *BlockQueue {
	mutex := &sync.Mutex{}
	return &BlockQueue{
		mutex:     mutex,
		emptyCond: sync.NewCond(mutex),
		fullCond:  sync.NewCond(mutex),
		queue:     make([]int, capacity+1),
		writePos:  0,
		readPos:   0,
		capacity:  capacity + 1,
	}
}

func (bq *BlockQueue) Push(item int) {
	bq.mutex.Lock()
	defer bq.mutex.Unlock()

	for (bq.writePos+1)%bq.capacity == bq.readPos {
		// 队列已满
		bq.fullCond.Wait()
	}

	bq.queue[bq.writePos] = item
	bq.writePos = (bq.writePos + 1) % bq.capacity

	bq.emptyCond.Broadcast()
}

func (bq *BlockQueue) Pop() int {
	bq.mutex.Lock()
	defer bq.mutex.Unlock()

	for bq.readPos == bq.writePos {
		// 队列已空
		bq.emptyCond.Wait()
	}

	result := bq.queue[bq.readPos]
	bq.readPos = (bq.readPos + 1) % bq.capacity
	bq.fullCond.Broadcast()

	return result
}

func (bq *BlockQueue) Empty() bool {
	return bq.readPos == bq.writePos
}

func (bq *BlockQueue) Full() bool {
	return (bq.writePos+1)%bq.capacity == bq.readPos
}

func main() {
	fmt.Println("vim-go")
	bq := NewBlockQueue(10)
	chDone := make(chan interface{})
	wg := &sync.WaitGroup{}
	go func(chDone chan interface{}) {
		done := false
		for !done {
			select {
			case <-chDone:
				done = true
			default:
				result := bq.Pop()
				fmt.Printf("pop:%d\n", result)
			}
		}
	}(chDone)

	for gid := 0; gid < 10; gid++ {
		wg.Add(1)
		go func(gid int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				bq.Push(i + gid*10)
				fmt.Printf("push:%d\n", i+gid*10)
				// time.Sleep(1 * time.Microsecond)
			}
		}(gid)
	}
	wg.Wait()
	chDone <- struct{}{}
}
