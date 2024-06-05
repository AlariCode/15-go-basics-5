package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Новый проект")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)
}
