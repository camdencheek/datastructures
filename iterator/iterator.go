package iterator

type Iterator[T any] interface {
	Next() bool
	Value() T
}
