package main

import "time"

type Task struct {
	TaskID    int       `json:"task_id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// payload para criação de usuário
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// payload para criação de tarefa
type CreateTaskRequest struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

// payload para atualização de tarefa
type UpdateTaskRequest struct {
	Status string `json:"status"` // "pending" ou "completed"
}