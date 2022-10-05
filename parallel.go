package enumerable

func (e Enumerable[T]) ForEachParallel(f func(T)) {
	jobs := make(chan T, len(e.values))
	results := make(chan bool, len(e.values))

	go func() {
		for _, value := range e.values {
			jobs <- value
		}
		close(jobs)
	}()

	go worker(jobs, results, f)
	<-results
}

func worker[T any](jobs <-chan T, results chan<- bool, f func(T)) {
	for j := range jobs {
		f(j)
		results <- true
	}
	close(results)
}
