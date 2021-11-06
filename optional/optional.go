package optional

type Optional[T any] struct {
	inner T
	isSome bool
}

func None[T any]() Optional[T] {
	return Optional[T]{}
}

func Some[T any](item T) Optional[T] {
	return Optional[T]{inner: item, isSome: true}
}

func (o Optional[T]) InnerMut() *T {
	if !o.isSome {
		return nil
	}
	return &o.inner
}

func (o Optional[T]) Unwrap() (res T) {
	return o.inner
}

func (o Optional[T]) TryUnwrap() (res T, ok bool) {
	return o.inner, o.isSome
}

func (o Optional[T]) IsSome() bool {
	return o.isSome
}

func (o Optional[T]) IsNone() bool {
	return !o.isSome
}
