package errors

// {
// 	"title": "Unauthorized_error",
// 	"details": "The email 'john@doe.com' is already on the system.",
// 	"status": 411
//   }

type UnauthorizedError struct {
	Message string `json:"message" example:"missing or malformed jwt"`
}

func Unauthorized() *UnauthorizedError {
	Unauthorized := &UnauthorizedError{
		Message: "missing or malformed jwt",
	}

	return Unauthorized
}
