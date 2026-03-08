package bookmark

import "fmt"

var bookMarks = make(map[string]string)

func AddBookmark(name string, url string) {
	bookMarks[name] = url
	fmt.Printf("Закладка '%s: %s' добавлена\n", name, url)
}

func DeleteBookmark(name string) {
	_, ok := bookMarks[name]
	if ok {
		fmt.Printf("Закладка - '%s: %s' удалена\n", name, bookMarks[name])
		delete(bookMarks, name)
	} else {
		fmt.Println("Такой закладки нет")
	}
}

func ListBookmarks() {
	if len(bookMarks) == 0 {
		fmt.Println("Пока нет закладок ")
	}
	for name, url := range bookMarks {
		fmt.Println(name, ":", url)
	}
}
