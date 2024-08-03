package trees

type Trie[T any] interface {
	// Insert inserts the item to Trie with custom comparator
	Insert(item T)

	// Search finds if the element exists in the Trie.
	// Return true is exists, otherwise returns false.
	Search(item T) bool

	// StartsWith returns true, if there are elements in the Trie starts with prefix
	StartsWith(prefix T) bool

	// CountByPrefix returns int number of elements which have prefix arg
	CountByPrefix(prefix T) int
}

type trieNode[T any] struct {
	val T

	children map[rune]*trieNode[T]
	isEnd    bool
}

type trie[T comparable] struct {
	root     *trieNode[T]
	compare  Comparator[T]
	toString func(T) string
}

func NewTrie[T comparable](cmp Comparator[T], toString func(T) string) *trie[T] {
	return &trie[T]{
		compare:  cmp,
		root:     &trieNode[T]{children: make(map[rune]*trieNode[T])},
		toString: toString,
	}
}

// Insert inserts the item to Trie with custom comparator
func (t *trie[T]) Insert(item T) {
	itemStr := t.toString(item)
	dummy := t.root
	for _, ch := range itemStr {
		if _, exists := dummy.children[ch]; !exists {
			dummy.children[ch] = &trieNode[T]{children: make(map[rune]*trieNode[T])}
		}
		dummy = dummy.children[ch]
	}
	dummy.isEnd = true
	dummy.val = item
}

// Search finds if the element exists in the Trie.
// Return true is exists, otherwise returns false.
func (t *trie[T]) Search(item T) bool {
	itemStr := t.toString(item)
	dummy := t.root
	for _, ch := range itemStr {
		if _, exists := dummy.children[ch]; !exists {
			return false
		}
		dummy = dummy.children[ch]
	}

	return dummy.isEnd
}

// StartsWith returns true, if there are elements in the Trie starts with prefix
func (t *trie[T]) StartsWith(prefix T) bool {
	itemStr := t.toString(prefix)
	dummy := t.root
	for _, ch := range itemStr {
		if _, exists := dummy.children[ch]; !exists {
			return false
		}
		dummy = dummy.children[ch]
	}

	return true
}

// CountByPrefix returns int number of elements which have prefix arg
func (t *trie[T]) CountByPrefix(prefix T) int {

	itemStr := t.toString(prefix)
	dummy := t.root
	var counter int
	for _, ch := range itemStr {
		if _, exists := dummy.children[ch]; !exists {
			return counter
		}
		dummy = dummy.children[ch]
	}

	return countEndNodes(dummy)
}

func countEndNodes[T any](node *trieNode[T]) int {
	count := 0
	if node.isEnd {
		count++
	}
	for _, child := range node.children {
		count += countEndNodes(child)
	}
	return count
}
