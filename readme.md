# Go Todo API

This is a simple REST API for managing a todo list, built with Go and the Gin framework.

## Features

- **Get all todos:** Retrieve the list of all todos.
- **Get a specific todo by ID:** Retrieve the details of a specific todo item by its ID.
- **Add a new todo:** Add a new todo item to the list.
- **Update todo status:** Update the completion status of a specific todo item.

## Endpoints

### Get All Todos

- **URL:** `/todos`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200
  - **Content:** `[{"Id": "1", "item": "clean room", "completed": false}, ...]`

### Get Todo by ID

- **URL:** `/todo/:id`
- **Method:** `GET`
- **URL Params:** `id=[string]`
- **Success Response:**
  - **Code:** 200
  - **Content:** `{"Id": "1", "item": "clean room", "completed": false}`
- **Error Response:**
  - **Code:** 404
  - **Content:** `{"message": "Todo not found"}`

### Add a New Todo

- **URL:** `/todos`
- **Method:** `POST`
- **Data Params:**
  - **Content:** `{"Id": "4", "item": "new todo item", "completed": false}`
- **Success Response:**
  - **Code:** 201
  - **Content:** `{"Id": "4", "item": "new todo item", "completed": false}`

### Update Todo Status

- **URL:** `/todo/:id`
- **Method:** `PATCH`
- **URL Params:** `id=[string]`
- **Success Response:**
  - **Code:** 200
  - **Content:** `{"Id": "1", "item": "clean room", "completed": true}`
- **Error Response:**
  - **Code:** 404
  - **Content:** `{"message": "Todo not found"}`

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/go-todo-api.git
