package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"
	"todo-app/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// In-memory storage for todos (replace with database in production)
var todos = make(map[string]models.Todo)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	// Convert map to slice for sorting
	todoList := make([]models.Todo, 0, len(todos))
	for _, todo := range todos {
		todoList = append(todoList, todo)
	}

	// Sort todos: non-completed first, then by creation date
	sort.Slice(todoList, func(i, j int) bool {
		// If completion status is different, non-completed comes first
		if todoList[i].Completed != todoList[j].Completed {
			return !todoList[i].Completed
		}
		// If completion status is the same, sort by creation date (older first)
		return todoList[i].CreatedAt.Before(todoList[j].CreatedAt)
	})

	log.Printf("GetTodos: Returning %d todos: %+v\n", len(todoList), todoList)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoList)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todo, exists := todos[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todoCreate models.TodoCreate
	if err := json.NewDecoder(r.Body).Decode(&todoCreate); err != nil {
		log.Printf("CreateTodo: Error decoding request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	log.Printf("CreateTodo: Received request with data: %+v\n", todoCreate)
	now := time.Now()
	todo := models.Todo{
		ID:          uuid.New().String(),
		Title:       todoCreate.Title,
		Description: todoCreate.Description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	todos[todo.ID] = todo

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todo, exists := todos[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		return
	}

	var todoUpdate models.TodoUpdate
	if err := json.NewDecoder(r.Body).Decode(&todoUpdate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	if todoUpdate.Title != nil {
		todo.Title = *todoUpdate.Title
	}
	if todoUpdate.Description != nil {
		todo.Description = *todoUpdate.Description
	}
	if todoUpdate.Completed != nil {
		todo.Completed = *todoUpdate.Completed
	}

	todo.UpdatedAt = time.Now()
	todos[id] = todo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, exists := todos[id]; !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		return
	}

	delete(todos, id)
	w.WriteHeader(http.StatusNoContent)
}
