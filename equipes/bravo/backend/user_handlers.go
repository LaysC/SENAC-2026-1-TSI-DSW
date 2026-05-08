package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// POST /api/v1/users
func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "corpo da requisição inválido")
		return
	}

	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "o campo 'name' é obrigatório")
		return
	}
	if req.Email == "" {
		writeError(w, http.StatusBadRequest, "o campo 'email' é obrigatório")
		return
	}
	if req.Password == "" {
		writeError(w, http.StatusBadRequest, "o campo 'password' é obrigatório")
		return
	}

	result, err := app.db.Exec(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		req.Name, req.Email, req.Password,
	)
	if err != nil {
		// email duplicado
		if isDuplicateEntry(err) {
			writeError(w, http.StatusConflict, "email já cadastrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "erro ao criar usuário")
		return
	}

	id, _ := result.LastInsertId()

	var user User
	err = app.db.QueryRow(
		"SELECT user_id, name, email, created_at, updated_at FROM users WHERE user_id = ?", id,
	).Scan(&user.UserID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "erro ao buscar usuário criado")
		return
	}

	writeJSON(w, http.StatusCreated, user)
}

// GET /api/v1/users
func (app *App) listUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := app.db.Query(
		"SELECT user_id, name, email, created_at, updated_at FROM users ORDER BY created_at DESC",
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "erro ao buscar usuários")
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.UserID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
			writeError(w, http.StatusInternalServerError, "erro ao ler usuário")
			return
		}
		users = append(users, u)
	}

	writeJSON(w, http.StatusOK, users)
}

// GET /api/v1/users/{id}
func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
	idStr := extractIDFromPath(r.URL.Path)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id inválido")
		return
	}

	var user User
	err = app.db.QueryRow(
		"SELECT user_id, name, email, created_at, updated_at FROM users WHERE user_id = ?", id,
	).Scan(&user.UserID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, "usuário não encontrado")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "erro ao buscar usuário")
		return
	}

	writeJSON(w, http.StatusOK, user)
}

// DELETE /api/v1/users/{id}
func (app *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := extractIDFromPath(r.URL.Path)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id inválido")
		return
	}

	result, err := app.db.Exec("DELETE FROM users WHERE user_id = ?", id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "erro ao deletar usuário")
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		writeError(w, http.StatusNotFound, "usuário não encontrado")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}