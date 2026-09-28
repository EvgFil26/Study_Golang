package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	score := 0

	for {
		fmt.Println("-----------------------------------")
		fmt.Println("Текущий счёт:", score)

		fmt.Println("Подлетаю к трубе")
		fmt.Println("100")

		fmt.Println("Пролетаю через трубу")
		fmt.Println("010")

		if rand.Intn(4) == 2 {
			fmt.Println("Упс! Я врезался в трубу, очков:", score)
			break
		}

		fmt.Println("Пролетел через трубу")
		fmt.Println("001")

		score++
		time.Sleep(500 * time.Millisecond)
	}
}
