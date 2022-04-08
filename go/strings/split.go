package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(8)

	fmt.Println("strings example")
	str := "a string  example"
	words := strings.Split(str, " ")
	fmt.Printf("%+v\n", words)

	char := '9'
	fmt.Println(reflect.TypeOf(char))
	fmt.Printf("%d\n", char)

	str = "zjw张"
	for _, char := range str {
		fmt.Printf("type:%s, code:%d\n", reflect.TypeOf(char), char)
	}

	for i := 0; i < len(str); i++ {
		fmt.Printf("type:%s, code:%d\n", reflect.TypeOf(str[i]), str[i])
	}
}
