package services

import (
	"fmt"
	"library-management/models"
)


type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

type LibraryManager interface {
	AddBook(book models.Book)
	AddMember(member models.Member)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	ListMembers() []models.Member
}

/* A function to add new book*/
func (lib *Library) AddBook(book models.Book){
	lib.Books[book.ID] = book
	fmt.Println("\n\tBook Added Successfully.")
}

/* A function to delete a book*/
func (lib *Library) RemoveBook(bookID int){
	book, exists := lib.Books[bookID]
	if !exists{
		fmt.Println("\n\tBook not available.")
		return
	}
	if book.Status == "Borrowed"{
		fmt.Println("\n\tYou can not delete a borrowed book.")
		return
	}
	delete(lib.Books, bookID)
	fmt.Println("\n\tBook Deleted Successfully.")
}

/* A function to handle book borrowing*/
func (lib *Library)BorrowBook(bookID int, memberID int) error{
	book, exists := lib.Books[bookID]
	if !exists{
		return fmt.Errorf("\n\tbook Not found")
	}
	if book.Status == "Borrowed"{
		return fmt.Errorf("\n\tbook already borrowed")
	}
	member, exist := lib.Members[memberID]
	if !exist{
		return fmt.Errorf("\n\tno member with this Id")
	}
	book.Status = "Borrowed"
	lib.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.Members[memberID] = member
	fmt.Println("\n\tBook borrowed Successfully")
	return nil
}

/* A function to handle book returns*/
func (lib *Library)ReturnBook(bookID int, memberID int) error{
	book, exists := lib.Books[bookID]
	if !exists{
		return fmt.Errorf("\n\tbook is not available")
	}
	if book.Status == "Available"{
		return fmt.Errorf("\n\tbook is not borrowed")
	}
	member, exist := lib.Members[memberID]
	if !exist{
		return fmt.Errorf("\n\tno member with this Id")
	}
	flag := true
	borrowedBook := member.BorrowedBooks
	for _, b := range borrowedBook{
		if b.ID == bookID{
			flag = false
			break
		}
	}
	if flag{
		return fmt.Errorf("\n\tmember doesn't borrow a book with this Id")
	}
	book.Status = "Available"
	lib.Books[bookID] = book
	borrowedBooks := member.BorrowedBooks
	var availableBooks []models.Book
	for _, v := range borrowedBooks{
		if v.ID == bookID{
			continue
		}
		availableBooks = append(availableBooks, v)
	}
	member.BorrowedBooks = availableBooks
	lib.Members[memberID] = member
	fmt.Print("\n\tBook Returned Successfully.")
	return nil
}

/* A function to list all available books*/
func (lib *Library) ListAvailableBooks() []models.Book{
	var books []models.Book
	allBooks := lib.Books
	for _, bookData := range allBooks{
		books = append(books, bookData)
	}
	return books
}

/* A function to list all available books*/
func (lib *Library) ListMembers() []models.Member{
	var members []models.Member
	allMembers := lib.Members
	for _, member := range allMembers{
		members = append(members, member)
	}
	return members
}

/* A function to list borrowed books by a user*/
func (lib *Library) ListBorrowedBooks(memberID int) []models.Book{
	member, exists := lib.Members[memberID]
	if !exists{
		return nil
	}
	return member.BorrowedBooks
}

/* A function to add a library member*/
func (lib *Library) AddMember(member models.Member){
	lib.Members[member.ID] = member
	fmt.Println("\n\tMember Added Successfully")
}