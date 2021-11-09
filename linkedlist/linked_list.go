package linkedlist

import (
	. "github.com/camdencheek/datastructures/iterator"
)

type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

type node[T any] struct {
	prev *node[T]
	next *node[T]
	val  T
}

func newNode[T any](item T, prev, next *node[T]) *node[T] {
	return &node[T]{
		prev: prev,
		next: next,
		val:  item,
	}
}

func NewLinkedList[T any](items ...T) LinkedList[T] {
	var (
		prev *node[T]
		head *node[T]
		tail *node[T]
	)
	for i, item := range items {
		n := newNode(item, prev, nil)
		if i == 0 {
			head = n
		}
		if i == len(items)-1 {
			tail = n
		}
		if prev != nil {
			prev.next = n
		}
		prev = n
	}

	return LinkedList[T]{
		head:   head,
		tail:   tail,
		length: len(items),
	}
}

// Len returns the length of the linked list.
func (ll *LinkedList[T]) Len() int {
	return ll.length
}

// Push will add the given value as a node to the end of the linked list.
func (ll *LinkedList[T]) Push(item T) {
	n := newNode(item, ll.tail, nil)
	if ll.tail == nil {
		ll.head = n
		ll.tail = n
	} else {
		ll.tail.next = n
		ll.tail = n
	}
	ll.length++
}

// Pop will remove and return the last node of the linked list.
// If the list is empty, it will return nil.
func (ll *LinkedList[T]) Pop() *T {
	popped := ll.tail
	if popped == nil {
		return nil
	}
	ll.tail = popped.prev
	if ll.tail == nil {
		ll.head = nil
	}
	ll.length--
	return &popped.val
}

// PushHead will add the provided value as a node at the front of the linked list.
func (ll *LinkedList[T]) PushHead(item T) {
	n := newNode(item, nil, ll.head)
	if ll.head == nil {
		ll.head = n
		ll.tail = n
	} else {
		ll.head.prev = n
		ll.head = n
	}
	ll.length++
}

// PopHead will remove the first node of the linked list and return it.
// If the list is empty, it will return nil.
func (ll *LinkedList[T]) PopHead() *T {
	popped := ll.head
	if popped == nil {
		return nil
	}
	ll.head = popped.next
	if ll.head == nil {
		ll.tail = nil
	}
	ll.length--
	return &popped.val
}

// Reverse reverses the linked list
func (ll *LinkedList[T]) Reverse() {
	current := ll.head
	for current != nil {
		current.prev, current.next = current.next, current.prev
		current = current.prev
	}
	ll.head, ll.tail = ll.tail, ll.head
}

// ToSlice collects the nodes of the linked list and returns them as a slice.
func (ll *LinkedList[T]) ToSlice() []T {
	res := make([]T, 0, ll.length)
	iter := ll.Iter()
	for iter.Next() {
		res = append(res, iter.Value())
	}
	return res
}

func (ll *LinkedList[T]) Iter() Iterator[T] {
	return &iterator[T]{cursor: ll.CursorGhost()}
}

func (ll *LinkedList[T]) IterReverse() Iterator[T] {
	return &iterator[T]{cursor: ll.CursorGhost(), reverse: true}
}

// CursorHead returns a cursor pointing to the first node in the linked list.
// If the list is empty, the cursor will point to the "ghost" node.
func (ll *LinkedList[T]) CursorHead() *Cursor[T] {
	return &Cursor[T]{ll: ll, current: ll.head}
}

// CursorTail returns a cursor pointing to the last node in the linked list.
// If the list is empty, the cursor will point to the "ghost" node.
func (ll *LinkedList[T]) CursorTail() *Cursor[T] {
	return &Cursor[T]{ll: ll, current: ll.tail}
}

// CursorGhost returns a cursor pointing to the "ghost" node, which
// logically lies before head and after tail.
func (ll *LinkedList[T]) CursorGhost() *Cursor[T] {
	return &Cursor[T]{ll: ll, current: nil}
}
