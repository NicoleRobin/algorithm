package main

import (
	"fmt"
)

func main() {
	fmt.Println("channel")
	c := make(chan string)
	c <- "channel"

	fmt.Println("hh")
}
