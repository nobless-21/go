package main

import (
	"testing"
)

func TestAddBook(t *testing.T) {
	facade := NewLibraryFacade()
	facade.bookManeger.Add("Go Programming")

	if _, exists := book["Go Programming"]; !exists {
		t.Errorf("Expected book 'Go Programming' to be added, but it was not found.")
	}
}

func TestSearchBook(t *testing.T) {
	facade := NewLibraryFacade()
	facade.bookManeger.Add("Go Programming")

	result := facade.bookManeger.Search("Go Programming")
	if result != 1 {
		t.Errorf("Expected to find the book, but got result %d", result)
	}

	result = facade.bookManeger.Search("Nonexistent Book")
	if result != 0 {
		t.Errorf("Expected to not find the book, but got result %d", result)
	}
}

func TestDeleteBook(t *testing.T) {
	facade := NewLibraryFacade()
	facade.bookManeger.Add("Go Programming")
	facade.bookManeger.Delete("Go Programming")

	result := facade.bookManeger.Search("Go Programming")
	if result != 0 {
		t.Errorf("Expected book 'Go Programming' to be deleted, but it was found.")
	}
}

func TestAddUser(t *testing.T) {
	facade := NewLibraryFacade()
	facade.readerManager.Add("Alice")

	if _, exists := user["Alice"]; !exists {
		t.Errorf("Expected user 'Alice' to be added, but it was not found.")
	}
}

func TestSearchUser(t *testing.T) {
	facade := NewLibraryFacade()
	facade.readerManager.Add("Alice")

	result := facade.readerManager.Search("Alice")
	if result != 1 {
		t.Errorf("Expected to find the user, but got result %d", result)
	}

	result = facade.readerManager.Search("Bob")
	if result != 0 {
		t.Errorf("Expected to not find the user, but got result %d", result)
	}
}

func TestDeleteUser(t *testing.T) {
	facade := NewLibraryFacade()
	facade.readerManager.Add("Alice")
	facade.readerManager.Delete("Alice")

	result := facade.readerManager.Search("Alice")
	if result != 0 {
		t.Errorf("Expected user 'Alice' to be deleted, but it was found.")
	}
}

func TestLoanBook(t *testing.T) {
	facade := NewLibraryFacade()
	facade.bookManeger.Add("Go Programming")
	facade.loanManager.Delete("Go Programming") // выдача книги

	result := facade.loanManager.Search("Go Programming")
	if result != 2 {
		t.Errorf("Expected book 'Go Programming' to be loaned out, but got result %d", result)
	}

	facade.loanManager.Add("Go Programming") // возврат книги
	result = facade.loanManager.Search("Go Programming")
	if result != 1 {
		t.Errorf("Expected book 'Go Programming' to be available, but got result %d", result)
	}
}