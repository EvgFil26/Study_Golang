package main

import "fmt"

// Создание структуры
type Apartment struct {
	Address string
	Square  float64
	Floor   int
	Price   float64
}

// Конструктор для создания новой квартиры
func NewApartment(
	address string,
	square float64,
	floor int,
	price float64,
) *Apartment {
	if address == "" || square <= 10 || (floor < 0 || floor > 100) {
		fmt.Println("Пользователь не прошёл валидацию")
		return nil
	} else {
		return &Apartment{
			Address: address,
			Square:  square,
			Floor:   floor,
			Price:   price,
		}
	}
}

// Метод изменения стоимости.
func (u *Apartment) PriceUPD(newPrice float64) {
	if newPrice > 0 {
		u.Price = newPrice
	} else {
		fmt.Println("Ошибка: стоимость должна быть больше нуля")
	}
}

func main() {
	// Использование конструктора
	u1 := NewApartment("ул. Школьная, 225", 12, 25, 900000)

	fmt.Println("До изменения:", *u1)

	// Обновление стоимости
	if u1 == nil {
		return
	}
	u1.PriceUPD(120000)

	fmt.Println("После изменения:", *u1)

}
