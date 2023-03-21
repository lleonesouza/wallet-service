package errors

import (
	"fmt"
	"net/http"
)

// {
// 	"title": "conflict_error",
// 	"details": "The email 'john@doe.com' is already on the system.",
// 	"status": 411
//   }

type ConflictError struct {
	Title   string `json:"title" example:"conflict_error"`
	Details string `json:"details" example:"The email 'john@doe.com' is already on the system."`
	Status  int    `json:"status" example:"409"`
}

func Conflict(key string, value string) *ConflictError {
	Conflict := &ConflictError{
		Title:   "conflict_error",
		Status:  http.StatusConflict,
		Details: fmt.Sprintf("The %s '%s' is already on the system.", key, value),
	}

	return Conflict
}
