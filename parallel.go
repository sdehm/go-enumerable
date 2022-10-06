package enumerable

import (
	"runtime"
	"sync"
)

func (e Enumerable[T]) ForEachParallel(f func(T), numWorkers ...int) {
	// set number of workers to GOMAXPROCS by default
	workers := setNumWorkers(numWorkers...)

	jobs := buildJobQueue(e)
	results := make(chan struct{}, len(e.values))

	// start workers
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for j := range jobs {
				f(j.value)
			}
			wg.Done()
		}()
	}

	// wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// wait for all results to be processed
	<-results
}

func (e Enumerable[T]) MapParallel(f func(T) T, numWorkers ...int) Enumerable[T] {
	workers := setNumWorkers(numWorkers...)
	jobs := buildJobQueue(e)
	results := make(chan workItem[T], len(e.values))

	// start workers
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for j := range jobs {
				results <- workItem[T]{f(j.value), j.index}
			}
			wg.Done()
		}()
	}

	// wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		e.values[r.index] = r.value
	}

	return e
}

type workItem[T comparable] struct {
	value T
	index int
}

func setNumWorkers(numWorkers ...int) int {
	if len(numWorkers) > 0 {
		return numWorkers[0]
	}
	// set number of workers to GOMAXPROCS by default
	return runtime.GOMAXPROCS(0)
}

func buildJobQueue[T comparable](e Enumerable[T]) chan workItem[T] {
	jobs := make(chan workItem[T], len(e.values))
	// populate jobs channel
	go func() {
		for i, v := range e.values {
			jobs <- workItem[T]{v, i}
		}
		close(jobs)
	}()
	return jobs
}
