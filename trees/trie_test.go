package trees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringComparator(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func stringToString(s string) string {
	return s
}

func intComparator(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func intToString(i int) string {
	return fmt.Sprintf("%d", i)
}

func TestTrie_InsertAndSearch(t *testing.T) {
	trieStrings := NewTrie[string](stringComparator, stringToString)

	trieStrings.Insert("apple")
	trieStrings.Insert("app")
	trieStrings.Insert("apricot")
	trieStrings.Insert("banana")
	trieStrings.Insert("band")
	trieStrings.Insert("bandana")

	assert.True(t, trieStrings.Search("apple"))
	assert.True(t, trieStrings.Search("app"))
	assert.True(t, trieStrings.Search("apricot"))
	assert.True(t, trieStrings.Search("banana"))
	assert.True(t, trieStrings.Search("band"))
	assert.True(t, trieStrings.Search("bandana"))
	assert.False(t, trieStrings.Search("cat"))
	assert.False(t, trieStrings.Search("apples"))
}

func TestTrie_StartsWith(t *testing.T) {
	trieStrings := NewTrie(stringComparator, stringToString)

	trieStrings.Insert("apple")
	trieStrings.Insert("app")
	trieStrings.Insert("apricot")
	trieStrings.Insert("banana")
	trieStrings.Insert("band")
	trieStrings.Insert("bandana")

	assert.True(t, trieStrings.StartsWith("app"))
	assert.True(t, trieStrings.StartsWith("ban"))
	assert.False(t, trieStrings.StartsWith("cat"))
	assert.False(t, trieStrings.StartsWith("bat"))
}

func TestTrie_CountByPrefix(t *testing.T) {
	trieStrings := NewTrie(stringComparator, stringToString)

	trieStrings.Insert("apple")
	trieStrings.Insert("appled")
	trieStrings.Insert("appleds")
	trieStrings.Insert("apps")
	trieStrings.Insert("apper")
	trieStrings.Insert("apply")
	trieStrings.Insert("app")
	trieStrings.Insert("apricot")
	trieStrings.Insert("banana")
	trieStrings.Insert("band")
	trieStrings.Insert("bandana")

	assert.Equal(t, 7, trieStrings.CountByPrefix("app"))
	assert.Equal(t, 8, trieStrings.CountByPrefix("ap"))
	assert.Equal(t, 3, trieStrings.CountByPrefix("ban"))
	assert.Equal(t, 2, trieStrings.CountByPrefix("band"))
	assert.Equal(t, 0, trieStrings.CountByPrefix("cat"))

	trieInts := NewTrie(intComparator, intToString)

	trieInts.Insert(123)
	trieInts.Insert(1)
	trieInts.Insert(162)
	trieInts.Insert(18)
	trieInts.Insert(199)
	trieInts.Insert(124)
	trieInts.Insert(125)
	trieInts.Insert(456)
	trieInts.Insert(457)
	trieInts.Insert(789)

	assert.Equal(t, 7, trieInts.CountByPrefix(1))
	assert.Equal(t, 3, trieInts.CountByPrefix(12))
	assert.Equal(t, 2, trieInts.CountByPrefix(45))
	assert.Equal(t, 1, trieInts.CountByPrefix(78))
	assert.Equal(t, 0, trieInts.CountByPrefix(99))
}

func TestTrie_InsertAndSearch_Int(t *testing.T) {
	trieInts := NewTrie(intComparator, intToString)

	trieInts.Insert(123)
	trieInts.Insert(1)
	trieInts.Insert(124)
	trieInts.Insert(125)
	trieInts.Insert(456)
	trieInts.Insert(457)
	trieInts.Insert(789)

	assert.True(t, trieInts.Search(123))
	assert.True(t, trieInts.Search(124))
	assert.True(t, trieInts.Search(125))
	assert.True(t, trieInts.Search(456))
	assert.True(t, trieInts.Search(457))
	assert.True(t, trieInts.Search(789))
	assert.True(t, trieInts.Search(1))
	assert.False(t, trieInts.Search(999))
}
