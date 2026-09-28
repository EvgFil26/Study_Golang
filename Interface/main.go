package main

import (
	"interface/payments"
	"interface/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewBank()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Бургер", 5)
	idphone := paymentModule.Pay("Телефон", 200)
	idticket := paymentModule.Pay("Билет", 55)

	paymentModule.Cancel(idphone)

	allInfo := paymentModule.AllInfo()
	pp.Println("Все оплаты:", allInfo)

	ticketinfo := paymentModule.Info(idticket)
	pp.Println("ticket info:", ticketinfo)
}
