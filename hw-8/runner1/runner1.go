package runner1

import "errors"

// Run выполняет задачи, одновременно не более n, пока не будут выполнены все или не будет получено m ошибок.
func Run(tasks []func() error, n int, m int) (err error, s int, e int) {
	// Буферизированный канал с помощью которого ограничивается одновременное выполнение задач.
	guardCh := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		guardCh <- struct{}{}
	}

	doneCh := make(chan bool)
	waitAllCh := make(chan bool)

	// Проверка результатов выполнения задач.
	go func() {
		for range tasks {
			done := <-doneCh
			if !done {
				e++
				if e == m {
					err = errors.New("too many errors")
					close(guardCh)
					break
				}
			} else {
				s++
			}

			// Новая горутина может стартовать.
			guardCh <- struct{}{}
		}

		// Можно завершать работу.
		waitAllCh <- true
	}()

	// Запуск новой задачи, если в буферизированный канал была записана пустая структура.
	for _, task := range tasks {
		_, ok := <-guardCh
		if !ok {
			break
		}
		go func(task func() error) {
			err := task()
			doneCh <- err == nil
		}(task)
	}

	// Ожидание завершения работы.
	<-waitAllCh

	return err, s, e
}
