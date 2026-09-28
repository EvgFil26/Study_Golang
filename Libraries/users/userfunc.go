package users

import "fmt"

func newApartment(
	address string,
	square float64,
	floor int,
	price float64,
) (*Apartment, error) {
	if address == "" || square <= 10 || floor < 0 || floor > 100 {
		return nil, fmt.Errorf("не прошёл валидацию")
	}
	return &Apartment{
		Address: address,
		Square:  square,
		Floor:   floor,
		Price:   price,
	}, nil
}

func (u *Apartment) priceUPD(newPrice float64) {
	if newPrice > 0 {
		u.Price = newPrice
	} else {
		fmt.Println("Ошибка: стоимость должна быть больше нуля")
	}
}
