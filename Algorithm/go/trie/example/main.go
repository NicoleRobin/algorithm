package main

import (
	"github.com/nicolerobin/algorithm/go/trie"
)

func main() {
	trieHead := trie.NewTrie()
	trieHead.Insert("apple")
	trieHead.Insert("app")
	trie.Print(trieHead, "")
}
