/* Пример распараллеливания задачи. Три функции-воркера читают из канала и выполняют работу */

package main

import (
	"fmt"
	"runtime"
	"time"
)

const goroutinesNum = 3

func formatWork(workerNum int, input string) string {
	return fmt.Sprintf("воркер номер %d; прочитаны данные: %s", workerNum, input)
}

func printFinishWork(workerNum int) {
	fmt.Printf("воркер %d работу закончил\n", workerNum)
}

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Println(formatWork(workerNum, input))
		runtime.Gosched() // для переключения горутин
	}
	printFinishWork(workerNum)
}

func workerPull() {
	worketInput := make(chan string, 2) // попробуйте увеличить размер канала

	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, worketInput)
	}

	months := []string{"Январь", "Февраль", "Март",
		"Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь",
		"Октябрь", "Ноябрь", "Декабрь",
	}

	for _, monthName := range months {
		worketInput <- monthName
	}
	close(worketInput) // для завершения пула воркеров

	time.Sleep(time.Millisecond)
}
