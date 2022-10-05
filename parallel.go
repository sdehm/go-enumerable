package enumerable

import "sync"

func (e Enumerable[T]) ForEachParallel(f func(T)) {
	var wg sync.WaitGroup
	wg.Add(len(e.values))
	for _, v := range e.values {
		go func(v T) {
			f(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
