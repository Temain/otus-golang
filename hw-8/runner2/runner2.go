package runner2

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
	for i := 1; i <= n; i++ {
		go worker(i, queueCh, doneCh, killCh)
	}

	// Отправка задач в очередь к воркерам.
	go func() {
		for _, task := range tasks {
			queueCh <- task
		}
	}()

	// Проверка результатов выполнения задач.
	for range tasks {
		done := <-doneCh
		if !done {
			e++
			if e == m {
				err = errors.New("too many errors")
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
