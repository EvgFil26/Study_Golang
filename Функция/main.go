package main

import "fmt"

var database string = ""

func main() {
	db()
}

func db() {
	defer func() {
		fmt.Println("Конец функции db")
	}()
	if database != "" {
		fmt.Println("Database connection established")
		return
	}
	fmt.Println("Database connection failed")
}
