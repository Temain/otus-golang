package runner4

import (
	"errors"
	"sync"
)

// Run выполняет задачи, одновременно не более n, пока не будут выполнены все или не будет получено m ошибок.
func Run(tasks []func() error, n int, m int) (err error) {
	if n < 0 || m < 0 {
		return errors.New("n and m can't be negative")
	}

	queueCh := make(chan func() error, len(tasks))
	var s, e int
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	// Запуск воркеров.
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go worker(queueCh, m, &e, &s, &mu)
		}()
	}

	// Отправка задач в очередь к воркерам.
	for _, task := range tasks {
		queueCh <- task
	}

	wg.Wait()
	close(queueCh)

	if e == m {
		err = errors.New("to many errors, stopped")
	}

	if e > m {
		err = errors.New("errors limit exceed")
	}

	if e == m && e+s > n+m {
		err = errors.New("bad count of completed tasks")
	}

	return err
}

func worker(queueCh <-chan func() error, errorsLimit int, errorsCount *int, successCount *int, mu *sync.Mutex) {
	for task := range queueCh {
		err := task()
		mu.Lock()
		if err != nil {
			*errorsCount++
			if *errorsCount >= errorsLimit {
				mu.Unlock()
				return
			}
		} else {
			*successCount++
		}
		mu.Unlock()
	}
}
