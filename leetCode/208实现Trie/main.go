package main

type Trie struct {
	child [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	n := t
	for _, ch := range word {
		ch -= 'a'
		if n.child[ch] == nil {
			n.child[ch] = new(Trie)
		}
		n = n.child[ch]
	}
	n.isEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	n := t.searchPrefix(word)
	return n != nil && n.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	n := t
	for _, ch := range prefix {
		ch -= 'a'
		if n.child[ch] == nil {
			return nil
		}
		n = n.child[ch]
	}
	return n
}
