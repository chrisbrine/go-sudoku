package server

import (
	"net/http"

	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func HealthCheck(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	RespondSuccess(w, true)
}

func AddMiscMethods(db *sql.DBData) {
	HandleGET(db, "/api/health", HealthCheck, false)
}