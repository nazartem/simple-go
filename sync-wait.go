/* Пример использования sync.WaitGroup для ожидания завершения горутин */

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	iterationsNum  = 7
	goroutinesNum2 = 5
)

func formatWork2(workerNum, iterationsNum int) string {
	return fmt.Sprintf("воркер номер %d; итерация: %d", workerNum, iterationsNum)
}

func startWorker2(in int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик на 1

	for j := 0; j < iterationsNum; j++ {
		fmt.Println(formatWork2(in, j))
		runtime.Gosched()
	}
}

func syncWait() {
	wg := &sync.WaitGroup{} // инициализируем группу

	for i := 0; i < goroutinesNum2; i++ {
		wg.Add(1) // добавляем воркер
		go startWorker2(i, wg)
	}

	time.Sleep(time.Millisecond)
	wg.Wait() // ожидаем, пока waiter.Done() не приведёт счетчик к 0
}
