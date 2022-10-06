package enumerable

import "sync"

func (e Enumerable[T]) ForEachParallel(f func(T)) {
	jobs := make(chan T, len(e.values))
	results := make(chan bool, len(e.values))

	go func() {
		for _, v := range e.values {
			jobs <- v
		}
		close(jobs)
	}()

	go worker(jobs, results, f)
	<-results
}

func worker[T any](jobs <-chan T, results chan<- bool, f func(T)) {
	wg := sync.WaitGroup{}
	for j := range jobs {
		wg.Add(1)
		go func(j T) {
			defer wg.Done()
			f(j)
		}(j)
	}
	wg.Wait()
	close(results)
}

type workItem[T any] struct {
	value T
	index int
}

func (e Enumerable[T]) MapParallel(f func(T) T) Enumerable[T] {
	jobs := make(chan workItem[T], 3)
	results := make(chan workItem[T], 3)

	go func() {
		for i, v := range e.values {
			jobs <- workItem[T]{v, i}
		}
		close(jobs)
	}()

	go workerT(jobs, results, f)

	for r := range results {
		e.values[r.index] = r.value
	}

	return e
}

func workerT[T any](jobs <-chan workItem[T], results chan<- workItem[T], f func(T) T) {
	wg := sync.WaitGroup{}
	for j := range jobs {
		wg.Add(1)
		go func(j workItem[T]) {
			results <- workItem[T]{f(j.value), j.index}
			wg.Done()
		}(j)
	}
	wg.Wait()
	close(results)
}
