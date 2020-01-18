package runner3

import (
	"errors"
	"sync"
)

// Run выполняет задачи, одновременно не более n, пока не будут выполнены все или не будет получено m ошибок.
func Run(tasks []func() error, n int, m int) (err error, s int, e int) {
	guardCh := make(chan struct{}, n)
	mu := sync.RWMutex{}
	var wg sync.WaitGroup
	for _, task := range tasks {
		guardCh <- struct{}{}
		mu.RLock()
		if e >= m {
			err = errors.New("too many errors")
			break
		}
		mu.RUnlock()

		wg.Add(1)
		go func(task func() error) {
			defer wg.Done()
			defer func() {
				<-guardCh
			}()
			err := task()
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				e++
			} else {
				s++
			}
		}(task)
	}

	close(guardCh)
	wg.Wait()

	return err, s, e
}
