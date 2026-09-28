package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type PaymentInfo struct {
	Description string `json:"Description"`
	USD         int    `json:"USD"`
	FullName    string `json:"FullName"`
	Adress      string `json:"Adress"`
	Time        time.Time
}

var mtx = sync.Mutex{}
var money = 1000
var PaymentHistory = make([]PaymentInfo, 0)

type InfoOut struct {
	Money          int
	PaymentHistory []PaymentInfo
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment PaymentInfo

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		msg := "Ошибка чтения запроса"
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}

	mtx.Lock()
	if money-payment.USD > 0 {
		money -= payment.USD
	}

	payment.Time = time.Now()
	PaymentHistory = append(PaymentHistory, payment)
	fmt.Println(money)
	fmt.Println(PaymentHistory)
	b := InfoOut{
		Money:          money,
		PaymentHistory: PaymentHistory,
	}
	a, _ := json.Marshal(b)
	w.Write([]byte(a))
	mtx.Unlock()

	/*if err := json.Unmarshal(httpRequestBody, &payment); err != nil {
		msg := "Ошибка парсинга запроса"
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}*/

}

func main() {
	http.HandleFunc("/pay", payHandler)
	if err := http.ListenAndServe(":9092", nil); err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}
