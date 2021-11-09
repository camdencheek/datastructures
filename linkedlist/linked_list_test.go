package linkedlist

import (
	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinkedList(t *testing.T) {
	t.Run("NewLinkedList", func(t *testing.T) {
		t.Run("zero items", func(t *testing.T) {
			ll := NewLinkedList[int]()
			require.True(t, equivalent(t, ll, []int{}))
		})
		t.Run("one item", func(t *testing.T) {
			ll := NewLinkedList(1)
			require.True(t, equivalent(t, ll, []int{1}))
		})
		t.Run("multiple items", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			require.True(t, equivalent(t, ll, []int{1, 2, 3}))
		})
	})

	t.Run("Push", func(t *testing.T) {
		t.Run("zero items", func(t *testing.T) {
			ll := NewLinkedList[int]()
			ll.Push(1)
			require.True(t, equivalent(t, ll, []int{1}))
		})

		t.Run("one item", func(t *testing.T) {
			ll := NewLinkedList[int](1)
			ll.Push(2)
			require.True(t, equivalent(t, ll, []int{1, 2}))
		})

		t.Run("multiple push", func(t *testing.T) {
			ll := NewLinkedList[int](1)
			ll.Push(2)
			ll.Push(3)
			require.True(t, equivalent(t, ll, []int{1, 2, 3}))
		})

		t.Run("pushing equals slice appending", func(t *testing.T) {
			f := func(items []int) bool {
				ll := NewLinkedList[int]()
				for _, item := range items {
					ll.Push(item)
				}
				return equivalent(t, ll, items)
			}
			require.NoError(t, quick.Check(f, nil))
		})
	})

	t.Run("Pop", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			ll := NewLinkedList[int]()
			require.Nil(t, ll.Pop())
		})

		t.Run("one element", func(t *testing.T) {
			ll := NewLinkedList(1)
			require.Equal(t, 1, *ll.Pop())
			require.Nil(t, ll.Pop())
		})

		t.Run("multiple elements", func(t *testing.T) {
			ll := NewLinkedList(1, 2)
			require.Equal(t, 2, *ll.Pop())
			require.Equal(t, 1, *ll.Pop())
			require.Nil(t, ll.Pop())
		})
	})

	t.Run("PopHead", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			ll := NewLinkedList[int]()
			require.Nil(t, ll.PopHead())
		})

		t.Run("one element", func(t *testing.T) {
			ll := NewLinkedList(1)
			require.Equal(t, 1, *ll.PopHead())
			require.Nil(t, ll.PopHead())
		})

		t.Run("multiple elements", func(t *testing.T) {
			ll := NewLinkedList(1, 2)
			require.Equal(t, 1, *ll.PopHead())
			require.Equal(t, 2, *ll.PopHead())
			require.Nil(t, ll.PopHead())
		})
	})

	t.Run("PushHead", func(t *testing.T) {
		t.Run("zero items", func(t *testing.T) {
			ll := NewLinkedList[int]()
			ll.PushHead(1)
			require.True(t, equivalent(t, ll, []int{1}))
		})

		t.Run("one item", func(t *testing.T) {
			ll := NewLinkedList[int](1)
			ll.PushHead(2)
			require.True(t, equivalent(t, ll, []int{2, 1}))
		})

		t.Run("multiple push", func(t *testing.T) {
			ll := NewLinkedList[int](1)
			ll.PushHead(2)
			ll.PushHead(3)
			require.True(t, equivalent(t, ll, []int{3, 2, 1}))
		})
	})

	t.Run("SetCurrent", func(t *testing.T) {
		t.Run("set head", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			cursor := ll.CursorHead()
			cursor.SetCurrent(4)
			require.True(t, equivalent(t, ll, []int{4, 2, 3}))
		})

		t.Run("set tail", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			cursor := ll.CursorTail()
			cursor.SetCurrent(4)
			require.True(t, equivalent(t, ll, []int{1, 2, 4}))
		})

		t.Run("set mid", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			cursor := ll.CursorHead()
			cursor.Next()
			cursor.SetCurrent(4)
			require.True(t, equivalent(t, ll, []int{1, 4, 3}))
		})
	})

	t.Run("Reverse", func(t *testing.T) {
		t.Run("zero", func(t *testing.T) {
			ll := NewLinkedList[int]()
			ll.Reverse()
			require.True(t, equivalent(t, ll, []int{}))
		})

		t.Run("one", func(t *testing.T) {
			ll := NewLinkedList(1)
			ll.Reverse()
			require.True(t, equivalent(t, ll, []int{1}))
		})

		t.Run("two", func(t *testing.T) {
			ll := NewLinkedList(1, 2)
			ll.Reverse()
			require.True(t, equivalent(t, ll, []int{2, 1}))
		})

		t.Run("three", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			ll.Reverse()
			require.True(t, equivalent(t, ll, []int{3, 2, 1}))
		})
	})

	t.Run("ToSlice", func(t *testing.T) {
		t.Run("roundtrip returns same values", func(t *testing.T) {
			f := func(items []int) bool {
				ll := NewLinkedList(items...)
				return reflect.DeepEqual(items, ll.ToSlice()) && ll.Len() == len(items)
			}
			require.NoError(t, quick.Check(f, nil))
		})
	})

	t.Run("Iter", func(t *testing.T) {
		t.Run("reverse forward iter equals reverse iter", func(t *testing.T) {
			f := func(items []int) []int {
				ll := NewLinkedList(items...)
				ll.Reverse()
				iter := ll.Iter()
				res := make([]int, 0, len(items))
				for iter.Next() {
					res = append(res, iter.Value())
				}
				return res
			}
			g := func(items []int) []int {
				ll := NewLinkedList(items...)
				iter := ll.IterReverse()
				res := make([]int, 0, len(items))
				for iter.Next() {
					res = append(res, iter.Value())
				}
				return res
			}
			err := quick.CheckEqual(f, g, nil)
			require.NoError(t, err)
		})
	})

}

// equivalent steps through each node in the given linked list and checks that it
// is forward connected, backward connected, and each node is equivalent to the same index
// in the slice.
func equivalent(t *testing.T, ll LinkedList[int], s []int) bool {
	cursor := ll.CursorGhost()
	for i := 0; i < len(s); i++ {
		assert.True(t, cursor.Next())
		assert.Equal(t, s[i], *cursor.Current())
		assert.Equal(t, i, *cursor.Index())
		if i == 0 {
			assert.False(t, cursor.Prev())
		} else {
			assert.True(t, cursor.Prev())
			assert.Equal(t, s[i-1], *cursor.Current())
		}
		assert.True(t, cursor.Next())
	}

	// Should hit the ghost node
	assert.False(t, cursor.Next())
	assert.Nil(t, cursor.Current())
	assert.Nil(t, cursor.Index())

	if len(s) > 0 {
		assert.True(t, cursor.Prev())
		assert.Equal(t, s[len(s)-1], *cursor.Current())
	} else {
		assert.False(t, cursor.Prev())
		assert.Nil(t, cursor.Current())
		assert.Nil(t, cursor.Index())
	}

	assert.Equal(t, len(s), ll.Len())
	assert.Equal(t, s, ll.ToSlice())
	return !t.Failed()
}
