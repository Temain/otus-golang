package runner2

import (
	"errors"
	"fmt"
)

// Run выполняет задачи, одновременно не более n, пока не будут выполнены все или не будет получено m ошибок.
func Run(tasks []func() error, n int, m int) (err error) {
	if n < 0 || m < 0 {
		return errors.New("n and m can't be negative")
	}

	queueCh := make(chan func() error, len(tasks))
	doneCh := make(chan bool)
	killCh := make(chan bool)
	var s, e int

	// Запуск воркеров.
	for i := 1; i <= n; i++ {
		go worker(i, queueCh, doneCh, killCh)
	}

	// Отправка задач в очередь к воркерам.
	for _, task := range tasks {
		queueCh <- task
	}

	// Проверка результатов выполнения задач.
	for range tasks {
		done := <-doneCh
		if !done {
			e++
			if e > m {
				err = errors.New("errors limit exceed")
				break
			}

			if e == m {
				err = errors.New("too many errors")
				if e+s > n+m {
					err = errors.New("bad count of completed tasks")
				}
				break
			}
		}
		s++
	}

	// Завершение работы воркеров.
	close(killCh)

	return err
}

func worker(num int, queueCh <-chan func() error, doneCh chan<- bool, killCh <-chan bool) {
	fmt.Printf("worker %d started\n", num)
	for {
		select {
		case task := <-queueCh:
			fmt.Printf("worker %d running task...\n", num)
			err := task()
			doneCh <- err == nil
		case <-killCh:
			fmt.Printf("worker %d stopped\n", num)
			return
		}
	}
}
