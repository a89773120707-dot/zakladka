package cli

import (
	"bufio"
	"fmt"
	"os"
)

func Menu(options map[string]string) string {
	fmt.Println("\nМЕНЮ:")

	for key, desc := range options {
		fmt.Printf("%s. %s\n", key, desc)
	}
	return ReadInput("Введите пункт меню: ")
}
func ReadInput(promt string) string {
	sacnner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(promt)
		sacnner.Scan()
		text := sacnner.Text()
		if text == "" {
			fmt.Println("Поле не может быть пустым, попробуйте еще раз")
			continue
		}
		return text
	}
}
