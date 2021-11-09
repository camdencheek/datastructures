package queue

type Queue[T any] struct {
	data []T
	start, end int
}

func (q *Queue[T]) Push(val T) {
	if 
}

func (q *Queue[T]) Cap() int {
	return len(q.data)
}

func (q *Queue[T]) Len() int {
	return q.end - q.start
}

func (q *Queue[T]) Reserve(additional int) {
	if q.Cap() - q.Len() >= additional {
		return
	}	
	
	var newCap int
	switch len(q.data) {
	case 0, 1, 2:
		newCap = 4		
	case 3, 4:
		newCap = 8
	default:
		newCap = len(q.data) * 2
	}

	if newCap - q.Len() >= additional {
		newCap = q.Len() + additional
	}

	newData := make([]T, newCap)
	copy(newData, q.data[q.start:q.end])
	q.data = newData

	newEnd := q.end - q.start
	q.start = 0
	q.end = newEnd
}

func (q *Queue[T]) Enqueue(val T) {
	q.Reserve(1)
	q.data[q.end%len(q.data)] = val
	q.end++
}

func (q *Queue[T]) Dequeue() *T {
	if q.Len() == 0 {
		return nil
	}

	val := q.data[q.start]
	q.start++
	return &val
}
