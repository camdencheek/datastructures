package vec

type Vec[T any] []T

func New[T any](vals ...T) Vec[T] {
	return Vec[T](vals)
}

func NewFromSlice[T any](vals []T) Vec[T] {
	return Vec[T](vals)
}

// Cap returns the capacity of the vec, or the number
// of items the vec can store before reallocating
func (v Vec[T]) Cap() int {
	return cap(v)
}

// Len returns the current number of items in the vec
func (v Vec[T]) Len() int {
	return len(v)
}

// Append appends the values of other to the end of the vec
func (v *Vec[T]) Append(other Vec[T]) {
	*v = append(*v, other...)
}

// AppendSlice appends the values of other to the end of the vec
func (v *Vec[T]) AppendSlice(other []T) {
	*v = append(*v, other...)
}

// Clear removes all values from the vec without changing
// its capacity
func (v *Vec[T]) Clear() {
	*v = (*v)[:0]
}

// Insert inserts the value into the vec at the given index,
// shifting all elements after it to the right.
// Insert panics if i > v.Len().
func (v *Vec[T]) Insert(i int, val T) {
	// add an element to grow the slice as needed
	*v = append(*v, val)
	// shift following elements right by one
	copy((*v)[i+1:], (*v)[i:])
	// insert the value
	(*v)[i] = val
}

// Push adds the value to the end of the fector
func (v *Vec[T]) Push(val T) {
	*v = append(*v, val)
}

// Pop removes and returns the last value from the vec,
// or nil if the vec has no values. 
func (v *Vec[T]) Pop() *T {
	if len(*v) == 0 {
		return nil
	}
	popped := (*v)[len(*v)-1]
	*v = (*v)[:len(*v)-1]
	return &popped
}

// PopOrZero removes the last element of the vec if it exists, 
// otherwise it returns the zero value of T.
func (v *Vec[T]) PopOrZero() (res T) {
	if len(*v) == 0 {
		return
	}
	popped := (*v)[len(*v)-1]
	*v = (*v)[:len(*v)-1]
	return popped
}

// Remove removes the value at index i from the vec, shifting all values
// after it to the left. It does not change the capacity of the vec.
func (v *Vec[T]) Remove(i int) {
	*v = append((*v)[:i], (*v)[i+1:]...)
}

// Reserve allocates enough capacity to add at least additional more values to
// the vec. It does nothing if the capacity of the vec can already handle
// that many additional values.
func (v *Vec[T]) Reserve(additional int) {
	requestedCap := len(*v) + additional
	if cap(*v) >= requestedCap {
		return
	}
	grown := make([]T, requestedCap)
	copy(grown, *v)
	*v = grown
}

// Grow increases the capacity of the vec to the given size. If the capacity
// of the vec is already larger than size, it does nothing.
func (v *Vec[T]) Grow(size int) {
	if cap(*v) >= size {
		return
	}
	grown := make([]T, size)
	copy(grown, *v)
	*v = grown
}

// ShrinkToFit shrinks the capacity of the vec to the length of the vec.
// If the length already equals capacity, it does nothing.
func (v *Vec[T]) ShrinkToFit() {
	if len(*v) == cap(*v) {
		return
	}
	shrunk := make([]T, len(*v))
	copy(shrunk, *v)
	*v = shrunk
}

// Truncate keeps `size` values, removing all values after that. It does
// not change the capacity of the vec.
func (v *Vec[T]) Truncate(size int) {
	if len(*v) <= size {
		return
	}
	*v = (*v)[:size]
}

// AsSlice returns the vec's backing slice
func (v Vec[T]) AsSlice() []T {
	return v
}

// AsCopiedSlice copies the values of the vec into a new slice
// and returns it.
func (v Vec[T]) AsCopiedSlice() []T {
	res := make([]T, len(v))
	copy(res, v)
	return res
}

// Copy returns a new vec that contains all the same elements as the vec.
func (v Vec[T]) Copy() Vec[T] {
	copied := make([]T, len(v))
	copy(copied, v)
	return NewFromSlice(copied)
}

// Filter removes all elements from the vec for which the given predicate
// function does not return true.
func (v *Vec[T]) Filter(predicate func(T) bool) {
	j := 0
	for i := 0; i < len(*v); {
		if predicate((*v)[i]) {
			(*v)[j] = (*v)[i]
			j++
		}
	}
	*v = (*v)[:j]
}
