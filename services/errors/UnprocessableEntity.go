package errors

import (
	"net/http"
	"strings"
)

// {
// 	"title": "body_error",
// 	"details": [
// 	  "Field validation for 'Password' failed on the 'password' tag",
// 	  "Field validation for 'CPF' failed on the 'cpf' tag"
// 	],
// 	"status": 422
//   }

type UnprocessableEntityError struct {
	Title   string   `json:"title" example:"body_error"`
	Details []string `json:"details" example:"Field validation for 'Password' failed on the 'password' tag,Field validation for 'CPF' failed on the 'cpf' tag"`
	Status  int      `json:"status" example:"422"`
}

func UnprocessableEntity(errors string) *UnprocessableEntityError {
	unprocessableEntity := &UnprocessableEntityError{
		Title:  "body_error",
		Status: http.StatusUnprocessableEntity,
	}

	details := strings.Split(errors, "\n")
	for _, detail := range details {
		_, after, _ := strings.Cut(detail, "Error:")
		unprocessableEntity.Details = append(unprocessableEntity.Details, after)
	}

	return unprocessableEntity
}
