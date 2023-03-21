package errors

import (
	"net/http"
)

// {
// 	"title": "BadRequest_error",
// 	"details": "The email 'john@doe.com' is already on the system.",
// 	"status": 411
//   }

type BadRequestError struct {
	Title   string `json:"title" example:"BadRequest_error"`
	Details string `json:"details" example:"The email 'john@doe.com' is already on the system."`
	Status  int    `json:"status" example:"409"`
}

func BadRequest() *BadRequestError {
	BadRequest := &BadRequestError{
		Title:   "BadRequest_error",
		Status:  http.StatusBadRequest,
		Details: "BadRequest",
	}

	return BadRequest
}
