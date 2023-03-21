package handlers

import (
	"net/http"
	"wallet-service/config"
	"wallet-service/services"
	"wallet-service/services/dtos"
	"wallet-service/services/errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *services.Services
	env     *config.Envs
}

// @Description	Get account information.
// @Tags			user
// @Accept			json
// @Produce		json
// @Success		200	{object}	dtos.UserResponseDTO
// @Failure		401	{object}	errors.UnauthorizedError
// @Security		ApiKeyAuth
// @Router			/user [get]
func (u *UserHandler) Get(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)

	user, err := u.service.User.GetById(claims.ID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, errors.BadRequest())
	}

	return c.JSON(http.StatusOK, u.service.User.Filter(user))
}

// @Description	Create a User account.
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			body	body		dtos.CreateUserDTO	true	"Create User Account Input"
// @Success		201		{object}	dtos.UserResponseDTO
// @Failure		409		{object}	errors.ConflictError
// @Failure		422		{object}	errors.UnprocessableEntityError
// @Failure		500		{object}	errors.InternalServerError
// @Router			/user [post]
func (u *UserHandler) Create(c echo.Context) error {
	// Bind User
	user := new(dtos.CreateUserDTO)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServer())
	}

	// Validate body
	err = u.service.Validation.Struct(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.UnprocessableEntity(err.Error()))
	}

	// Check if email exists
	_, err = u.service.User.GetByEmail(user.Email)
	if err == nil {
		return c.JSON(http.StatusConflict, errors.Conflict("email", user.Email))
	}

	// Check if CPF exists
	_, err = u.service.User.GetByCPF(user.CPF)
	if err == nil {
		return c.JSON(http.StatusConflict, errors.Conflict("cpf", user.CPF))
	}

	// Create Wallet
	wallet, err := u.service.Wallet.Create()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServer())
	}

	// Create User
	response, err := u.service.User.Create(user, wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServer())
	}

	return c.JSON(http.StatusCreated, u.service.User.Filter(response))
}

// @Description	Update 'Name' and/or 'Lastname' of User account.
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			body	body		dtos.UpdateUserDTO	true	"User"
// @Success		200		{object}	dtos.UserResponseDTO
// @Failure		401		{object}	errors.UnauthorizedError
// @Failure		422		{object}	errors.UnprocessableEntityError
// @Failure		400		{object}	errors.InternalServerError
// @Security		ApiKeyAuth
// @Router			/user [put]
func (u *UserHandler) Update(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)

	// Bind User
	user := new(dtos.UpdateUserDTO)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServer())
	}

	// Validate Body
	err = u.service.Validation.Struct(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.UnprocessableEntity(err.Error()))
	}

	// Update
	response, err := u.service.User.Update(claims.ID, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServer())
	}

	return c.JSON(http.StatusOK, u.service.User.Filter(response))
}

// @Description	Login
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			body	body		dtos.LoginUserDTO	true	"User"
// @Success		200		{object}	dtos.LoginResponseDTO
// @Failure		401		{object}	errors.UnauthorizedError
// @Failure		500		{object}	errors.InternalServerError
// @Failure		422		{object}	errors.UnprocessableEntityError
// @Router			/user/login [post]
func (u *UserHandler) Login(c echo.Context) error {
	// Bind LoginUserDTO
	user := new(dtos.LoginUserDTO)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServer())
	}

	// Validate LoginUserDTO
	err = u.service.Validation.Struct(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.UnprocessableEntity(err.Error()))
	}

	// Get User By Email
	completeUser, err := u.service.User.GetByEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.NotFound("email", user.Email))
	}

	// Compare Password
	err = u.service.User.CheckPasswordHash(user.Password, completeUser.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, errors.Unauthorized())
	}

	// Login
	token, err := u.service.User.Login(completeUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServer())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
