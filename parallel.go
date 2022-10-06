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

	workerFunc := func(j workItem[T]) {
		f(j.value)
	}

	wg := sync.WaitGroup{}
	startWorkers(jobs, &wg, workerFunc, workers)

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

	workerFunc := func(j workItem[T]) {
		results <- workItem[T]{f(j.value), j.index}
	}

	wg := sync.WaitGroup{}
	startWorkers(jobs, &wg, workerFunc, workers)

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

func startWorkers[T comparable](jobs chan workItem[T], wg *sync.WaitGroup, f func(workItem[T]), workers int) {
	// start workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for j := range jobs {
				f(j)
			}
			wg.Done()
		}()
	}
}
