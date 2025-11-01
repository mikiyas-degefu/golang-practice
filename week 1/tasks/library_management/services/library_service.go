package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	if book.Status == "" {
		book.Status = "Available"
	}
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("cannot remove a borrowed book")
	}
	delete(l.books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}
	member, ok := l.members[memberID]
	if !ok {
		return errors.New("member not found")
	}
	// mark book as borrowed and save it back to the map
	book.Status = "Borrowed"
	l.books[bookID] = book
	// add a copy of the book to the member's BorrowedBooks
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	// save the updated member back to the map
	l.members[memberID] = member
	return nil
}

// ReturnBook allows a member to return a borrowed book.
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	member, ok := l.members[memberID]
	if !ok {
		return errors.New("member not found")
	}
	// find the book in member's borrowed slice
	index := -1
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("this member did not borrow that book")
	}
	// remove the book from the slice (simple method)
	member.BorrowedBooks = append(member.BorrowedBooks[:index], member.BorrowedBooks[index+1:]...)
	// update the member back into the map
	l.members[memberID] = member
	// mark book as available and update map
	book.Status = "Available"
	l.books[bookID] = book
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	out := []models.Book{}
	for _, b := range l.books {
		if b.Status == "Available" {
			out = append(out, b)
		}
	}
	return out
}

// ListBorrowedBooks returns the books borrowed by a specific member.
func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, ok := l.members[memberID]
	if !ok {
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}

// AddMember is a helper to add a member into the library.
// It returns an error if an ID conflict exists.
func (l *Library) AddMember(member models.Member) error {
	if _, exists := l.members[member.ID]; exists {
		return errors.New("member ID already exists")
	}
	l.members[member.ID] = member
	return nil
}

// GetAllMembers returns a slice of all members (simple helper used by the controller).
func (l *Library) GetAllMembers() []models.Member {
	out := []models.Member{}
	for _, m := range l.members {
		out = append(out, m)
	}
	return out
}
