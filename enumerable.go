package enumerable

type Enumerable[T comparable] struct {
	values []T
	stack  []func(IEnumerable[T]) IEnumerable[T]
}

// nested enumerable interface
type IEnumerable[T comparable] interface {
	Append(value T) IEnumerable[T]
	Map(f func(T) T) IEnumerable[T]
	Reverse() IEnumerable[T]
	Filter(f func(T) bool) IEnumerable[T]
	Take(n int) IEnumerable[T]
	TakeWhile(f func(T) bool) IEnumerable[T]
	Skip(n int) IEnumerable[T]
	SkipWhile(f func(T) bool) IEnumerable[T]
	Contains(value T) bool
	Any(f func(T) bool) bool
	All(f func(T) bool) bool
	Reduce(f func(T, T) T) T
	ForEach(f func(T))
	Apply() IEnumerable[T]
	ToList() []T
	lazy(f func(IEnumerable[T]) IEnumerable[T]) IEnumerable[T]
	getValues() []T
	setValues(values []T) IEnumerable[T]
	reverseValues() IEnumerable[T]
	// Parallel functions
	ForEachParallel(f func(T), numWorkers ...int)
	MapParallel(f func(T) T, numWorkers ...int) IEnumerable[T]
}

// Create a new IEnumerable[T] from a slice of T
func New[T comparable](values []T) IEnumerable[T] {
	return Enumerable[T]{values, []func(IEnumerable[T]) IEnumerable[T]{}}
}

// Append a value to the IEnumerable[T] and return a new IEnumerable[T]
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Append(value T) IEnumerable[T] {
	return e.lazy(func(e IEnumerable[T]) IEnumerable[T] {
		return e.setValues(append(e.getValues(), value))
	})
}

// Map a function over the IEnumerable[T], returning a new IEnumerable[T]
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Map(f func(T) T) IEnumerable[T] {
	return e.lazy(func(e IEnumerable[T]) IEnumerable[T] {
		values := e.getValues()
		for i, v := range values {
			values[i] = f(v)
		}
		return e.setValues(values)
	})
}

// Reverse the order of the IEnumerable[T]
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Reverse() IEnumerable[T] {
	return e.lazy(func(e IEnumerable[T]) IEnumerable[T] {
		return e.reverseValues()
	})
}

// Filter an IEnumerable[T] by a predicate function
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Filter(f func(T) bool) IEnumerable[T] {
	return e.lazy(func(e IEnumerable[T]) IEnumerable[T] {
		values := e.getValues()
		index := 0
		for _, v := range values {
			if f(v) {
				values = append(values[:index], v)
				index++
			}
		}
		return e.setValues(values[:index])
	})
}

// Take the first n values of the IEnumerable[T]
// If n is greater than the length of the IEnumerable[T], returns the IEnumerable[T]
// If n is negative, returns the last n values of the IEnumerable[T]
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Take(n int) IEnumerable[T] {
	return e.lazy(func(e IEnumerable[T]) IEnumerable[T] {
		values := e.getValues()
		reversed := false
		if n < 0 {
			values = reverse(values)
			n = -n
			reversed = true
		}
		if n > len(values) {
			n = len(values)
		}
		values = values[:n]
		if reversed {
			values = reverse(values)
		}
		return e.setValues(values)
	})
}

// Take the first n values of the IEnumerable[T] that satisfy a predicate function
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) TakeWhile(f func(T) bool) IEnumerable[T] {
	return e.lazy(func(e IEnumerable[T]) IEnumerable[T] {
		values := e.getValues()
		index := 0
		for i, v := range values {
			if !f(v) {
				index = i
				break
			}
		}
		values = values[:index]
		return e.setValues(values)
	})
}

// Skip the first n values of the IEnumerable[T]
// If n is greater than the length of the IEnumerable[T], returns an empty IEnumerable[T]
// If n is negative, returns all but the last n values of the IEnumerable[T]
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) Skip(n int) IEnumerable[T] {
	return e.lazy(func(e IEnumerable[T]) IEnumerable[T] {
		values := e.getValues()
		reversed := false
		if n < 0 {
			values = reverse(values)
			n = -n
			reversed = true
		}
		if n > len(values) {
			n = len(values)
		}
		values = values[n:]
		if reversed {
			values = reverse(values)
		}
		return e.setValues(values)
	})
}

// Skip the first values of the IEnumerable[T] that satisfy a predicate function
// Evaluates lazily, call apply to evaluate
func (e Enumerable[T]) SkipWhile(f func(T) bool) IEnumerable[T] {
	return e.lazy(func(e IEnumerable[T]) IEnumerable[T] {
		values := e.getValues()
		index := 0
		for i, v := range values {
			if !f(v) {
				index = i
				break
			}
		}
		values = values[index:]
		return e.setValues(values)
	})
}

// Contains returns true if the IEnumerable[T] contains the value
func (e Enumerable[T]) Contains(value T) bool {
	for _, v := range e.Apply().getValues() {
		if v == value {
			return true
		}
	}
	return false
}

// Any returns true if the IEnumerable[T] contains a value that satisfies the predicate
func (e Enumerable[T]) Any(f func(T) bool) bool {
	for _, v := range e.Apply().getValues() {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all values in the IEnumerable[T] satisfy the predicate
func (e Enumerable[T]) All(f func(T) bool) bool {
	result := true
	for _, v := range e.Apply().getValues() {
		if !f(v) {
			result = false
		}
	}
	return result
}

// Reduce the IEnumerable[T] to a single value
func (e Enumerable[T]) Reduce(f func(T, T) T) T {
	e = e.Apply().(Enumerable[T])
	result := e.values[0]
	for _, v := range e.values[1:] {
		result = f(result, v)
	}
	return result
}

// Iterate over the IEnumerable[T], calling the function for each value
func (e Enumerable[T]) ForEach(f func(T)) {
	for _, v := range e.Apply().getValues() {
		f(v)
	}
}

// Map a function over the IEnumerable[T] but return a new IEnumerable of a different type
func Transform[T comparable, U comparable](e IEnumerable[T], f func(T) U) IEnumerable[U] {
	result := New([]U{})
	for _, v := range e.Apply().getValues() {
		// TODO: Lazy evaluation broken here
		result = result.Append(f(v))
	}
	return result.Apply()
}

// Apply all the functions from the stack to the IEnumerable[T]
func (e Enumerable[T]) Apply() IEnumerable[T] {
	for _, f := range e.stack {
		// recreate the enumerable without the stack
		e = Enumerable[T]{values: e.values}
		e = f(e).(Enumerable[T])
	}
	return e
}

// Apply any pending operations and return the values as a slice
func (e Enumerable[T]) ToList() []T {
	return e.Apply().getValues()
}

func (e Enumerable[T]) lazy(f func(IEnumerable[T]) IEnumerable[T]) IEnumerable[T] {
	e.stack = append(e.stack, f)
	return e
}

func (e Enumerable[T]) getValues() []T {
	return e.values
}

func (e Enumerable[T]) setValues(values []T) IEnumerable[T] {
	e.values = values
	return e
}

func (e Enumerable[T]) reverseValues() IEnumerable[T] {
	e.values = reverse(e.values)
	return e
}

func reverse[T comparable](values []T) []T {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return values
}