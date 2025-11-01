package models

// BorrowedBooks holds the books this member has borrowed
type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}
