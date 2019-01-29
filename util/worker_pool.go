package util

import "sync"

// WorkerPool will launch ws workers at the same time and wait for them all to
// complete.
func WorkerPool(ws int, fn func(int)) {
	wg := sync.WaitGroup{}
	for w := 0; w < ws; w++ {
		wg.Add(1)
		go func(w int) {
			fn(w)
			wg.Done()
		}(w)
	}

	wg.Wait()
}
