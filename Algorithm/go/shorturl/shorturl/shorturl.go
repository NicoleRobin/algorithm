package shorturl

import (
	"bytes"
	"fmt"
)

const (
	BASE = 62
)

var digitCharMap = map[int]byte{}

func init() {
	for i := 0; i < 62; i++ {
		if i < 10 {
			digitCharMap[i] = byte('0' + i)
		} else if i >= 10 && i < 36 {
			digitCharMap[i] = byte('a' + i - 10)
		} else if i >= 36 {
			digitCharMap[i] = byte('A' + i - 36)
		}
	}
	fmt.Println(digitCharMap)
}

func ShortUrl(id int) string {
	out := bytes.Buffer{}

	for id > 0 {
		digit := id % BASE
		out.WriteByte(digitCharMap[digit])
		id /= BASE
	}
	for out.Len() < 6 {
		out.WriteByte(byte('0'))
	}
	return out.String()
}
