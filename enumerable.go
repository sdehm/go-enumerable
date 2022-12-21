package enumerable

type enumerable[T comparable] struct {
	values []T
	stack  []func(Enumerable[T]) Enumerable[T]
}

// nested enumerable interface
type Enumerable[T comparable] interface {
	Append(value T) Enumerable[T]
	Map(f func(T) T) Enumerable[T]
	Reverse() Enumerable[T]
	Filter(f func(T) bool) Enumerable[T]
	Take(n int) Enumerable[T]
	TakeWhile(f func(T) bool) Enumerable[T]
	Skip(n int) Enumerable[T]
	SkipWhile(f func(T) bool) Enumerable[T]
	Contains(value T) bool
	Any(f func(T) bool) bool
	All(f func(T) bool) bool
	Reduce(f func(T, T) T) T
	ForEach(f func(T))
	Apply() Enumerable[T]
	ToList() []T
	lazy(f func(Enumerable[T]) Enumerable[T]) Enumerable[T]
	getValues() []T
	setValues(values []T) Enumerable[T]
	reverseValues() Enumerable[T]
	// Parallel functions
	ForEachParallel(f func(T), numWorkers ...int)
	MapParallel(f func(T) T, numWorkers ...int) Enumerable[T]
}

// Create a new Enumerable[T] from a slice of T
func New[T comparable](values []T) Enumerable[T] {
	return enumerable[T]{values, []func(Enumerable[T]) Enumerable[T]{}}
}

// Append a value to the Enumerable[T] and return a new Enumerable[T]
// Evaluates lazily, call apply to evaluate
func (e enumerable[T]) Append(value T) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
		return e.setValues(append(e.getValues(), value))
	})
}

// Map a function over the Enumerable[T], returning a new Enumerable[T]
// Evaluates lazily, call apply to evaluate
func (e enumerable[T]) Map(f func(T) T) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
		values := e.getValues()
		for i, v := range values {
			values[i] = f(v)
		}
		return e.setValues(values)
	})
}

// Reverse the order of the Enumerable[T]
// Evaluates lazily, call apply to evaluate
func (e enumerable[T]) Reverse() Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
		return e.reverseValues()
	})
}

// Filter an Enumerable[T] by a predicate function
// Evaluates lazily, call apply to evaluate
func (e enumerable[T]) Filter(f func(T) bool) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
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

// Take the first n values of the Enumerable[T]
// If n is greater than the length of the Enumerable[T], returns the Enumerable[T]
// If n is negative, returns the last n values of the Enumerable[T]
// Evaluates lazily, call apply to evaluate
func (e enumerable[T]) Take(n int) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
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

// Take the first n values of the Enumerable[T] that satisfy a predicate function
// Evaluates lazily, call apply to evaluate
func (e enumerable[T]) TakeWhile(f func(T) bool) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
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

// Skip the first n values of the Enumerable[T]
// If n is greater than the length of the Enumerable[T], returns an empty Enumerable[T]
// If n is negative, returns all but the last n values of the Enumerable[T]
// Evaluates lazily, call apply to evaluate
func (e enumerable[T]) Skip(n int) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
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

// Skip the first values of the Enumerable[T] that satisfy a predicate function
// Evaluates lazily, call apply to evaluate
func (e enumerable[T]) SkipWhile(f func(T) bool) Enumerable[T] {
	return e.lazy(func(e Enumerable[T]) Enumerable[T] {
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

// Contains returns true if the Enumerable[T] contains the value
func (e enumerable[T]) Contains(value T) bool {
	for _, v := range e.Apply().getValues() {
		if v == value {
			return true
		}
	}
	return false
}

// Any returns true if the Enumerable[T] contains a value that satisfies the predicate
func (e enumerable[T]) Any(f func(T) bool) bool {
	for _, v := range e.Apply().getValues() {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all values in the Enumerable[T] satisfy the predicate
func (e enumerable[T]) All(f func(T) bool) bool {
	result := true
	for _, v := range e.Apply().getValues() {
		if !f(v) {
			result = false
		}
	}
	return result
}

// Reduce the Enumerable[T] to a single value
func (e enumerable[T]) Reduce(f func(T, T) T) T {
	e = e.Apply().(enumerable[T])
	result := e.values[0]
	for _, v := range e.values[1:] {
		result = f(result, v)
	}
	return result
}

// Iterate over the Enumerable[T], calling the function for each value
func (e enumerable[T]) ForEach(f func(T)) {
	for _, v := range e.Apply().getValues() {
		f(v)
	}
}

// Map a function over the Enumerable[T] but return a new IEnumerable of a different type
func Transform[T comparable, U comparable](e Enumerable[T], f func(T) U) Enumerable[U] {
	// TODO: Lazy evaluation broken here, would need some way to convert the stack to the new type
	values := e.Apply().getValues()
	newValues := make([]U, len(values))
	for i, v := range values {
		newValues[i] = f(v)
	}
	return New(newValues)
}

// Apply all the functions from the stack to the Enumerable[T]
func (e enumerable[T]) Apply() Enumerable[T] {
	for _, f := range e.stack {
		// recreate the enumerable without the stack
		e = enumerable[T]{values: e.values}
		e = f(e).(enumerable[T])
	}
	return e
}

// Apply any pending operations and return the values as a slice
func (e enumerable[T]) ToList() []T {
	return e.Apply().getValues()
}

func (e enumerable[T]) lazy(f func(Enumerable[T]) Enumerable[T]) Enumerable[T] {
	e.stack = append(e.stack, f)
	return e
}

func (e enumerable[T]) getValues() []T {
	return e.values
}

func (e enumerable[T]) setValues(values []T) Enumerable[T] {
	e.values = values
	return e
}

func (e enumerable[T]) reverseValues() Enumerable[T] {
	e.values = reverse(e.values)
	return e
}

func reverse[T comparable](values []T) []T {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return values
}