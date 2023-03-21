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

type TransactionHandler struct {
	service *services.Services
	env     *config.Envs
}

// @Description	Create a transaction.
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Param			body	body		dtos.CreateTransactionDTO	true	"Create Transaction Input"
// @Success		201		{object}	dtos.ResponseTransactionDTO
// @Failure		401		{object}	errors.UnauthorizedError
// @Failure		400		{object}	errors.BadRequestError
// @Security		ApiKeyAuth
// @Router			/transaction [post]
func (tx *TransactionHandler) Create(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)

	txInput := new(dtos.CreateTransactionDTO)
	err := c.Bind(txInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	txOutput, err := tx.service.Transaction.CreateTransaction(claims.ID, claims.Type, *txInput)

	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	return c.JSON(http.StatusCreated, txOutput)
}

// @Description	Get transaction from Wallet.
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Success		200	{object}	dtos.ResponseTransactionDTO
// @Failure		400	{object}	errors.BadRequestError
// @Failure		401	{object}	errors.UnauthorizedError
// @Security		ApiKeyAuth
// @Router			/transaction [get]
func (tx *TransactionHandler) List(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)

	txOutputs, err := tx.service.Transaction.List(claims.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.BadRequest())
	}

	return c.JSON(http.StatusOK, txOutputs)
}
