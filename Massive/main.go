package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	userArray := []User{
		{
			"Толя",
			6.5,
			true,
		},
		{
			"Ваня",
			8.5,
			true,
		},
		{
			"Вася",
			6.5,
			false,
		},
		{
			"Петя",
			3.1,
			false,
		},
		{
			"Игорь",
			6.1,
			false,
		},
	}
	fmt.Println("________ДО____________")
	pp.Println(userArray)

	for i := 0; i < len(userArray); i++ {
		if userArray[i].Premium {
			userArray[i].Rating += 1
		}
	}

	/*
		for index, user := range userArray {
			if user.Premium {
				userArray[index].Rating += 1
			}
		}

		fmt.Println("________ПОСЛЕ_________")
		for index := range userArray {
			pp.Println(userArray[index])
		}
	*/

	fmt.Println("________ПОСЛЕ_________")
	pp.Println(userArray)

	fmt.Println(
		len(userArray),
		cap(userArray),
	)

	userArray = append(userArray, User{
		"Виталик",
		4.1,
		true,
	})

	fmt.Println("________ПОСЛЕ_append_________")
	pp.Println(userArray)

	fmt.Println(
		len(userArray),
		cap(userArray),
	)

	fmt.Println("________Создание пустого массива слайса")

	intSlice := make([]int, 0 /* длина массива */, 5 /*начальная способность вместить в массив*/)

	fmt.Println(intSlice,
		len(intSlice),
		cap(intSlice),
	)

	intSlice = append(intSlice, 1, 5, 3)

	fmt.Println(intSlice,
		len(intSlice),
		cap(intSlice),
	)
}
