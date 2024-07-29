package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque(t *testing.T) {
	var d Deque[int]
	d = NewDeque[int]()

	t.Run("Initial state", func(t *testing.T) {
		assert.True(t, d.IsEmpty())
		assert.Equal(t, 0, d.Size())
	})

	t.Run("Add elements to the front", func(t *testing.T) {
		assert.NoError(t, d.FrontEnqueue(1))
		assert.NoError(t, d.FrontEnqueue(2))
		assert.NoError(t, d.FrontEnqueue(3))
		assert.False(t, d.IsEmpty())
		assert.Equal(t, 3, d.Size())
	})

	t.Run("Add elements to the back", func(t *testing.T) {
		assert.NoError(t, d.Enqueue(4))
		assert.NoError(t, d.Enqueue(5))
		assert.NoError(t, d.Enqueue(6))
		assert.False(t, d.IsEmpty())
		assert.Equal(t, 6, d.Size())
	})

	t.Run("Remove elements from the front", func(t *testing.T) {
		val, err := d.FrontDequeue()
		assert.NoError(t, err)
		assert.Equal(t, 3, *val)
		val, err = d.FrontDequeue()
		assert.NoError(t, err)
		assert.Equal(t, 2, *val)
		val, err = d.FrontDequeue()
		assert.NoError(t, err)
		assert.Equal(t, 1, *val)
	})

	t.Run("Remove elements from the back", func(t *testing.T) {
		val, err := d.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 6, *val)
		val, err = d.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 5, *val)
		val, err = d.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 4, *val)
		assert.True(t, d.IsEmpty())
	})

	t.Run("Peek elements", func(t *testing.T) {
		assert.NoError(t, d.Enqueue(10))
		assert.NoError(t, d.Enqueue(20))
		val, err := d.Peek()
		assert.NoError(t, err)
		assert.Equal(t, 20, *val)
		val, err = d.FrontPeek()
		assert.NoError(t, err)
		assert.Equal(t, 10, *val)
	})

	t.Run("Complex operations", func(t *testing.T) {
		d = NewDeque[int]()
		assert.NoError(t, d.Enqueue(1))
		assert.NoError(t, d.FrontEnqueue(2))
		assert.NoError(t, d.Enqueue(3))
		assert.NoError(t, d.FrontEnqueue(4))
		val, err := d.FrontPeek()
		assert.NoError(t, err)
		assert.Equal(t, 4, *val)
		val, err = d.Peek()
		assert.NoError(t, err)
		assert.Equal(t, 3, *val)

		val, err = d.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 3, *val)
		val, err = d.FrontDequeue()
		assert.NoError(t, err)
		assert.Equal(t, 4, *val)
		assert.Equal(t, 2, d.Size())

		val, err = d.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 1, *val)
		val, err = d.FrontDequeue()
		assert.NoError(t, err)
		assert.Equal(t, 2, *val)
		assert.True(t, d.IsEmpty())
	})
}
