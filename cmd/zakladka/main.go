package main

import (
	"fmt"
	"zakladka/internal/bookmark"
	"zakladka/internal/cli"
)

func main() {
	fmt.Println("===ЗАКЛАДКА===")

	menuOptions := map[string]string{
		"1": "Посмотреть закладки",
		"2": "Добавить закладку",
		"3": "Удалить закладку",
		"4": "Выход",
	}

	commands := map[string]func(){
		"1": func() { bookmark.ListBookmarks() },
		"2": func() {
			name := cli.ReadInput("Введите название заклдаки: ")
			url := cli.ReadInput("Введите URL заклдаки: ")
			bookmark.AddBookmark(name, url)
		},
		"3": func() {
			name := cli.ReadInput("Введите название заклдаки: ")
			bookmark.DeleteBookmark(name)
		},
	}

	for {

		choice := cli.Menu(menuOptions)
		fmt.Println("Ваш выбор: ", choice)

		if choice == "4" {
			fmt.Println("Выход из программы ")
			return
		}

		command, ok := commands[choice]
		if ok {
			command()
		} else {
			fmt.Println("Неверный пункт меню")
		}

	}

}
