package linkedlist

type iterator[T any] struct {
	cursor  *Cursor[T]
	reverse bool
}

func (i *iterator[T]) Next() bool {
	if i.reverse {
		return i.cursor.Prev()
	}
	return i.cursor.Next()
}

func (i *iterator[T]) Value() T {
	return *i.cursor.Current()
}
