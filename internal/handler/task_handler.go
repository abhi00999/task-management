package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/abhi00999/task-management/internal/service"
	"github.com/abhi00999/task-management/models"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{service: service.NewTaskService()}
}

func (h *TaskHandler) Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("healhty")
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	ctx := context.TODO()
	t, err := h.service.Create(ctx, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(t)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	skip, _ := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)
	ctx := context.TODO()
	tasks, err := h.service.List(ctx, status, limit, skip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	ctx := context.TODO()
	err = h.service.Update(ctx, id, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	ctx := context.TODO()
	err = h.service.Delete(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}