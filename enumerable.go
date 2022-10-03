package enumerable

type	Enumerable[T any] struct {
	values []T
}

// Create a new Enumerable[T] from a slice of T
func New[T any](values []T) Enumerable[T] {
	return Enumerable[T]{values}
}

// Map a function over the Enumerable[T], returning a new Enumerable[T]
func (e Enumerable[T]) Map(f func(T) T) Enumerable[T]{
	for i, v := range e.values {
		e.values[i] = f(v)
	}
	return e
}

// Reduce the Enumerable[T] to a single value
func (e Enumerable[T]) Reduce(f func(T, T) T) T {
	result := e.values[0]
	for _, v := range e.values[1:] {
		result = f(result, v)
	}
	return result
}
