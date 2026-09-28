package main

import (
	"fmt"

	"todolistsimple/tasks"
)

func main() {
	text := tasks.Readtext()
	fmt.Println(text)
}
