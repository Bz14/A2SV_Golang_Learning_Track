package controllers

import (
	"bufio"
	"fmt"
	"library-management/models"
	"library-management/services"
	"os"
	"strconv"
	"strings"
)

var bookIdValue = 1
var memberIdValue = 1

/* A function to add a new book*/
func AddNewBook(reader *bufio.Reader, library services.LibraryManager){
	flag := true
	for flag{
		invalidTitle := true
		invalidAuthor := true
		var title, author string
		for invalidTitle{
			fmt.Printf("\tEnter Book Title: ")
			title, _ = reader.ReadString('\n')
			if strings.TrimSpace(title) == ""{
				fmt.Println("\tInvalid input. Try again.")
				continue
			}
			invalidTitle = false
		}
		for invalidAuthor{
			fmt.Printf("\tEnter Book Author: ")
			author, _ = reader.ReadString('\n')
			if strings.TrimSpace(author) == ""{
				fmt.Println("\tInvalid input. Try again.")
				continue
			}
			invalidAuthor = false
		}
		newBook := models.Book{
			ID: int(bookIdValue), 
			Title: strings.TrimSpace(title), 
			Author: strings.TrimSpace(author), 
			Status: "Available",
		}
		library.AddBook(newBook)
		bookIdValue += 1
		flag = false
	}
}

/* A function to add a new member*/
func AddMember(reader *bufio.Reader, library services.LibraryManager){
	flag := true
	for flag{
		invalidMember := true
		var name string
		if invalidMember{
			fmt.Printf("\tEnter Member Name: ")
			name, _ = reader.ReadString('\n')
			if strings.TrimSpace(name) == ""{
				fmt.Println("\tInvalid input. Try again.")
				continue
			}
			invalidMember = false
		}
		newMember := models.Member{
			ID: int(memberIdValue),
			Name: strings.TrimSpace(name),
		}
		library.AddMember(newMember)
		memberIdValue += 1
		flag = false
	}
}

/* A function to add remove a book*/
func DeleteBook(reader *bufio.Reader, library services.LibraryManager){
	flag := true
	for flag{
		fmt.Printf("\tEnter Book Id to delete: ")
		id, _ := reader.ReadString('\n')
		bookId, err := strconv.ParseInt(strings.TrimSpace(id), 10, 36)
		if err != nil{
			fmt.Println("\tBook Id must be integer. Please insert an integer.")
			continue
		}
		library.RemoveBook(int(bookId))
		flag = false	
	}
}

/* A function to borrow book*/
func BorrowAvailableBook(reader *bufio.Reader, library services.LibraryManager){
	flag := true
	for flag{
		fmt.Printf("\tEnter Book Id to borrow: ")
		id1, _ := reader.ReadString('\n')
		bookId, err := strconv.ParseInt(strings.TrimSpace(id1), 10, 36)
		if err != nil{
			fmt.Println("\tBook Id must be integer. Please insert an integer.")
			continue
		}
		fmt.Printf("\tEnter Member Id: ")
		id2 , _ := reader.ReadString('\n')
		memberId, err := strconv.ParseInt(strings.TrimSpace(id2), 10, 36)
		if err != nil{
			fmt.Println("\tMember Id must be integer. Please insert an integer.")
			continue
		}
		err = library.BorrowBook(int(bookId), int(memberId))
		if err != nil{
			fmt.Println(err)
		}
		flag = false
	}
}

/* A function to return a book*/
func ReturnBorrowedBook(reader *bufio.Reader, library services.LibraryManager){
	flag := true
	for flag{
		fmt.Printf("\tEnter Book Id to return: ")
		id1, _ := reader.ReadString('\n')
		bookId, err := strconv.ParseInt(strings.TrimSpace(id1), 10, 36)
		if err != nil{
			fmt.Println("\tBook Id must be integer. Please insert an integer.")
			continue
		}
		fmt.Printf("\tEnter Member Id: ")
		id2 , _ := reader.ReadString('\n')
		memberId, err := strconv.ParseInt(strings.TrimSpace(id2), 10, 36)
		if err != nil{
			fmt.Println("\tMember Id must be integer. Please insert an integer.")
			continue
		}
		err = library.ReturnBook(int(bookId), int(memberId))
		if err != nil{
			fmt.Println(err)
		}
		flag = false
	}
}

