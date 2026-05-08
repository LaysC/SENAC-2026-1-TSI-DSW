package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// extrai o último segmento da URL: /api/v1/tasks/42 → "42"
func extractIDFromPath(path string) string {
	parts := strings.Split(strings.TrimSuffix(path, "/"), "/")
	return parts[len(parts)-1]
}

// verifica se o erro é de entrada duplicada no MySQL (ex: email único)
func isDuplicateEntry(err error) bool {
	var mysqlErr *mysql.MySQLError
	return strings.Contains(err.Error(), "Duplicate entry") ||
		(mysqlErr != nil && mysqlErr.Number == 1062)
}
