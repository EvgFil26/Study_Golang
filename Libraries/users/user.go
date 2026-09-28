package users

import "fmt"

type Apartment struct {
	Address string
	Square  float64
	Floor   int
	Price   float64
}

func CreateUser() (*Apartment, error) {
	a, err := newApartment("ул. Школьная, 225", 12, 25, 900000)
	if err != nil {
		return nil, err
	}
	fmt.Println("До изменения:", *a)
	return a, nil
}

func UpdatePrice(a *Apartment) *Apartment {
	a.priceUPD(20000)
	fmt.Println("После изменения:", *a)
	return a
}
