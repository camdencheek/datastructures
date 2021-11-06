package linkedlist

type Cursor[T any] struct {
	ll *LinkedList[T]
	// current is nil when pointing to the ghost node
	current *node[T]
	// index is only valid when current != nil
	index int
}

func (c *Cursor[T]) Current() *T {
	if c.current == nil {
		return nil
	}
	return &c.current.val
}

func (c *Cursor[T]) Index() *int {
	if c.current == nil {
		return nil
	}
	return &c.index
}

func (c *Cursor[T]) Next() bool {
	if c.current != nil {
		c.current = c.current.next
		c.index++
	} else {
		c.current = c.ll.head
		c.index = 0
	}
	return c.current != nil
}

func (c *Cursor[T]) Prev() bool {
	if c.current != nil {
		c.current = c.current.prev
		c.index--
	} else {
		c.current = c.ll.tail
		c.index = c.ll.length - 1
	}
	return c.current != nil
}

func (c *Cursor[T]) InsertBefore(item T) {
	if c.current == nil {
		c.ll.Push(item)
		return
	}
	n := newNode(item, c.current.prev, c.current)
	if c.current.prev == nil {
		c.ll.head = n
	} else {
		c.current.prev.next = n
	}
	c.current.prev = n
	c.index++
	c.ll.length++
}

func (c *Cursor[T]) InsertAfter(item T) {
	if c.current == nil {
		c.ll.PushHead(item)
		return
	}
	n := newNode(item, c.current, c.current.next)
	if c.current.next == nil {
		c.ll.tail = n
	} else {
		c.current.next.prev = n
	}
	c.current.next = n
	c.ll.length++
}

func (c *Cursor[T]) RemoveCurrent() *T {
	if c.current == nil {
		return nil
	}
	unlinked := c.current
	c.current = unlinked.next

	if unlinked.prev != nil {
		unlinked.prev.next = unlinked.next
	} else {
		c.ll.head = unlinked.next
	}

	if unlinked.next != nil {
		unlinked.next.prev = unlinked.prev
	} else {
		c.ll.tail = unlinked.prev
	}

	c.ll.length--

	return &unlinked.val
}
