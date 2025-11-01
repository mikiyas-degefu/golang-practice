# Library Management System Documentation

## Overview

This is a simple, console-based Library Management System written in Go. It demonstrates the use of structs, interfaces, slices, maps, and basic Go methods in a beginner-friendly way.

The project is designed with a clear folder structure and manual ID input for both books and members.

---

## Folder Structure

```
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   ├── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
```

---

## Project Components

### main.go

* Entry point of the application.
* Calls `controllers.StartConsole()` to start the interactive console.

### models/

* `book.go`: Defines the `Book` struct with ID, Title, Author, and Status.
* `member.go`: Defines the `Member` struct with ID, Name, and BorrowedBooks slice.

### services/

* `library_service.go`: Contains the `LibraryManager` interface and its implementation in the `Library` struct.
* Methods include AddBook, RemoveBook, BorrowBook, ReturnBook, ListAvailableBooks, ListBorrowedBooks.
* Also includes helper methods to add members and list all members.

### controllers/

* `library_controller.go`: Handles console input and interacts with the `Library` service.
* Provides menu-driven interface to:

  * Add a new book
  * Remove a book
  * Borrow a book
  * Return a book
  * List available books
  * List borrowed books by a member
  * List all members

### docs/

* `documentation.md`: Contains this documentation.

### go.mod

* Defines the module path and Go version.

---

## How to Run

1. Open terminal and navigate to the `library_management` folder.
2. Initialize Go module if not already done:

```
go mod init library_management
```

3. Tidy dependencies:

```
go mod tidy
```

4. Run the application:

```
go run main.go
```

5. Follow the menu prompts to manage books and members.

---

## Features

* Add and remove books
* Borrow and return books
* List available books
* List books borrowed by a specific member
* List all members
* Manual input for IDs (beginner-friendly)
* Clear, simple console interface

---

## Error Handling

* Borrowing a book that is already borrowed
* Returning a book not borrowed by the member
* Removing a borrowed book
* Using invalid member or book IDs

All error messages are displayed in the console for user guidance.

---

## Notes for Learners

* The system uses maps to store books and members for quick access.
* Borrowed books are stored as copies in each member's BorrowedBooks slice.
* The authoritative book status is kept in the Library's books map.
* The project structure allows beginners to understand Go concepts step-by-step.
