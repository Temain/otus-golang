package runner3

import (
	"errors"
	"sync"
)

// Run выполняет задачи, одновременно не более n, пока не будут выполнены все или не будет получено m ошибок.
func Run(tasks []func() error, n int, m int) (err error) {
	if n < 0 || m < 0 {
		return errors.New("n and m can't be negative")
	}

	guardCh := make(chan struct{}, n)
	mu := sync.Mutex{}
	var (
		wg   sync.WaitGroup
		s, e int
		kill bool
	)

	for _, task := range tasks {
		guardCh <- struct{}{}
		if kill {
			err = errors.New("too many errors")
			break
		}

		wg.Add(1)
		go func(task func() error) {
			defer wg.Done()
			defer func() {
				<-guardCh
			}()
			err := task()
			checkResults(&s, &e, &n, &m, &kill, &mu, err)
		}(task)
	}

	close(guardCh)
	wg.Wait()

	return err
}

func checkResults(s *int, e *int, n *int, m *int, kill *bool, mu *sync.Mutex, err error) {
	mu.Lock()
	defer mu.Unlock()

	if err != nil {
		*e++
		if *e >= *m {
			*kill = true
		}
		if *e > 0 && *e+*s > *n+*m {
			*kill = true
		}
	} else {
		*s++
	}
}
