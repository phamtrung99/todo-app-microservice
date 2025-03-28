package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/phamtrung99/todo-app-microservice/api-gateway/client"
)

type ITodoHandler interface {
	CreateTodo(w http.ResponseWriter, r *http.Request)
	GetTodos(w http.ResponseWriter, r *http.Request)
}

type todoHandler struct {
	todoClient client.ITodoClient
}

func NewTodoHandler(todoClient client.ITodoClient) ITodoHandler {
	return &todoHandler{todoClient: todoClient}
}

func (h *todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Parsed request body: Title=%s, Description=%s", req.Title, req.Description)

	resp, err := h.todoClient.CreateTodo(req.Title, req.Description)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *todoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	resp, err := h.todoClient.GetList()
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
