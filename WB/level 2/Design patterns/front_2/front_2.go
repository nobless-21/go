package main

import "fmt"

type Option interface{
	Add(string)
	Search(string) int
	Delete(string)
}

type BookManeger struct{}

var book = make(map[string]bool)

var user = make(map[string]bool)

func (b *BookManeger) Add(name string){ //добавление книги
	book[name] = true
	fmt.Println("Книга добавлена")
}

func (b *BookManeger) Search(name_book string) int{ //поиск книги
	for key, value := range book{
		if key == name_book && value{
			fmt.Println("Kнига найдена")
			return 1
		}else if key == name_book && !value{
			fmt.Println("Книга выдана")
			return 2
		} else {
			fmt.Println("Такая книга не числится в библиотеке")
			return 0
		}
	}
	return 0
}

func (b *BookManeger) Delete(name_book string){ //удаление книги
	if b.Search(name_book) == 1 {
		delete(book, name_book)
		fmt.Println("Книга удалена")
	}
}

type ReaderManager struct{}

func (r *ReaderManager) Add(name_user string){ //добавление пользователя
	user[name_user] = true
	fmt.Println("Пользователь добавлен")
}

func (r *ReaderManager) Search(name_user string) int {	//поиск пользователя
	for name := range user{
		if name_user == name {
			fmt.Println("Такой пользователь существует")
			return 1
		}else {
			fmt.Println("Такого пользователя нет")
			return 0
		}
	}
	return 0
}

func (r *ReaderManager) Delete(name_user string){
	if r.Search(name_user) == 1 {
		delete(user, name_user)
		fmt.Println("Пользователь удалён")
	}
}

type LoanManager struct{}

func (l *LoanManager) Add(name_book string) { //возврат книги
	for name := range book {
		if name == name_book {
			book[name] = true
			fmt.Println("Книгу вернули")
		}
	}
}

func (l *LoanManager) Search(name_book string) int{ //поиск книги
	for key, value := range book{
		if key == name_book && value{
			fmt.Println("Kнига найдена")
			return 1
		}else if key == name_book && !value{
			fmt.Println("Книга уже выдана")
			return 2
		} else {
			fmt.Println("Такая книга не числится в библиотеке")
			return 0
		}
	}
	return 0
}

func (l *LoanManager) Delete(name_book string) { //выдача книги
	if l.Search(name_book) == 1{
		book[name_book] = false
		fmt.Println("Книга выдана")
	}
}

type LibraryFacade struct{
	bookManeger Option
	readerManager Option
	loanManager Option
}

func NewLibraryFacade() *LibraryFacade{
	return &LibraryFacade{
		bookManeger: &BookManeger{},
		readerManager: &ReaderManager{},
		loanManager: &LoanManager{},
	}
}