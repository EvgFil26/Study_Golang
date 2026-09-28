package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp/v3"
)

func Readtext() string {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Введите команду: ")
		if ok := scanner.Scan(); !ok {
			return ("Ошибка ввода")
		}

		text := scanner.Text()
		fmt.Println("Текст: ", text)
		cmd, header, textmain := parttext(text)
		fmt.Println(cmd)
		switch cmd {
		case "add":
			added := AddTask(header, textmain)
			pp.Println("Добавлен(add): ", added)
		case "list":
			all := All()
			pp.Println("Все(list): ", all)
		case "del":
			all := Del(header)
			pp.Println("Все(del): ", all)
		case "done":
			all := Complete(header)
			pp.Println("Все(complete): ", all)
		}
	}
}

func parttext(text string) (string, string, string) {
	fields := strings.Fields(text)
	pp.Println(fields)

	cmd := fields[0]
	pp.Println(cmd)
	if cmd == "list" {
		return cmd, "", ""
	}

	header := fields[1]
	pp.Println(header)

	textmain := ""
	for i := 2; i < len(fields); i++ {
		textmain += fields[i] + " "
	}
	pp.Println(textmain)
	return cmd, header, textmain
}
