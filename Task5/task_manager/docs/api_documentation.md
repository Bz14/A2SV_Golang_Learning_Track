# Task Management API

An api endpoint for performing CRUD operations.

## Project Structure

The project is organized into several packages:

- **main.go**: This is the entry point of the application.
- **models/**: This package contains the data structures used in the application, including `Task`.
- **controllers/**: This package contains the `task_controller.go` file, which Handles incoming HTTP requests and invokes the appropriate service methods.
- **router/**: This package contains the `router.go` file, which sets up the routes and initializes the Gin router and Defines the routing configuration for the API
- **data/**: This package contains the `task_service.go` file, which contains business logic and data manipulation functions.

## Installation Instructions

1. Clone the repository:
   ```bash
   git clone <https://github.com/Bz14/A2SV_Golang_Learning_Track.git>
   ```
2. cd Task4
3. cd task_manager
4. Install MongoDB (if you haven't already)

Follow the instructions for your operating system on the [official MongoDB installation page](https://docs.mongodb.com/manual/installation/).

5. Start the MongoDB server

Ensure your MongoDB server is running. You can start it using:

```bash
  mongod
```

4. go run main.go

## Usage Instructions

- Start the application:

```bash
  go run main.go
```

## Endpoints

### GET /tasks

**Description:** Get a list of all tasks.

**Response:**

- **200 OK:** Returns the list of tasks.

---

### GET /tasks/:id

**Description:** Get the details of a specific task.

**Parameters:**

- `id` (path): The ID of the task.

**Response:**

- **200 OK:** Returns the task details.
- **400 Bad Request:** Invalid task ID.
- **404 Not Found:** Task not found.

---

### POST /tasks

**Description:** Create a new task.

**Request Body:**

- `title` (string): The title of the task.
- `description` (string): The description of the task.
- `due_date` (string): The due date of the task.
- `status` (string): The status of the task.

**Response:**

- **201 Created:** Returns the created task.
- **400 Bad Request:** Invalid input.

---

### PUT /tasks/:id

**Description:** Update a specific task.

**Parameters:**

- `id` (path): The ID of the task.

**Request Body:**

- `title` (string): The new title of the task.
- `description` (string): The new description of the task.
- `due_date` (string): The new due date of the task.
- `status` (string): The new status of the task.

**Response:**

- **200 OK:** Returns the updated task.
- **400 Bad Request:** Invalid task ID or input.
- **404 Not Found:** Task not found.

---

### DELETE /tasks/:id

**Description:** Delete a specific task.

**Parameters:**

- `id` (path): The ID of the task.

**Response:**

- **400 Bad Request:** Invalid task ID.
- **404 Not Found:** Task not found.

[Postman Documentation](https://documenter.getpostman.com/view/34226868/2sA3rwKtQ1)
