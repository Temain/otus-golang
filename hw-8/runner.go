package runner

import (
	"errors"
	"fmt"
	"time"
)

// Run выполняет задачи, одновременно не более n, пока не будут выполнены все или не будет получено m ошибок.
func Run(tasks []func() error, n int, m int) (err error, s int, e int) {
	queueCh := make(chan func() error)
	doneCh := make(chan bool)
	killCh := make(chan bool)

	// Запуск воркеров.
	for i := 0; i < n; i++ {
		go worker(queueCh, doneCh, killCh)
	}

	// Отправка задач в очередь к воркерам.
	for _, task := range tasks {
		go func(task func() error) {
			queueCh <- task
		}(task)
	}

	// Проверка результатов выполнения задач.
	for range tasks {
		done := <-doneCh
		if !done {
			e++
			if e == m {
				err = errors.New("превышено предельное количество ошибок")
				break
			}
		}
		s++
	}

	// Завершение работы воркеров.
	close(killCh)
	time.Sleep(2 * time.Second)

	return err, s, e
}

func worker(queueCh chan func() error, doneCh chan bool, killCh chan bool) {
	fmt.Println("worker started")
	for {
		select {
		case task := <-queueCh:
			err := task()
			doneCh <- err == nil
		case <-killCh:
			fmt.Println("worker stopped")
			return
		}
	}
}
