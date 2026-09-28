package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	n int,
	mail string,
	transferPoint chan<- string,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я почтальон номер", n, "и мой рабочий день завершён")
			return
		default:
			fmt.Println("Я почтальон номер", n, ", получил письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Я почтальон номер", n, ", донёс письмо до почты:", mail)

			transferPoint <- mail
			fmt.Println("Я почтальон номер", n, ", передал письмо", mail)
		}
	}
}

func PostmanPool(
	ctx context.Context,
	postmanCount int,
) <-chan string {
	MailTranferPoint := make(chan string)

	wg := &sync.WaitGroup{}
	for i := 1; i < postmanCount; i++ {
		wg.Add(1)
		go postman(ctx, wg, i, postmanToMail(i), MailTranferPoint)
	}

	go func() {
		wg.Wait()
		close(MailTranferPoint)
	}()

	return MailTranferPoint
}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Привет от семьи",
		2: "Рабочее уведомление",
		3: "Письмо с учёбы",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Лотерея"
	}
	return mail
}
