# Go Todo App

A simple command-line todo application written in Go that helps you manage your tasks efficiently.

## Features

- Add new todos
- Edit existing todos
- Mark todos as complete/incomplete
- Delete todos
- List all todos in a formatted table
- Persistent storage using JSON

## Installation

Make sure you have Go installed on your system, then clone and build the application:

```bash
git clone <repository-url>
cd go-todo-app
go mod tidy
go build -o todo-app
```

## Usage

The application uses command-line flags to perform different operations:

### Add a new todo
```bash
./todo-app -add "Your todo title here"
```

### List all todos
```bash
./todo-app -list
```

### Edit a todo
```bash
./todo-app -edit "0:New title for todo"
```
Replace `0` with the index of the todo you want to edit.

### Toggle todo completion status
```bash
./todo-app -toggle 0
```
Replace `0` with the index of the todo you want to toggle.

### Delete a todo
```bash
./todo-app -del 0
```
Replace `0` with the index of the todo you want to delete.

## Data Storage

Todos are automatically saved to and loaded from `todos.json` in the application directory. The data persists between application runs.

## Requirements

- Go 1.24.4 or later
- Dependencies are managed via `go.mod`
