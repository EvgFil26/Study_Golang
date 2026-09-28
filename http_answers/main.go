package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var money int
var bank int
var mtx = sync.Mutex{}

func main() {
	money = (1000)
	http.HandleFunc("/pay", payhandler)
	http.HandleFunc("/save", savehandler)

	http.ListenAndServe(":9091", nil)
}

func payhandler(w http.ResponseWriter, r *http.Request) {
	HttpBodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Ошибка чтения запроса")
	}

	HttpBodyRequestString := string(HttpBodyRequest)

	payamount, err := strconv.Atoi(HttpBodyRequestString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Ошибка конвертации запроса")
		return
	}

	mtx.Lock()
	if money >= payamount {
		time.Sleep(3 * time.Second)
		money -= payamount
		fmt.Println("Оплата прошла успешно", money)
		strconv.Itoa(money)
		w.Write([]byte("Оплата прошла успешно " + strconv.Itoa(money)))
	} else {
		fmt.Println("На балансе не хватает денег")
		w.Write([]byte("На балансе не хватает денег"))
	}
	mtx.Unlock()
}

func savehandler(w http.ResponseWriter, r *http.Request) {
	HttpBodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Ошибка чтения запроса")
	}

	HttpBodyRequestString := string(HttpBodyRequest)

	saveamount, err := strconv.Atoi(HttpBodyRequestString)
	if err != nil {
		fmt.Println("Ошибка конвертации запроса")
	}
	mtx.Lock()
	if money >= saveamount {
		bank += saveamount
		money -= saveamount
		fmt.Println("Пополнение счёта прошло успешно", money, bank)
	} else {
		fmt.Println("На балансе не хватает денег для пополнения счёта")
		return
	}
	mtx.Unlock()
}
