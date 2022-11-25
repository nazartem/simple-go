/* Пример ограничения по ресурсам с помощью буферизированных каналов. Размер буфера равен
ограничению на количество горутин. Если канал quotaCh заполнен, то лимит воркеров уже работают. */

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	iterationsNum3 = 5
	goroutinesNum3 = 5
	quotaLimit     = 2
)

func formatWork3(workerNum, iterationsNum3 int) string {
	return fmt.Sprintf("воркер номер %d; итерация: %d", workerNum, iterationsNum3)
}

func startWorker3(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{} // берём свободный слот
	defer wg.Done()

	for j := 0; j < iterationsNum3; j++ {
		fmt.Println(formatWork3(in, j))
		if j%2 == 0 {
			<-quotaCh             // возвращаем слот
			quotaCh <- struct{}{} // берём слот
		}
		runtime.Gosched() // даём поработать другим горутинам
	}

	<-quotaCh // возвращаем слот
}

func workersLimit() {
	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit)

	for i := 0; i < goroutinesNum3; i++ {
		wg.Add(1)
		go startWorker3(i, wg, quotaCh)
	}

	time.Sleep(time.Millisecond)
	wg.Wait()
}
