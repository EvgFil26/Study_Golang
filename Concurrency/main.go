package main

import (
	"Concurrency/miner"
	"Concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	t := time.Now()
	mtx := sync.Mutex{}
	var coal atomic.Int64
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("Рабочий день шахтёров завершён")
		minerCancel()
	}()

	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("Рабочий день почтальонов завершён")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 3)
	mailTransferPoint := postman.PostmanPool(postmanContext, 3)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := range coalTransferPoint {
			coal.Add(int64(c))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for mail := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, mail)
			mtx.Unlock()
		}
	}()

	wg.Wait()
	mtx.Lock()
	fmt.Println(len(mails), coal.Load(), time.Since(t))
	mtx.Unlock()
}
