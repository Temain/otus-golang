package runner1

import (
	"errors"
	"sync"
)

// Run выполняет задачи, одновременно не более n, пока не будут выполнены все или не будет получено m ошибок.
func Run(tasks []func() error, n int, m int) (err error) {
	if n < 0 || m < 0 {
		return errors.New("n and m can't be negative")
	}

	// Буферизированный канал с помощью которого ограничивается одновременное выполнение задач.
	guardCh := make(chan struct{}, n)
	doneCh := make(chan bool, n)
	var s, e int
	var wg sync.WaitGroup

	// Запуск новой задачи, если в буферизированный канал была записана пустая структура.
	for _, task := range tasks {
		guardCh <- struct{}{}

		wg.Add(1)
		go func(task func() error) {
			defer wg.Done()
			defer func() {
				<-guardCh
			}()

			err := task()
			doneCh <- err == nil
		}(task)

		err = checkResults(doneCh, &n, &m, &s, &e)
		if err != nil {
			break
		}
	}

	wg.Wait()
	close(guardCh)
	close(doneCh)

	return err
}

func checkResults(doneCh <-chan bool, n *int, m *int, s *int, e *int) (err error) {
	done := <-doneCh
	if done {
		*s++
	} else {
		*e++
		if *e > *m {
			err = errors.New("errors limit exceed")
		}

		if *e == *m {
			err = errors.New("too many errors")
			if *e+*s > *n+*m {
				err = errors.New("bad count of completed tasks")
			}
			return err
		}
	}

	return err
}
