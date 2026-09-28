package main

import (
	"fmt"
	"time"
)

func mine(transferpoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался")
	fmt.Println("Поход в шахту номер", n, "закончился")
	transferpoint <- 10
	fmt.Println("Рабочий", n, "уголь передал")
	fmt.Println("Я код и я пошёл дальше")
}

func main() {
	go func() {
		fmt.Println("Я анонимная горутина, могу работать и так")
	}()
	start := time.Now()
	fmt.Println(start)
	coal := 0
	transferpoint := make(chan int, 5) /* При добавлении ", 5)"
	горутина идёт дальше отдавая значение в буфер, т.о горутина может не ждать передачи,
	а положить его в буфер и продолжаться дальше */
	go mine(transferpoint, 1)
	go mine(transferpoint, 2)
	go mine(transferpoint, 3)
	go mine(transferpoint, 4)
	go mine(transferpoint, 5)

	coal += <-transferpoint

	time.Sleep(1 * time.Second)
	coal += <-transferpoint

	time.Sleep(1 * time.Second)
	coal += <-transferpoint

	time.Sleep(1 * time.Second)
	coal += <-transferpoint

	time.Sleep(1 * time.Second)
	coal += <-transferpoint

	fmt.Println(coal)
	fmt.Println(time.Since(start))
}
