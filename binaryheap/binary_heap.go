package binaryheap

import (
	"fmt"
	"github.com/camdencheek/datastructures/compare"
	"github.com/camdencheek/datastructures/vector"
)

type BinaryHeap[T any] struct {
	data vec.Vec[T]
	cmp  compare.CompareFunc[T]
}

func New[T any](cmp compare.CompareFunc[T]) *BinaryHeap[T] {
	return &BinaryHeap[T]{cmp: cmp}
}

func (b *BinaryHeap[T]) Len() int {
	return len(b.data)
}

func (b *BinaryHeap[T]) Push(val T) {
	b.data.Push(val)
	b.siftUp(0, b.data.Len() - 1)
	fmt.Printf("b.data: %#v\n", b.data)
}

func (b *BinaryHeap[T]) Pop() *T {
	replacement := b.data.Pop()
	if replacement == nil {
		return nil
	}

	if b.data.Len() == 0 {
		return replacement
	}

	max := b.data[0]
	b.data[0] = *replacement
	b.siftDown(0)
	return &max
}

func (b *BinaryHeap[T]) Peek() *T {
	if b.data.Len() == 0 {
		return nil
	}
	return &b.data[0]
}

func (b *BinaryHeap[T]) siftDown(pos int) int {
	for {
		left := 2 * pos + 1
		right := 2 * pos + 2
		largest := pos

		if left < b.data.Len() && b.cmp.Greater(b.data[left], b.data[largest]) {
			largest = left
		}

		if right < b.data.Len() && b.cmp.Greater(b.data[right],  b.data[largest]) {
			largest = right
		}

		if largest == pos {
			break
		}

		b.data[largest], b.data[pos] = b.data[pos], b.data[largest]
		pos = largest
	}
	return pos
}

func (b *BinaryHeap[T]) siftUp(start, pos int) int {
	for pos > start {
		parent := (pos - 1) / 2
		if b.cmp.Less(b.data[pos], b.data[parent]) {
			break
		}
		b.data[parent], b.data[pos] = b.data[pos], b.data[parent]
		pos = parent
	}
	return pos
}
