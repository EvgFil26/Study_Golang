package main

import "fmt"

func main() {
	number := 20
	fmt.Println("До функций: ", number)
	pointer := &number
	fmt.Println(pointer)

	foo(number)
	fmt.Println("После функции числа: ", *pointer)
	boo(pointer)
	fmt.Println("После функции указателя: ", *pointer)
}

func foo(n int) {
	n = 10
	fmt.Println(n)
}

func boo(n *int) {
	*n = 10
	fmt.Println(*n)
}
