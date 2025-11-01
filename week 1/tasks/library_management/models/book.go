package models

// Status should be either "Available" or "Borrowed".
type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}
