package main

import (
	"log"

	users "trylibraries/Users"

	"github.com/k0kubun/pp/v3"
)

func main() {
	pp.Println("Про")

	a, err := users.CreateUser()
	if err != nil {
		log.Fatal(err)
	}

	a = users.UpdatePrice(a)
	pp.Println(a)
}
