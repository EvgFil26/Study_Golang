package main

import "fmt"

func main() {
	criminal := make(map[string]bool, 10)
	criminal["Никита"] = true
	criminal["Артём"] = false
	fmt.Println(criminal)

	c, ok := criminal["Виталий"]
	if !ok {
		fmt.Println("Человека нет в базе")
		return
	}
	if ok {
		fmt.Println("Человек найден в базе")
	}
	if c {
		fmt.Println("Человек судим")
	} else {
		fmt.Println("Человек НЕ судим")
	}

}

/* func main() {
/*  weather := make(map[int]int, 10) можно предвыделить кол-во значений */
/* weather := map[int]int{
		1: -5,
		2: -3,
		3: 1,
		4: -2,
	}

	fmt.Println(weather)

	for k, v := range weather {
		v += 1
		fmt.Println(k, v)
	}

	fmt.Println(weather) // После этого цикла мапа не поменяется!!! (обращение к копии)

	for k, v := range weather {
		weather[k] += 1
		fmt.Println(k, v)
	}

	fmt.Println(weather) // После этого цикла мапа меняется!!! (обращение по индексу)

	c, ok := weather[1] // weather[11] (по ключу)
	if ok {
		fmt.Println("Okay", c)
	} else {
		fmt.Println("Ошибка")
	}
} */
