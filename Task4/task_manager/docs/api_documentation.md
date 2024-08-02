# Task Management API Documentation

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
