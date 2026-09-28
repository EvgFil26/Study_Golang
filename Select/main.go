package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)
	go func() {
		i := 0
		for {
			intCh <- i
			i++
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		i := 0
		for {
			strCh <- "hi" + strconv.Itoa(i)
			i++
			time.Sleep(200 * time.Millisecond)
		}
	}()
	// Селект сам выбирает что делать в зависимости от очереди, а не ждёт долгий intCh
	for {
		select {
		case number := <-intCh:
			fmt.Println(number)
		case str := <-strCh:
			fmt.Println(str)
		}
		// можно удалить бесконечность и дописать default:
		// если в момент прихода программы case ещё не готовы

	}
}
