# Todo App on Svelte and Go

A simple and elegant todo application built with Svelte and Vite for frontend and Go for backend.

![screenshot](screenshot.png)

## Features

- Add new todos
- Mark todos as complete
- Delete todos
- Local storage persistance
- Responsive design
- Go server with in memory storage

## Getting Started (Frontend)

### Prerequisites

- Node.js (v14 or later)
- npm

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   npm install
   ```

### Running the App

- Development mode:
  ```bash
  npm run dev
  ```

- Build for production:
  ```bash
  npm run build
  ```

## Getting Started (Backend)

### Prerequisites

- Go (v1.20 or later)

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go get github.com/gorilla/mux
   ```

### Running the App

- Development mode:
  ```bash
  go run main.go
  ```

- Build for production:
  ```bash
  go build main.go
  ```

## Project Structure

- `src/App.svelte`: Main application component
- `src/lib/TodoList.svelte`: Todo list management component
- `src/lib/TodoItem.svelte`: Individual todo item component
- `backend/cmd/api/main.go`: Go backend server
- `backend/internal/handlers/todo.go`: Endpoint
- `backend/internal/models/todo.go`: Todo model

## Technologies

- Svelte
- Vite
- JavaScript

## License

MIT License
