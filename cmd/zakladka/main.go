package main

import (
	"fmt"
	"zakladka/internal/bookmark"
	"zakladka/internal/cli"
)

func main() {
	fmt.Println("===ЗАКЛАДКА===")

	service := bookmark.NewService()
	menuOptions := map[string]string{
		"1": "Посмотреть закладки",
		"2": "Добавить закладку",
		"3": "Удалить закладку",
		"4": "Выход",
	}

	commands := map[string]func(){
		"1": func() {
			bookmarks := service.ListBookmarks()
			if len(bookmarks) == 0 {
				fmt.Println("Пока нет закладок ")
				return
			}
			for name, url := range bookmarks {
				fmt.Println(name, ":", url)
			}

		},
		"2": func() {
			name := cli.ReadInput("Введите название заклдаки: ")
			url := cli.ReadInput("Введите URL заклдаки: ")
			service.AddBookmark(name, url)
			fmt.Printf("Закладка '%s: %s' добавлена\n", name, url)
		},
		"3": func() {
			name := cli.ReadInput("Введите название заклдаки: ")
			url, deleted := service.DeleteBookmark(name)
			if deleted {
				fmt.Printf("Закладка - '%s: %s' удалена\n", name, url)
			} else {
				fmt.Println("Такой закладки нет")
			}
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
