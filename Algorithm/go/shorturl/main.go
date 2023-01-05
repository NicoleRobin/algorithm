package main

import (
	"github.com/nicolerobin/algorithm/go/shorturl/shorturl"
	"github.com/nicolerobin/log"
)

func main() {
	for i := 0; i < 100000; i += 10000 {
		shortUrl := shorturl.ShortUrl(i)
		log.Debug("id:%d, shortUrl:%s", i, shortUrl)
	}
}
