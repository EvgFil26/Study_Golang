package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp/v3"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите команду")
	for {
		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода")
			return
		}

		text := scanner.Text()
		fmt.Println("Текст:", text)

		fields := strings.Fields(text)
		pp.Println(fields)

		cmd := fields[0]

		if cmd == "выйти" {
			return
		}

		fmt.Println(cmd)
	}
}
