package main

type TrieNode struct {
	IsEnd    bool
	Children []*TrieNode
}

func NewTriNode() *TrieNode {
	return &TrieNode{
		IsEnd:    false,
		Children: make([]*TrieNode, 26, 26),
	}
}

type Trie struct {
	Root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Root: NewTriNode()}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this.Root
	for _, r := range word {
		idx := r - 'a'
		if root.Children[idx] == nil {
			root.Children[idx] = NewTriNode()
		}
		root = root.Children[idx]
	}
	root.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.searchEndNode(word)
	return node != nil && node.IsEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.searchEndNode(prefix) != nil
}

func (this *Trie) searchEndNode(word string) *TrieNode {
	root := this.Root
	for _, r := range word {
		if root == nil {
			break
		}
		root = root.Children[r-'a']
	}
	return root
}
