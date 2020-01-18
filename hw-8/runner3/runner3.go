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
		err = checkErrors(&e, &m, &mu)
		if err != nil {
			break
		}

		wg.Add(1)
		go func(task func() error) {
			defer wg.Done()
			defer func() {
				<-guardCh
			}()
			err := task()
			checkResults(&s, &e, &mu, err)
		}(task)
	}

	close(guardCh)
	wg.Wait()

	return err, s, e
}

func checkErrors(e *int, m *int, mu *sync.RWMutex) error {
	mu.RLock()
	defer mu.RUnlock()
	if *e >= *m {
		return errors.New("too many errors")
	}
	return nil
}

func checkResults(s *int, e *int, mu *sync.RWMutex, err error) {
	mu.Lock()
	defer mu.Unlock()
	if err != nil {
		*e++
	} else {
		*s++
	}
}
