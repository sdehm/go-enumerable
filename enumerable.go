package enumerable

type Enumerable[T any] struct {
	values []T
	stack  []func(Enumerable[T]) Enumerable[T]
}

// Create a new Enumerable[T] from a slice of T
func New[T any](values []T) Enumerable[T] {
	return Enumerable[T]{values, []func(Enumerable[T]) Enumerable[T]{}}
}

// Append a value to the Enumerable[T] and return a new Enumerable[T]
func (e Enumerable[T]) Append(value T) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
		e.values = append(e.values, value)
		return e
	})
}

// Map a function over the Enumerable[T], returning a new Enumerable[T]
func (e Enumerable[T]) Map(f func(T) T) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
		for i, v := range e.values {
			e.values[i] = f(v)
		}
		return e
	})
}

// Reduce the Enumerable[T] to a single value
func (e Enumerable[T]) Reduce(f func(T, T) T) T {
	result := e.values[0]
	for _, v := range e.values[1:] {
		result = f(result, v)
	}
	return result
}

// Iterate over the Enumerable[T], calling the function for each value
func (e Enumerable[T]) ForEach(f func(T)) {
	for _, v := range e.values {
		f(v)
	}
}

// Map a function over the Enumerable[T] but return a new Enumerable of a different type
func Transform[T any, U any](e Enumerable[T], f func(T) U) Enumerable[U] {
	result := New([]U{})
	for _, v := range e.values {
		// TODO: Lazy evaluation broken here
		result = result.Append(f(v))
	}
	return result.Apply()
}

// Reverse the order of the Enumerable[T]
func (e Enumerable[T]) Reverse() Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
		result := New([]T{})
		for i := len(e.values) - 1; i >= 0; i-- {
			result.values = append(result.values, e.values[i])
		}
		return result
	})
}

// Filter an Enumerable[T] by a predicate function
func (e Enumerable[T]) Filter(f func(T) bool) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
		index := 0
		for _, v := range e.values {
			if f(v) {
				e.values = append(e.values[:index], v)
				index++
			}
		}
		e.values = e.values[:index]
		return e
	})
}

func (e Enumerable[T]) Apply() Enumerable[T] {
	for _, f := range e.stack {
		e = f(e)
	}
	return e
}

// Apply any pending operations and return the values as a slice
func (e Enumerable[T]) ToList() []T {
	return e.Apply().values
}

func (e Enumerable[T]) lazy(f func(Enumerable[T]) Enumerable[T]) Enumerable[T] {
	e.stack = append(e.stack, f)
	return e
}
