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

type ShopkeeperHandler struct {
	service *services.Services
	env     *config.Envs
}

// @Description	Get account information.
// @Tags			shopkeeper
// @Accept			json
// @Produce		json
// @Success		200	{object}	dtos.ShopkeeperResponseDTO
// @Failure		400	{object}	errors.BadRequestError
// @Failure		401	{object}	errors.UnauthorizedError
// @Security		ApiKeyAuth
// @Router			/shopkeeper [get]
func (s *ShopkeeperHandler) Get(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)

	shopkeeper, err := s.service.Shopkeeper.GetById(claims.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	return c.JSON(http.StatusOK, s.service.Shopkeeper.Filter(shopkeeper))
}

// @Description	Create a Shopkeeper account.
// @Tags			shopkeeper
// @Accept			json
// @Produce		json
// @Param			body	body		dtos.CreateShopkeeperDTO	true	"Create Shopkeeper Account Input"
// @Success		201		{object}	dtos.ShopkeeperResponseDTO
// @Failure		422		{object}	errors.UnprocessableEntityError
// @Failure		409		{object}	errors.ConflictError
// @Failure		400		{object}	errors.BadRequestError
// @Router			/shopkeeper [post]
func (s *ShopkeeperHandler) Create(c echo.Context) error {
	// Bind Shopkeeper
	shopkeeper := new(dtos.CreateShopkeeperDTO)
	err := c.Bind(shopkeeper)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	// Validate body
	err = s.service.Validation.Struct(shopkeeper)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.UnprocessableEntity(err.Error()))
	}

	// Check if email exists
	_, err = s.service.Shopkeeper.GetByEmail(shopkeeper.Email)
	if err == nil {
		return c.JSON(http.StatusConflict, errors.Conflict("email", shopkeeper.Email))

	}

	// Check if CNPJ exists
	_, err = s.service.User.GetByCPF(shopkeeper.CNPJ)
	if err == nil {
		return c.JSON(http.StatusConflict, errors.Conflict("cnpj", shopkeeper.CNPJ))
	}

	// Create Wallet
	wallet, err := s.service.Wallet.Create()
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	// Create Shopkeeper
	response, err := s.service.Shopkeeper.Create(shopkeeper, wallet)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	return c.JSON(http.StatusCreated, s.service.Shopkeeper.Filter(response))
}

// @Description	Update 'Name' and/or 'Lastname' of Shopkeeper account.
// @Tags			shopkeeper
// @Accept			json
// @Produce		json
// @Param			body	body		dtos.UpdateShopkeeperDTO	true	"Shopkeeper"
// @Success		200		{object}	dtos.ShopkeeperResponseDTO
// @Failure		400		{object}	errors.BadRequestError
// @Failure		401		{object}	errors.UnauthorizedError
// @Failure		422		{object}	errors.UnprocessableEntityError
// @Security		ApiKeyAuth
// @Router			/shopkeeper [put]
func (s *ShopkeeperHandler) Update(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)

	// Bind User
	shopkeeper := new(dtos.UpdateUserDTO)
	err := c.Bind(shopkeeper)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	// Validate Body
	err = s.service.Validation.Struct(shopkeeper)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.UnprocessableEntity(err.Error()))
	}

	// Update
	response, err := s.service.User.Update(claims.ID, shopkeeper)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	return c.JSON(http.StatusOK, s.service.User.Filter(response))
}

// @Description	Login
// @Tags			shopkeeper
// @Accept			json
// @Produce		json
// @Param			body	body		dtos.LoginShopkeeperDTO	true	"Shopkeeper"
// @Success		200		{object}	dtos.LoginResponseDTO
// @Failure		400		{object}	errors.BadRequestError
// @Failure		401		{object}	errors.UnauthorizedError
// @Failure		422		{object}	errors.UnprocessableEntityError
// @Router			/shopkeeper/login [post]
func (s *ShopkeeperHandler) Login(c echo.Context) error {
	// Bind LoginShopeekerDTO
	input := new(dtos.LoginShopkeeperDTO)
	err := c.Bind(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	// Validate Body
	err = s.service.Validation.Struct(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.UnprocessableEntity(err.Error()))
	}

	// Get Shopkeeper By Email
	shopkeeper, err := s.service.Shopkeeper.GetByEmail(input.Email)
	if shopkeeper == nil {
		return c.JSON(http.StatusUnauthorized, errors.Unauthorized())
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	// Compare Password
	err = s.service.Shopkeeper.CheckPasswordHash(input.Password, shopkeeper.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, errors.Unauthorized())
	}

	// Login
	token, err := s.service.User.Login(shopkeeper)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
