package main

import (
	"library_management_system/controllers"
	"library_management_system/models"
	"library_management_system/services"
)

func main() {
	lib := services.NewLibrary()

	// Adding some initial members
	lib.AddMember(models.Member{ID: 1, Name: "Tamiru"})
	lib.AddMember(models.Member{ID: 2, Name: "Haileamlak"})
	lib.AddMember(models.Member{ID: 3, Name: "Dagmawi"})
	lib.AddMember(models.Member{ID: 4, Name: "Dawit"})
	lib.AddMember(models.Member{ID: 5, Name: "Alice"})
	lib.AddMember(models.Member{ID: 6, Name: "Bob"})

	// Adding some initial books
	lib.AddBook(models.Book{ID: 1, Title: "The Hobbit", Author: "J.R.R. Tolkien", Status: "Available"})
	lib.AddBook(models.Book{ID: 2, Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Status: "Available"})
	lib.AddBook(models.Book{ID: 9, Title: "Harry Potter and the Deathly Hallows", Author: "J.K. Rowling", Status: "Available"})
	lib.AddBook(models.Book{ID: 10, Title: "The Da Vinci Code", Author: "Dan Brown", Status: "Available"})
	lib.AddBook(models.Book{ID: 11, Title: "Angels & Demons", Author: "Dan Brown", Status: "Available"})
	lib.AddBook(models.Book{ID: 12, Title: "Inferno", Author: "Dan Brown", Status: "Available"})

	
	controllers.HandleConsoleInput(lib)
}
