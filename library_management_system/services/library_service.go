package services

import (
	"errors"
	"library_management_system/models"
)

// LibraryManager defines the methods for library management
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book

	AddMember(member models.Member)
	RemoveMember(memberID int)
}

// Library implements the LibraryManager interface
type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() Library{
	return Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (lib *Library) AddMember(member models.Member) {
	lib.Members[member.ID] = member
}

func (lib *Library) RemoveMember(memberID int) {
	if _, exists := lib.Members[memberID]; !exists {
		return
	}

	delete(lib.Members, memberID)
}

func (lib *Library) AddBook(book models.Book) {
	lib.Books[book.ID] = book
}

func (lib *Library) RemoveBook(bookID int) {
	delete(lib.Books, bookID)
}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExists := lib.Books[bookID]
	member, memberExists := lib.Members[memberID]

	if !bookExists {
		return errors.New("book not found")
	}
	if !memberExists {
		return errors.New("member not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	book.Status = "Borrowed"

	member.BorrowedBooks[bookID] = book

	lib.Books[bookID] = book

	return nil
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := lib.Books[bookID]
	member, memberExists := lib.Members[memberID]

	if !bookExists {
		return errors.New("book not found")
	}
	if !memberExists {
		return errors.New("member not found")
	}
	if book.Status == "Available" {
		return errors.New("book was not borrowed")
	}

	book.Status = "Available"
	lib.Books[bookID] = book
	delete(member.BorrowedBooks, bookID)

	return nil
}

func (lib *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range lib.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := lib.Members[memberID]
	if !exists {
		return nil
	}

	var borrowedBooks []models.Book
	for _, book := range member.BorrowedBooks {
		borrowedBooks = append(borrowedBooks, book)
	}
	return borrowedBooks
}
