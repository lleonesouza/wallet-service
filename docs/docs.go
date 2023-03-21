// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "leone de souza",
            "url": "https://github.com/lleonesouza",
            "email": "lleonesouza@live.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/shopkeeper": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get account information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shopkeeper"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ShopkeeperResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.UnauthorizedError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update 'Name' and/or 'Lastname' of Shopkeeper account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shopkeeper"
                ],
                "parameters": [
                    {
                        "description": "Shopkeeper",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateShopkeeperDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ShopkeeperResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.UnauthorizedError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.UnprocessableEntityError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a Shopkeeper account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shopkeeper"
                ],
                "parameters": [
                    {
                        "description": "Create Shopkeeper Account Input",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateShopkeeperDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.ShopkeeperResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.BadRequestError"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/errors.ConflictError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.UnprocessableEntityError"
                        }
                    }
                }
            }
        },
        "/shopkeeper/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shopkeeper"
                ],
                "parameters": [
                    {
                        "description": "Shopkeeper",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginShopkeeperDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.UnauthorizedError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.UnprocessableEntityError"
                        }
                    }
                }
            }
        },
        "/transaction": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get transaction from Wallet.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ResponseTransactionDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.UnauthorizedError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a transaction.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "parameters": [
                    {
                        "description": "Create Transaction Input",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateTransactionDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.ResponseTransactionDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.UnauthorizedError"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get account information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserResponseDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.UnauthorizedError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update 'Name' and/or 'Lastname' of User account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "User",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.InternalServerError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.UnauthorizedError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.UnprocessableEntityError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a User account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "Create User Account Input",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserResponseDTO"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/errors.ConflictError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.UnprocessableEntityError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.InternalServerError"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "User",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginResponseDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.UnauthorizedError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.UnprocessableEntityError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.InternalServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CreateShopkeeperDTO": {
            "type": "object",
            "required": [
                "cnpj",
                "email",
                "lastname",
                "name",
                "password"
            ],
            "properties": {
                "cnpj": {
                    "type": "string",
                    "example": "12345678"
                },
                "email": {
                    "type": "string",
                    "example": "john@doe.com"
                },
                "lastname": {
                    "type": "string",
                    "example": "doe"
                },
                "name": {
                    "type": "string",
                    "example": "john"
                },
                "password": {
                    "type": "string",
                    "example": "12345678"
                }
            }
        },
        "dtos.CreateTransactionDTO": {
            "type": "object",
            "required": [
                "to",
                "value"
            ],
            "properties": {
                "to": {
                    "type": "string",
                    "example": "john@doe.com"
                },
                "value": {
                    "type": "integer",
                    "example": 20
                }
            }
        },
        "dtos.CreateUserDTO": {
            "type": "object",
            "required": [
                "cpf",
                "email",
                "lastname",
                "name",
                "password"
            ],
            "properties": {
                "cpf": {
                    "type": "string",
                    "example": "42971056830"
                },
                "email": {
                    "type": "string",
                    "example": "john@doe.com"
                },
                "lastname": {
                    "type": "string",
                    "maxLength": 19,
                    "minLength": 3,
                    "example": "doe"
                },
                "name": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3,
                    "example": "john"
                },
                "password": {
                    "type": "string",
                    "example": "JohnDoe!@#0"
                }
            }
        },
        "dtos.LoginResponseDTO": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5AZG9lLmNvbSIsImlkIjoiMWE4MjQwM2YtYWNhOS00YjA1LTliNTEtYjRmZWE4OGM2MWQ5IiwidHlwZSI6InNob3BrZWVwZXIiLCJleHAiOjE2NzU1NDkyODd9.MSgwyCzvbC6tfH7ZYNrEhhv_XbmKqVEX-rEe6Y7EMKI"
                }
            }
        },
        "dtos.LoginShopkeeperDTO": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john@doe.com"
                },
                "password": {
                    "type": "string",
                    "example": "12345678"
                }
            }
        },
        "dtos.LoginUserDTO": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john@doe.com"
                },
                "password": {
                    "type": "string",
                    "example": "JohnDoe!@#0"
                }
            }
        },
        "dtos.ResponseTransactionDTO": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string",
                    "example": "2023-01-31 12:47:27.072 +0000 UTC"
                },
                "from_user_id": {
                    "type": "string",
                    "example": "06901d3b-134b-4ea6-ba0f-3a00ca5836b7"
                },
                "id": {
                    "type": "string",
                    "example": "06901d3b-134b-4ea6-ba0f-3a00ca5836b7"
                },
                "to_user_id": {
                    "type": "string",
                    "example": "06901d3b-134b-4ea6-ba0f-3a00ca5836b7"
                },
                "update_at": {
                    "type": "string",
                    "example": "2023-01-31 12:47:27.072 +0000 UTC"
                },
                "value": {
                    "type": "integer",
                    "example": 50
                }
            }
        },
        "dtos.ShopkeeperResponseDTO": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer",
                    "example": 50
                },
                "cnpj": {
                    "type": "string",
                    "example": "12345789"
                },
                "create_at": {
                    "type": "string",
                    "example": "2023-01-31 12:47:27.072 +0000 UTC"
                },
                "email": {
                    "type": "string",
                    "example": "jhon@doe.com"
                },
                "id": {
                    "type": "string",
                    "example": "06901d3b-134b-4ea6-ba0f-3a00ca5836b7"
                },
                "lastname": {
                    "type": "string",
                    "example": "doe"
                },
                "name": {
                    "type": "string",
                    "example": "john"
                },
                "update_at": {
                    "type": "string",
                    "example": "2023-01-31 12:47:27.072 +0000 UTC"
                }
            }
        },
        "dtos.UpdateShopkeeperDTO": {
            "type": "object",
            "required": [
                "lastname",
                "name"
            ],
            "properties": {
                "lastname": {
                    "type": "string",
                    "example": "doe"
                },
                "name": {
                    "type": "string",
                    "example": "john"
                }
            }
        },
        "dtos.UpdateUserDTO": {
            "type": "object",
            "required": [
                "lastname",
                "name"
            ],
            "properties": {
                "lastname": {
                    "type": "string",
                    "maxLength": 19,
                    "minLength": 3,
                    "example": "doe"
                },
                "name": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3,
                    "example": "john"
                }
            }
        },
        "dtos.UserResponseDTO": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer",
                    "example": 50
                },
                "cpf": {
                    "type": "string",
                    "example": "12345678"
                },
                "create_at": {
                    "type": "string",
                    "example": "2023-01-31 12:47:27.072 +0000 UTC"
                },
                "email": {
                    "type": "string",
                    "example": "john@doe.com"
                },
                "id": {
                    "type": "string",
                    "example": "06901d3b-134b-4ea6-ba0f-3a00ca5836b7"
                },
                "lastname": {
                    "type": "string",
                    "example": "doe"
                },
                "name": {
                    "type": "string",
                    "example": "john"
                },
                "update_at": {
                    "type": "string",
                    "example": "2023-01-31 12:47:27.072 +0000 UTC"
                },
                "wallet_id": {
                    "type": "string",
                    "example": "06901d3b-134b-4ea6-ba0f-3a00ca5836b7"
                }
            }
        },
        "errors.BadRequestError": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "The email 'john@doe.com' is already on the system."
                },
                "status": {
                    "type": "integer",
                    "example": 409
                },
                "title": {
                    "type": "string",
                    "example": "BadRequest_error"
                }
            }
        },
        "errors.ConflictError": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "The email 'john@doe.com' is already on the system."
                },
                "status": {
                    "type": "integer",
                    "example": 409
                },
                "title": {
                    "type": "string",
                    "example": "conflict_error"
                }
            }
        },
        "errors.InternalServerError": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "Something in our services is not right. We are working on it. Please try again."
                },
                "status": {
                    "type": "integer",
                    "example": 500
                },
                "title": {
                    "type": "string",
                    "example": "InternalServer_error"
                }
            }
        },
        "errors.UnauthorizedError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "missing or malformed jwt"
                }
            }
        },
        "errors.UnprocessableEntityError": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Field validation for 'Password' failed on the 'password' tag",
                        "Field validation for 'CPF' failed on the 'cpf' tag"
                    ]
                },
                "status": {
                    "type": "integer",
                    "example": 422
                },
                "title": {
                    "type": "string",
                    "example": "body_error"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Token used authenticate 'User' and 'Shopkeeper'",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "wallet-service API",
	Description:      "A service to exchange coins between Shopkeeper wallets and Users Wallets",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}