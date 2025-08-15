package main

import (
	"fmt"
	"strconv"
)

func maximum69Number(num int) int {
	strNUm := fmt.Sprintf("%d", num)
	byteNum := []byte(strNUm)
	for i := 0; i < len(byteNum); i++ {
		if byteNum[i] == '6' {
			byteNum[i] = '9'
			break
		}
	}

	iNum, _ := strconv.ParseInt(string(byteNum), 10, 64)
	return int(iNum)
}

func main() {

}
