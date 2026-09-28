package miner

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	n int,
	power int,
	transferPoint chan<- int,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр номер", n, "и мой рабочий день завершён")
			return
		default:
			fmt.Println("Я шахтёр номер", n, ", начал добычу угля")
			time.Sleep(1 * time.Second)
			fmt.Println("Я шахтёр номер", n, ", добыл", power, "угля")

			transferPoint <- power
			fmt.Println("Я шахтёр номер", n, ", отнёс уголь на пункт приёма")
		}
	}
}

func MinerPool(
	ctx context.Context,
	minerCount int,
) <-chan int {
	CoalTranferPoint := make(chan int)

	wg := &sync.WaitGroup{}
	for i := 1; i < minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, i, rand.Intn(20), CoalTranferPoint)
	}

	go func() {
		wg.Wait()
		close(CoalTranferPoint)
	}()

	return CoalTranferPoint
}
