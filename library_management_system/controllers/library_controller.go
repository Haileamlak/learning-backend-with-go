package controllers

import (
	"fmt"
	"library_management_system/models"
	"library_management_system/services"
	"os"
	"os/exec"
)

func ShowMenu() {
	fmt.Println("Library Management System")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")
	fmt.Println("7. Exit")
}

func HandleConsoleInput(lib services.LibraryManager) {
	for {
		clearConsole()
		ShowMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Println("Enter book ID:")
			fmt.Scan(&id)
			fmt.Println("Enter book title:")
			fmt.Scan(&title)
			fmt.Println("Enter book author:")
			fmt.Scan(&author)
			lib.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "Available"})
			fmt.Printf("Book with id: %v, title: %v, and author: %v added successfully.", id, title, author)

			wait()

		case 2:
			var id int
			fmt.Println("Enter book ID to remove:")
			fmt.Scan(&id)
			lib.RemoveBook(id)
			fmt.Printf("Book with id: %v removed successfully.", id)

			wait()

		case 3:
			var bookID, memberID int
			fmt.Println("Enter book ID to borrow:")
			fmt.Scan(&bookID)
			fmt.Println("Enter member ID:")
			fmt.Scan(&memberID)
			err := lib.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Book with book id: %v has been borrowed to member with member id: %v successfully.", bookID, memberID)
			}
			wait()
		case 4:
			var bookID, memberID int
			fmt.Println("Enter book ID to return:")
			fmt.Scan(&bookID)
			fmt.Println("Enter member ID:")
			fmt.Scan(&memberID)
			err := lib.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Book with book id: %v has been returned from member with member id: %v successfully.", memberID, bookID)
			}
			wait()

		case 5:
			books := lib.ListAvailableBooks()
			if len(books) == 0 {
				fmt.Println("No available books.")

			} else {

				fmt.Println("Available Books:")
				for _, book := range books {
					fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
				}
			}

			wait()
		case 6:
			var memberID int
			fmt.Println("Enter member ID to list borrowed books:")
			fmt.Scan(&memberID)
			books := lib.ListBorrowedBooks(memberID)

			if len(books) == 0 {
				fmt.Println("No borrowed books by this member.")

			} else {

				fmt.Println("Borrowed Books:")
				for _, book := range books {
					fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
				}
			}

			wait()
		case 7:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func wait() {
	fmt.Println()
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	fmt.Scanln()
}
