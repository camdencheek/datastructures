package linkedlist

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func TestCursor(t *testing.T) {
	t.Run("Current", func(t *testing.T) {
		ll := NewLinkedList(1, 2)

		t.Run("CursorHead", func(t *testing.T) {
			ch := ll.CursorHead()
			require.Equal(t, 1, *ch.Current())
			require.True(t, ch.Next())
			require.Equal(t, 2, *ch.Current())
			require.True(t, ch.Prev())
			require.Equal(t, 1, *ch.Current())
			require.True(t, ch.Next())
			require.Equal(t, 2, *ch.Current())
			require.False(t, ch.Next())
			require.Nil(t, ch.Current())
			require.True(t, ch.Next())
			require.Equal(t, 1, *ch.Current())
		})

		t.Run("CursorTail", func(t *testing.T) {
			ct := ll.CursorTail()
			require.Equal(t, 2, *ct.Current())
			require.True(t, ct.Prev())
			require.Equal(t, 1, *ct.Current())
			require.False(t, ct.Prev())
			require.Nil(t, ct.Current())
		})

		t.Run("empty list returns nil", func(t *testing.T) {
			t.Run("CursorHead", func(t *testing.T) {
				ll := NewLinkedList[int]()
				cursor := ll.CursorHead()
				require.Nil(t, cursor.Current())
				require.False(t, cursor.Next())
				require.Nil(t, cursor.Current())
				require.False(t, cursor.Prev())
				require.Nil(t, cursor.Current())
			})

			t.Run("CursorTail", func(t *testing.T) {
				ll := NewLinkedList[int]()
				cursor := ll.CursorTail()
				require.Nil(t, cursor.Current())
				require.False(t, cursor.Next())
				require.Nil(t, cursor.Current())
				require.False(t, cursor.Prev())
				require.Nil(t, cursor.Current())
			})
		})
	})

	t.Run("RemoveCurrent", func(t *testing.T) {
		t.Run("empty list", func(t *testing.T) {
			ll := NewLinkedList[int]()
			cursor := ll.CursorHead()
			require.Nil(t, cursor.RemoveCurrent())
			require.True(t, equivalent(t, ll, []int{}))
		})

		t.Run("single element", func(t *testing.T) {
			ll := NewLinkedList(1)
			cursor := ll.CursorHead()
			require.Equal(t, 1, *cursor.RemoveCurrent())
			require.True(t, equivalent(t, ll, []int{}))
		})

		t.Run("remove head", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			cursor := ll.CursorHead()
			require.Equal(t, 1, *cursor.RemoveCurrent())
			require.True(t, equivalent(t, ll, []int{2, 3}))
		})

		t.Run("remove middle element", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			cursor := ll.CursorHead()
			require.True(t, cursor.Next())
			require.Equal(t, 2, *cursor.RemoveCurrent())
			require.True(t, equivalent(t, ll, []int{1, 3}))
		})

		t.Run("remove last element", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			cursor := ll.CursorTail()
			require.Equal(t, 3, *cursor.RemoveCurrent())
			require.True(t, equivalent(t, ll, []int{1, 2}))
		})

		t.Run("remove ghost", func(t *testing.T) {
			ll := NewLinkedList(1, 2, 3)
			cursor := ll.CursorGhost()
			require.Nil(t, cursor.RemoveCurrent())
			require.True(t, equivalent(t, ll, []int{1, 2, 3}))
		})

		t.Run("quick check removing random from list and slice is equivalent", func(t *testing.T) {
			f := func(slice []int) bool {
				ll := NewLinkedList(slice...)
				cursor := ll.CursorHead()

				for i := 0; i < len(slice); {
					// Remove even nodes from both the slice and the linked list
					if slice[i]%2 == 0 {
						slice = append(slice[:i], slice[i+1:]...)
						cursor.RemoveCurrent()
						continue
					}
					cursor.Next()
					i++
				}
				// Expect that the chopped slice and the linked list are equivalent
				return equivalent(t, ll, slice)
			}
			require.NoError(t, quick.Check(f, nil))
		})
	})

	t.Run("InsertAfter", func(t *testing.T) {
		t.Run("empty list", func(t *testing.T) {
			ll := NewLinkedList[int]()
			cursor := ll.CursorHead()
			cursor.InsertAfter(1)
			require.True(t, equivalent(t, ll, []int{1}))
		})

		t.Run("nonempty list before head", func(t *testing.T) {
			ll := NewLinkedList[int](1)
			cursor := ll.CursorHead()
			cursor.InsertAfter(2)
			require.True(t, equivalent(t, ll, []int{1, 2}))
		})

		t.Run("nonempty list before tail", func(t *testing.T) {
			ll := NewLinkedList[int](1)
			cursor := ll.CursorTail()
			cursor.InsertAfter(2)
			require.True(t, equivalent(t, ll, []int{1, 2}))
		})

		t.Run("larger list before head", func(t *testing.T) {
			ll := NewLinkedList[int](1, 2)
			cursor := ll.CursorHead()
			cursor.InsertAfter(3)
			require.True(t, equivalent(t, ll, []int{1, 3, 2}))
		})

		t.Run("larger list before tail", func(t *testing.T) {
			ll := NewLinkedList[int](1, 2)
			cursor := ll.CursorTail()
			cursor.InsertAfter(3)
			require.True(t, equivalent(t, ll, []int{1, 2, 3}))
		})

		t.Run("before ghost is tail", func(t *testing.T) {
			ll := NewLinkedList[int](1, 2)
			cursor := ll.CursorGhost()
			cursor.InsertAfter(3)
			require.True(t, equivalent(t, ll, []int{3, 1, 2}))
		})

		t.Run("quick", func(t *testing.T) {
			f := func(slice []int) bool {
				ll := NewLinkedList(slice...)
				cursor := ll.CursorHead()

				for i := 0; i < len(slice); i++ {
					// If the item is even, duplicate it
					if slice[i]%2 == 0 {
						slice = append(slice[:i+1], slice[i:]...)
						i++
						cursor.InsertAfter(slice[i])
						cursor.Next()
					}
					cursor.Next()
				}
				// Expect that the chopped slice and the linked list are equivalent
				return equivalent(t, ll, slice)
			}
			require.NoError(t, quick.Check(f, nil))
		})
	})

	t.Run("InsertBefore", func(t *testing.T) {
		t.Run("empty list", func(t *testing.T) {
			ll := NewLinkedList[int]()
			cursor := ll.CursorHead()
			cursor.InsertBefore(1)
			require.True(t, equivalent(t, ll, []int{1}))
		})

		t.Run("nonempty list before head", func(t *testing.T) {
			ll := NewLinkedList[int](1)
			cursor := ll.CursorHead()
			cursor.InsertBefore(2)
			require.True(t, equivalent(t, ll, []int{2, 1}))
		})

		t.Run("nonempty list before tail", func(t *testing.T) {
			ll := NewLinkedList[int](1)
			cursor := ll.CursorTail()
			cursor.InsertBefore(2)
			require.True(t, equivalent(t, ll, []int{2, 1}))
		})

		t.Run("larger list before head", func(t *testing.T) {
			ll := NewLinkedList[int](1, 2)
			cursor := ll.CursorHead()
			cursor.InsertBefore(3)
			require.True(t, equivalent(t, ll, []int{3, 1, 2}))
		})

		t.Run("larger list before tail", func(t *testing.T) {
			ll := NewLinkedList[int](1, 2)
			cursor := ll.CursorTail()
			cursor.InsertBefore(3)
			require.True(t, equivalent(t, ll, []int{1, 3, 2}))
		})

		t.Run("before ghost is tail", func(t *testing.T) {
			ll := NewLinkedList[int](1, 2)
			cursor := ll.CursorGhost()
			cursor.InsertBefore(3)
			require.True(t, equivalent(t, ll, []int{1, 2, 3}))
		})

		t.Run("quick", func(t *testing.T) {
			f := func(slice []int) bool {
				ll := NewLinkedList(slice...)
				cursor := ll.CursorHead()

				for i := 0; i < len(slice); i++ {
					// If the item is even, duplicate it
					if slice[i]%2 == 0 {
						slice = append(slice[:i+1], slice[i:]...)
						i++
						cursor.InsertBefore(slice[i])
					}
					cursor.Next()
				}
				// Expect that the chopped slice and the linked list are equivalent
				return equivalent(t, ll, slice)
			}
			require.NoError(t, quick.Check(f, nil))
		})
	})
}
