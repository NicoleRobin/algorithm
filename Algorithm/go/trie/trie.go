package trie

import "fmt"

// Trie trie node
type Trie struct {
	Children map[rune]*Trie
	IsEnd    bool
}

func NewTrie() *Trie {
	return &Trie{
		Children: map[rune]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	currNode := t
	for _, c := range word {
		if currNode.Children[c] == nil {
			currNode.Children[c] = &Trie{
				Children: map[rune]*Trie{},
			}
		}
		currNode = currNode.Children[c]
	}
	currNode.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	currNode := t
	for _, c := range word {
		if currNode.Children[c] == nil {
			return false
		}
		currNode = currNode.Children[c]
	}
	return currNode.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	currNode := t
	for _, c := range prefix {
		if currNode.Children[c] == nil {
			return false
		}
		currNode = currNode.Children[c]
	}
	return true
}

func Print(t *Trie, prefix string) {
	if t == nil {
		return
	}

	if t.IsEnd {
		fmt.Println(prefix)
	}

	for k, v := range t.Children {
		Print(v, prefix+string(k))
	}
}