/* A function to display all books*/
func ShowBooks(reader *bufio.Reader, library services.LibraryManager){
	allBooks := library.ListAvailableBooks()
	if len(allBooks) < 1{
		fmt.Println("\n\tNo Available Books.")
		return
	}
	fmt.Println("\t----------------------------------------------------------------------")
	fmt.Println("\t\tBookId\t\tTitle\t\tAuthor\t\tStatus")
	fmt.Println("\t----------------------------------------------------------------------")
	for _, book := range allBooks{
		fmt.Printf("\t\t%d\t\t%s\t\t%s\t\t%s", book.ID, book.Title,book.Author,book.Status)
		fmt.Println("\n\t----------------------------------------------------------------------")
	}
}

/* A function to display all memebers*/
func ShowMembers(reader *bufio.Reader, library services.LibraryManager){
	allMembers := library.ListMembers()
	if len(allMembers) < 1{
		fmt.Println("\n\tNo Available Members.")
		return
	}
	fmt.Println("\t-----------------------------------------------------------")
	fmt.Println("\t\tMemberId\t\tName")
	fmt.Println("\t------------------------------------------------------------")
	for _, member := range allMembers{
		fmt.Printf("\t\t%d\t\t\t%s", member.ID, member.Name)
		fmt.Println("\n\t-----------------------------------------------------------")
	}
}


/* A function to display borrowed books*/
func ShowBorrowedBooks(reader *bufio.Reader, library services.LibraryManager){
	flag := true
	for flag{
		fmt.Printf("\tEnter Member Id: ")
		id2 , _ := reader.ReadString('\n')
		memberId, err := strconv.ParseInt(strings.TrimSpace(id2), 10, 36)
		if err != nil{
			fmt.Println("\tMember Id must be integer. Please insert an integer.")
			continue
		}
		borrowedBooks := library.ListBorrowedBooks(int(memberId))
		if borrowedBooks == nil{
			fmt.Println("\tMember not found.")
			return
		}
		if len(borrowedBooks) < 1{
			fmt.Println("\tNo available books.")
			return
		}
		fmt.Println("\t----------------------------------------------------------------------")
		fmt.Println("\t\tBookId\t\tTitle\t\tAuthor\t\tStatus")
		fmt.Println("\t----------------------------------------------------------------------")
		for _, book := range borrowedBooks{
			fmt.Printf("\t\t%d\t\t%s\t\t%s\t\t%s", book.ID, book.Title,book.Author,book.Status)
			fmt.Println("\n\t----------------------------------------------------------------------")
		}
		flag = false
	}
	
}

/* A function to show options*/
func ShowOptions(){
	options := "\n\t=========================================================\n"
	options += "\t\t\t1.Add new book.\n\t\t\t2.Add new Member.\n\t\t\t3.Remove a book.\n\t\t\t4.Borrow a book.\n\t\t\t"
	options += "5.Return a book.\n\t\t\t6.List all books.\n\t\t\t7.List all borrowed books by a member.\n\t\t\t8.Show all members\n\t\t\t9.Exit."
	options += "\n\t========================================================="
	fmt.Println(options)
}

func Menu() {
	reader := bufio.NewReader(os.Stdin)
	newLibrary := services.Library{
		Books: map[int]models.Book{},
		Members: map[int]models.Member{},
	}
	var library services.LibraryManager  = &newLibrary
	flag := true
	for flag{
		ShowOptions()
		fmt.Print("\tEnter Your choice: ")
		choice, _ := reader.ReadString('\n')
		fmt.Println("\n\t=========================================================")
		choice = strings.TrimSpace(choice)
		c, _ := strconv.ParseInt(choice, 10, 35)
		switch c{
		case 1:
			AddNewBook(reader, library)
		case 2:
			AddMember(reader, library)
		case 3:
			DeleteBook(reader, library)
		case 4:
			BorrowAvailableBook(reader, library)
		case 5:
			ReturnBorrowedBook(reader, library)
		case 6:
			ShowBooks(reader, library)
		case 7:
			ShowBorrowedBooks(reader, library)
		case 8:
			ShowMembers(reader, library)
		case 9:
			flag = false
		default:
			fmt.Println("\tInvalid input.Try Again.")
		}
	}
}