package compare

type CompareFunc[T any] func(T, T) Result

func (f CompareFunc[T]) Less(a, b T) bool {
	return f(a, b) == Less
}

func (f CompareFunc[T]) Greater(a, b T) bool {
	return f(a, b) == Greater
}

func (f CompareFunc[T]) Equal(a, b T) bool {
	return f(a, b) == Equal
}

func (f CompareFunc[T]) Invert() CompareFunc[T] {
	return func(a, b T) Result {
		switch f(a, b) {
		case Less:
			return Greater
		case Greater: 
			return Less
		case Equal:
			return Equal
		default:
			panic("unknown CompareResult")
		}
	}
}

type Result int8

const (
	Less Result = -1
	Equal Result = 0
	Greater Result = 1
)
