package main

import (
	"log"
	"net/http"
	"todo-app/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/api/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/api/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/api/todos/{id}", handlers.GetTodo).Methods("GET")
	r.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	// Create a CORS handler
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:5174", "http://127.0.0.1:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:           true, // Enable for debugging, remove in production
	})

	// Wrap router with CORS handler
	handler := corsHandler.Handler(r)

	// Start server
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
