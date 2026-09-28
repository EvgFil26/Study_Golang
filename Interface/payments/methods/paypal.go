package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func (c PayPal) Pay(usd int) int {
	fmt.Println("Оплата через PayPal")
	fmt.Println("Размер оплаты:", usd, "долларов")

	return rand.Int()
}

func NewPayPal() PayPal {
	return PayPal{}
}

func (c PayPal) Cancel(id int) {
	fmt.Println("Отмена PayPal-операции. ID:", id)
}
