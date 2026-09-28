package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var money atomic.Int64
var bank atomic.Int64
var mtx = sync.Mutex{}

func main() {
	money.Add(1000)
	http.HandleFunc("/pay", payhandler)
	http.HandleFunc("/save", savehandler)

	http.ListenAndServe(":9091", nil)
}

func payhandler(w http.ResponseWriter, r *http.Request) {
	HttpBodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Ошибка чтения запроса")
	}

	HttpBodyRequestString := string(HttpBodyRequest)

	payamount, err := strconv.Atoi(HttpBodyRequestString)
	if err != nil {
		fmt.Println("Ошибка конвертации запроса")
	}
	mtx.Lock()
	if money.Load() > int64(payamount) {
		time.Sleep(3 * time.Second)
		money.Add(int64(-payamount))
	} else {
		fmt.Println("На балансе не хватает денег")
		return
	}
	mtx.Unlock()
	moneyint := int(money.Load())
	moneystring := strconv.Itoa(moneyint)
	Answer := []byte(moneystring)
	w.Write(Answer)
	fmt.Println("Оплата прошла успешно", moneyint)
}

func savehandler(w http.ResponseWriter, r *http.Request) {
	HttpBodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Ошибка чтения запроса")
	}

	HttpBodyRequestString := string(HttpBodyRequest)

	saveamount, err := strconv.Atoi(HttpBodyRequestString)
	if err != nil {
		fmt.Println("Ошибка конвертации запроса")
	}
	mtx.Lock()
	if money.Load() > int64(saveamount) {
		bank.Add(int64(saveamount))
		money.Add(-int64(saveamount))
	} else {
		fmt.Println("На балансе не хватает денег для пополнения счёта")
		return
	}
	mtx.Unlock()
	moneyint := int(money.Load())
	moneystring := strconv.Itoa(moneyint)
	Answer := []byte(moneystring)
	w.Write(Answer)
	bankint := int(bank.Load())
	bankstring := strconv.Itoa(bankint)
	Answerbank := []byte(bankstring)
	w.Write(Answerbank)
	fmt.Println("Пополнение счёта прошло успешно", moneyint, bankint)
}
