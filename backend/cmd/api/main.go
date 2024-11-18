package main

import (
	"log"
	"net/http"
	"todo-app/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// CORS middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	})

	// Routes
	r.HandleFunc("/api/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/api/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/api/todos/{id}", handlers.GetTodo).Methods("GET")
	r.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	// Start server
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
