package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// nil канал - плохая идея, лучше его инициализировать через make
	// close(ch) - при записи и при двойном закрытии паника, читается из него default value (проверить через ok можно)
	ch := make(chan int)

	go func() {
		rand := rand.IntN(4)
		iterations := 3 + rand
		for i := 1; i < iterations; i++ {
			ch <- 10
		}
		close(ch)
	}()

	coal := 0
	for {
		value, ok := <-ch
		if !ok {
			break
		}
		coal += value
		fmt.Println(coal)
	}
}
