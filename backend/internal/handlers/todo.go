package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"todo-app/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// In-memory storage for todos (replace with database in production)
var todos = make(map[string]models.Todo)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todoList := make([]models.Todo, 0, len(todos))
	for _, todo := range todos {
		todoList = append(todoList, todo)
	}
	
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}
	
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
