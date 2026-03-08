package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("===ЗАКЛАДКА===")
	bookMark := make(map[string]string)
menu:
	for {
		fmt.Println("\nМЕНЮ:")
		fmt.Println("1.Посмтреть закладки")
		fmt.Println("2.Добавить закладку")
		fmt.Println("3.Удалить закладку")
		fmt.Println("4.Выход")

		choice := readInput("Введите пункт меню: ")
		fmt.Println("Ваш выбор: ", choice)
		switch choice {
		case "1":
			viewBookmark(bookMark)
		case "2":
			addBookMark(bookMark)
		case "3":
			deleteBookmark(bookMark)
		case "4":
			fmt.Println("Вы нажали выход ")
			break menu
		default:
			fmt.Println("Неверный пункт меню ")
		}

	}

}

func deleteBookmark(bookMark map[string]string) {

	name := readInput("Введите закладку которую надо удалить: ")
	_, ok := bookMark[name]
	if ok {
		fmt.Printf("Закладка - '%s: %s' удалена\n", name, bookMark[name])
		delete(bookMark, name)
	} else {
		fmt.Println("Такой закладки нет ")
	}

}
func readInput(promt string) string {
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
func addBookMark(bookMark map[string]string) {

	nameBook := readInput("Введите название заклдаки: ")

	urlBook := readInput("Введите URL заклдаки: ")

	bookMark[nameBook] = urlBook
	fmt.Printf("Закладка '%s:%s' добавлена\n", nameBook, urlBook)

}

func viewBookmark(bookMark map[string]string) {
	if len(bookMark) == 0 {
		fmt.Println("Пока нет закладок ...")
		return
	}

	fmt.Println("\nВаши закладки:")

	for i, v := range bookMark {
		fmt.Println(i, ":", v)
	}
}
