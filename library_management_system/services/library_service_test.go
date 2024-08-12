package services

import (
	"library_management_system/models"
	"testing"
)

func TestAddBook(t *testing.T) {
	lib := NewLibrary()
	book := models.Book{ID: 1, Title: "Go Programming", Author: "John Doe", Status: "Available"}

	lib.AddBook(book)

	if _, exists := lib.Books[book.ID]; !exists {
		t.Errorf("Expected book with ID %d to be added, but it wasn't", book.ID)
	}
}

func TestRemoveBook(t *testing.T) {
	lib := NewLibrary()
	book := models.Book{ID: 1, Title: "Go Programming", Author: "John Doe", Status: "Available"}

	lib.AddBook(book)
	lib.RemoveBook(book.ID)

	if _, exists := lib.Books[book.ID]; exists {
		t.Errorf("Expected book with ID %d to be removed, but it wasn't", book.ID)
	}
}

func TestBorrowBook(t *testing.T) {
	lib := NewLibrary()
	book := models.Book{ID: 1, Title: "Go Programming", Author: "John Doe", Status: "Available"}
	member := models.Member{ID: 1, Name: "Alice", BorrowedBooks: map[int]models.Book{}}

	lib.AddBook(book)
	lib.Members[member.ID] = member

	err := lib.BorrowBook(book.ID, member.ID)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if lib.Books[book.ID].Status != "Borrowed" {
		t.Errorf("Expected book status to be 'Borrowed', but got %s", lib.Books[book.ID].Status)
	}

	if len(lib.Members[member.ID].BorrowedBooks) != 1 {
		t.Errorf("Expected member to have 1 borrowed book, but got %d", len(lib.Members[member.ID].BorrowedBooks))
	}
}

func TestReturnBook(t *testing.T) {
	lib := NewLibrary()
	book := models.Book{ID: 1, Title: "Go Programming", Author: "John Doe", Status: "Available"}
	member := models.Member{ID: 1, Name: "Alice", BorrowedBooks: map[int]models.Book{}}

	lib.AddBook(book)
	lib.Members[member.ID] = member

	lib.BorrowBook(book.ID, member.ID)
	err := lib.ReturnBook(book.ID, member.ID)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if lib.Books[book.ID].Status != "Available" {
		t.Errorf("Expected book status to be 'Available', but got %s", lib.Books[book.ID].Status)
	}

	if len(lib.Members[member.ID].BorrowedBooks) != 0 {
		t.Errorf("Expected member to have 0 borrowed books, but got %d", len(lib.Members[member.ID].BorrowedBooks))
	}
}

func TestListAvailableBooks(t *testing.T) {
	lib := NewLibrary()
	book1 := models.Book{ID: 1, Title: "Go Programming", Author: "John Doe", Status: "Available"}
	book2 := models.Book{ID: 2, Title: "Go Advanced", Author: "Jane Smith", Status: "Borrowed"}

	lib.AddBook(book1)
	lib.AddBook(book2)

	availableBooks := lib.ListAvailableBooks()
	if len(availableBooks) != 1 {
		t.Errorf("Expected 1 available book, but got %d", len(availableBooks))
	}

	if availableBooks[0].ID != book1.ID {
		t.Errorf("Expected available book ID to be %d, but got %d", book1.ID, availableBooks[0].ID)
	}
}

func TestListBorrowedBooks(t *testing.T) {
	lib := NewLibrary()
	book := models.Book{ID: 1, Title: "Go Programming", Author: "John Doe", Status: "Available"}
	member := models.Member{ID: 1, Name: "Alice", BorrowedBooks: map[int]models.Book{}}

	lib.AddBook(book)
	lib.AddMember(member)

	lib.BorrowBook(book.ID, member.ID)

	borrowedBooks := lib.ListBorrowedBooks(member.ID)

	if len(borrowedBooks) != 1 {
		t.Errorf("Expected 1 borrowed book, but got %d", len(borrowedBooks))
	}

	if borrowedBooks[0].ID != book.ID {
		t.Errorf("Expected borrowed book ID to be %d, but got %d", book.ID, borrowedBooks[0].ID)
	}
}
