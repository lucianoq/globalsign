package main

type Adder interface {
	Add([]byte, int)
}

func NewTrie() *trie {
	return &trie{
		&node{ptr: make([]*node, 0)},
	}
}

type trie struct {
	root *node
}

func (t *trie) Add(word []byte) {
	t.root.add(word)
}

func (t *trie) AddAll(add Adder) {
	t.root.addAll(make([]byte, 0), add)
}

type node struct {
	char  byte
	count int
	ptr   []*node
}

func (n *node) add(word []byte) {
	if len(word) == 0 {
		n.count++
		return
	}

	curr := word[0]
	suffix := word[1:]

	for _, next := range n.ptr {
		if next.char == curr {
			next.add(suffix)
			return
		}
	}

	newNode := &node{
		char: curr,
		ptr:  make([]*node, 0),
	}
	n.ptr = append(n.ptr, newNode)
	newNode.add(suffix)
}

// addAll visits the trie depth-first and `Add`s all the words
// with at least 1 occurrence inside the Adder in input
func (n *node) addAll(prefix []byte, add Adder) {
	if n.count != 0 {
		var word = make([]byte, len(prefix))
		copy(word, prefix)
		add.Add(word, n.count)
	}

	for _, next := range n.ptr {
		next.addAll(append(prefix, next.char), add)
	}
}
