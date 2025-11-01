package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

func StartConsole() {
	lib := services.NewLibrary()
	seedSampleData(lib)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		fmt.Print("Enter choice: ")

		if !scanner.Scan() {
			fmt.Println("Goodbye!")
			return
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			handleAddBook(scanner, lib)
		case "2":
			handleRemoveBook(scanner, lib)
		case "3":
			handleBorrowBook(scanner, lib)
		case "4":
			handleReturnBook(scanner, lib)
		case "5":
			handleListAvailable(lib)
		case "6":
			handleListBorrowed(scanner, lib)
		case "7":
			handleListMembers(lib)
		case "0":
			fmt.Println("Exiting. Bye!")
			return
		default:
			fmt.Println("Unknown choice, please try again.")
		}
		fmt.Println()
	}
}

func printMenu() {
	fmt.Println("=== Simple Library ===")
	fmt.Println("1) Add book")
	fmt.Println("2) Remove book")
	fmt.Println("3) Borrow book")
	fmt.Println("4) Return book")
	fmt.Println("5) List available books")
	fmt.Println("6) List borrowed books (by member)")
	fmt.Println("7) List members")
	fmt.Println("0) Exit")
}

// Helpers for handling each menu action follow.
func handleAddBook(scanner *bufio.Scanner, lib *services.Library) {
	fmt.Print("Enter book ID (number): ")
	id := readInt(scanner)
	fmt.Print("Enter title: ")
	title := readLine(scanner)
	fmt.Print("Enter author: ")
	author := readLine(scanner)
	book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
	lib.AddBook(book)
	fmt.Println("Book added.")
}

func handleRemoveBook(scanner *bufio.Scanner, lib *services.Library) {
	fmt.Print("Enter book ID to remove: ")
	id := readInt(scanner)
	if err := lib.RemoveBook(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Book removed.")
}

func handleBorrowBook(scanner *bufio.Scanner, lib *services.Library) {
	fmt.Print("Enter member ID: ")
	mid := readInt(scanner)
	fmt.Print("Enter book ID to borrow: ")
	bid := readInt(scanner)
	if err := lib.BorrowBook(bid, mid); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Book borrowed.")
}

func handleReturnBook(scanner *bufio.Scanner, lib *services.Library) {
	fmt.Print("Enter member ID: ")
	mid := readInt(scanner)
	fmt.Print("Enter book ID to return: ")
	bid := readInt(scanner)
	if err := lib.ReturnBook(bid, mid); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Book returned.")
}

func handleListAvailable(lib *services.Library) {
	books := lib.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}
	fmt.Println("Available books:")
	for _, b := range books {
		fmt.Printf("ID: %d | %s — %s", b.ID, b.Title, b.Author)
	}
}

func handleListBorrowed(scanner *bufio.Scanner, lib *services.Library) {
	fmt.Print("Enter member ID: ")
	mid := readInt(scanner)
	books, err := lib.ListBorrowedBooks(mid)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("This member has not borrowed any books.")
		return
	}

	fmt.Printf("Books borrowed by member %d:", mid)

	for _, b := range books {
		fmt.Printf("ID: %d | %s — %s", b.ID, b.Title, b.Author)
	}
}

func handleListMembers(lib *services.Library) {
	members := lib.GetAllMembers()
	if len(members) == 0 {
		fmt.Println("No members found.")
		return
	}
	fmt.Println("Members:")
	for _, m := range members {
		fmt.Printf("ID: %d | %s | Borrowed: %d", m.ID, m.Name, len(m.BorrowedBooks))
	}
}

func readLine(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		return ""
	}
	return strings.TrimSpace(scanner.Text())
}

// readInt reads an integer from the user. It loops until a valid integer is entered.
func readInt(scanner *bufio.Scanner) int {
	for {
		if !scanner.Scan() {
			return 0
		}
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			fmt.Print("Please enter a number: ")
			continue
		}
		i, err := strconv.Atoi(text)
		if err != nil {
			fmt.Print("Invalid number, try again: ")
			continue
		}
		return i
	}
}

// seedSampleData adds a few books and members so the user can try commands immediately.
func seedSampleData(lib *services.Library) {
	lib.AddBook(models.Book{ID: 1, Title: "The Go Programming Language", Author: "Yohannes Belete"})
	lib.AddBook(models.Book{ID: 2, Title: "Clean Code", Author: "Abel Berhe"})
	lib.AddMember(models.Member{ID: 1, Name: "Mikiyas"})
	lib.AddMember(models.Member{ID: 2, Name: "Kaleab"})
}
