# Library Management System

## Overview
This is a simple console-based library management system implemented in Go. It has the functionalities of adding, removing, borrowing, and returning books. The system also keeps track of the books borrowed by each library member. The application also provides a simple console interface to interact with the library management system. 

## Structs
- **Book**: Represents a book with fields for ID, title, author, and status.
- **Member**: Represents a library member with fields for ID, name, and borrowed books.

## Interfaces
- **LibraryManager**: Defines the methods for library management including adding, removing, borrowing, and returning books.

## Implementation
- **Library**: Implements the LibraryManager interface and contains the logic for managing books and members.

## Console Interface
The application provides a simple console interface to interact with the library management system. Users can add, remove, borrow, and return books, as well as list available and borrowed books. The application also provides an option to list all library members and their borrowed books. It has also an option to add and remove library members.

## Running the Application
To run the application, use the following command:

```sh
go run main.go
