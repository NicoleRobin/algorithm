package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("strings example")
	str := "a string  example"
	words := strings.Split(str, " ")
	fmt.Printf("%+v\n", words)
}
