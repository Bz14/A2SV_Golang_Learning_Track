# Library Management System

A console application for managing library operations, including adding, borrowing, returning of a book and also members.

## Project Structure

The project is organized into several packages:

- **main.go**: This is the entry point of the application.
- **models/**: This package contains the data structures used in the application, including `Book` and `Member`.
- **controllers/**: This package contains the `library_controller.go` file, which handles user input and controls the flow of the application.
- **services/**: This package contains the `library_service.go` file, which provides the core functionality of the library management system, such as adding and removing books and members.

## Installation Instructions

1. Clone the repository:
   ```bash
   git clone <https://github.com/Bz14/A2SV_Golang_Learning_Track.git>
   ```
2. cd Task3
3. cd library-management
4. go run main.go

## Usage Instructions

- Start the application:
  ```bash
  go run main.go
  ```

## Features

- Registering a new book.
- Registering a new member.
- Removing a book.
- Track book loans and returns.
- Display all available books.
- Display a book borrowed by a member.
- Display all members
