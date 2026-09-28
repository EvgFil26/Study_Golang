package main

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/k0kubun/pp/v3"
)

type Car struct {
	armour int
}

func (c *Car) Gas() (int, int, error) {
	speed := rand.IntN(150)
	if speed > 100 {
		c.armour -= 10
	}
	if c.armour <= 0 {
		err := errors.New("Нельзя так быстро ездить, машина уже сломалась")
		return 0, 0, err
	}
	return speed, c.armour, nil
}

func main() {
	/* defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была паника,", p)
		}
	}()
	a := 0
	b := 1 / a
	fmt.Println(b) */
	car := Car{
		armour: 100,
	}
	for {
		pp.Println("Машина до поездки:", car)
		speed, armour, err := car.Gas()
		pp.Println("Машина после поездки:", car)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println("Разгон:", speed, "Состояние:", armour)
	}
}
