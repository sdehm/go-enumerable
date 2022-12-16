package enumerable

// nested enumerable interface
type IEnumerable[T comparable] interface {
	Append(value T) Enumerable[T]
	Apply() Enumerable[T]
}

type Enumerable[T comparable] struct {
	values []T
	stack  []func(Enumerable[T]) Enumerable[T]
}

// Create a new Enumerable[T] from a slice of T
func New[T comparable](values []T) Enumerable[T] {
	return Enumerable[T]{values, []func(Enumerable[T]) Enumerable[T]{}}
}

// Append a value to the Enumerable[T] and return a new Enumerable[T]
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Append(value T) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
		e.values = append(e.values, value)
		return e
	})
}

// Map a function over the Enumerable[T], returning a new Enumerable[T]
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Map(f func(T) T) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
		for i, v := range e.values {
			e.values[i] = f(v)
		}
		return e
	})
}

// Reverse the order of the Enumerable[T]
// Evaluates lazily, call apply to evaluate
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
// Evaluates lazily, call apply to evaluate
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

// Take the first n values of the Enumerable[T]
// If n is greater than the length of the Enumerable[T], returns the Enumerable[T]
// If n is negative, returns the last n values of the Enumerable[T]
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Take(n int) Enumerable[T] {
	return e.lazy(func(Enumerable[T]) Enumerable[T] {
		reversed := false
		if n < 0 {
			e = e.Reverse().Apply()
			n = -n
			reversed = true
		}
		if n > len(e.values) {
			n = len(e.values)
		}
		e.values = e.values[:n]
		if reversed {
			e = e.Reverse().Apply()
		}
		return e
	})
}

// Take the first n values of the Enumerable[T] that satisfy a predicate function
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) TakeWhile(f func(T) bool) Enumerable[T] {
	return e.lazy(func(Enumerable[T]) Enumerable[T] {
		index := 0
		for i, v := range e.values {
			if !f(v) {
				index = i
				break
			}
		}
		e.values = e.values[:index]
		return e
	})
}

// Skip the first n values of the Enumerable[T]
// If n is greater than the length of the Enumerable[T], returns an empty Enumerable[T]
// If n is negative, returns all but the last n values of the Enumerable[T]
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Skip(n int) Enumerable[T] {
	return e.lazy(func(Enumerable[T]) Enumerable[T] {
		reversed := false
		if n < 0 {
			e = e.Reverse().Apply()
			n = -n
			reversed = true
		}
		if n > len(e.values) {
			n = len(e.values)
		}
		e.values = e.values[n:]
		if reversed {
			e = e.Reverse().Apply()
		}
		return e
	})
}

// Skip the first values of the Enumerable[T] that satisfy a predicate function
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) SkipWhile(f func(T) bool) Enumerable[T] {
	return e.lazy(func(Enumerable[T]) Enumerable[T] {
		index := 0
		for i, v := range e.values {
			if !f(v) {
				index = i
				break
			}
		}
		e.values = e.values[index:]
		return e
	})
}

// Contains returns true if the Enumerable[T] contains the value
func (e Enumerable[T]) Contains(value T) bool {
	for _, v := range e.Apply().values {
		if v == value {
			return true
		}
	}
	return false
}

// Any returns true if the Enumerable[T] contains a value that satisfies the predicate
func (e Enumerable[T]) Any(f func(T) bool) bool {
	for _, v := range e.Apply().values {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all values in the Enumerable[T] satisfy the predicate
func (e Enumerable[T]) All(f func(T) bool) bool {
	result := true
	for _, v := range e.Apply().values {
		if !f(v) {
			result = false
		}
	}
	return result
}

// Reduce the Enumerable[T] to a single value
func (e Enumerable[T]) Reduce(f func(T, T) T) T {
	e = e.Apply()
	result := e.values[0]
	for _, v := range e.values[1:] {
		result = f(result, v)
	}
	return result
}

// Iterate over the Enumerable[T], calling the function for each value
func (e Enumerable[T]) ForEach(f func(T)) {
	for _, v := range e.Apply().values {
		f(v)
	}
}

// Map a function over the Enumerable[T] but return a new Enumerable of a different type
func Transform[T comparable, U comparable](e Enumerable[T], f func(T) U) Enumerable[U] {
	result := New([]U{})
	for _, v := range e.Apply().values {
		// TODO: Lazy evaluation broken here
		result = result.Append(f(v))
	}
	return result.Apply()
}

// Apply all the functions from the stack to the Enumerable[T]
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
