package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/middleware"
	"github.com/farhapartex/lending-platform/core-service/pkg/idmask"
)

type AccountHandlerParams struct {
	Transactions domain.TransactionService
	Masker       *idmask.Masker
}

type AccountHandler struct {
	transactions domain.TransactionService
	masker       *idmask.Masker
}

func NewAccountHandler(params AccountHandlerParams) *AccountHandler {
	return &AccountHandler{
		transactions: params.Transactions,
		masker:       params.Masker,
	}
}

func (h *AccountHandler) GetTransaction(c *gin.Context) {
	address := c.Param("address")

	id, err := h.masker.Unmask(idmask.KindTransaction, c.Param("transactionId"))
	if err != nil {
		respondBadRequest(c, "That transaction identifier is not valid.")

		return
	}

	transaction, err := h.transactions.ByID(c.Request.Context(), address, id)
	if err != nil {
		respondDomainError(c, err, "That transaction does not exist.")

		return
	}

	publicID, err := h.masker.Mask(idmask.KindTransaction, transaction.ID)
	if err != nil {
		respondInternalError(c)

		return
	}

	c.JSON(http.StatusOK, dto.NewTransactionResponse(transaction, publicID))
}

func respondDomainError(c *gin.Context, err error, notFoundMessage string) {
	switch {
	case errors.Is(err, domain.ErrInvalidInput):
		respondBadRequest(c, "That request could not be read.")
	case errors.Is(err, domain.ErrNotFound):
		respondNotFound(c, notFoundMessage)
	default:
		respondInternalError(c)
	}
}

func respondBadRequest(c *gin.Context, message string) {
	c.JSON(
		http.StatusBadRequest,
		dto.NewErrorResponse(dto.CodeBadRequest, message, middleware.RequestIDFrom(c)),
	)
}

func respondNotFound(c *gin.Context, message string) {
	c.JSON(
		http.StatusNotFound,
		dto.NewErrorResponse(dto.CodeNotFound, message, middleware.RequestIDFrom(c)),
	)
}

func respondInternalError(c *gin.Context) {
	c.JSON(
		http.StatusInternalServerError,
		dto.NewErrorResponse(dto.CodeInternalError, "Something went wrong on our side.", middleware.RequestIDFrom(c)),
	)
}
