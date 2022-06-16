package main

import (
	"fmt"
	"sync"
)

func print1_100() {
	const MAX = 10
	wg := sync.WaitGroup{}
	oddChan := make(chan int)
	evenChan := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range evenChan {
			if num > MAX {
				close(oddChan)
				break
			}
			oddChan <- num + 1
			fmt.Printf("goroutine1, num:%d\n", num)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range oddChan {
			if num > MAX {
				close(evenChan)
				break
			}
			evenChan <- num + 1
			fmt.Printf("goroutine2, num:%d\n", num)
		}
	}()
	evenChan <- 0
	wg.Wait()
}

func main() {
	fmt.Println("print 1 to 100")
	print1_100()
}
