package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/abhi00999/task-management/internal/handler"
	"github.com/abhi00999/task-management/pkg/db"
)

func main() {
	db.InitMongoDB()
	r := mux.NewRouter()

	h := handler.NewTaskHandler()

	r.HandleFunc("/health", h.Health).Methods("GET")

	r.HandleFunc("/tasks", h.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", h.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", h.UpdateTask).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}