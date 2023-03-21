package errors

import (
	"net/http"
)

// {
// 	"title": "InternalServer_error",
// 	"details": "The email 'john@doe.com' is already on the system.",
// 	"status": 411
//   }

type InternalServerError struct {
	Title   string `json:"title" example:"InternalServer_error"`
	Details string `json:"details" example:"Something in our services is not right. We are working on it. Please try again."`
	Status  int    `json:"status" example:"500"`
}

func InternalServer() *InternalServerError {
	InternalServer := &InternalServerError{
		Title:   "internal_server_error",
		Status:  http.StatusInternalServerError,
		Details: "Something in our services is not right. We are working on it. Please try again.",
	}

	return InternalServer
}
