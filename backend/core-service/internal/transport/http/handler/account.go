package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/middleware"
	"github.com/farhapartex/lending-platform/core-service/pkg/idmask"
	"github.com/farhapartex/lending-platform/core-service/pkg/queryparam"
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

func (h *AccountHandler) ListTransactions(c *gin.Context) {
	request, err := dto.ParseTransactionListRequest(c.Param("address"), c.Request.URL.Query())
	if err != nil {
		respondBadRequest(c, queryparam.Message(err))

		return
	}

	page, err := h.transactions.List(c.Request.Context(), request)
	if err != nil {
		respondDomainError(c, err, "That wallet has no history.")

		return
	}

	response, err := dto.NewTransactionListResponse(page, h.maskTransactionID)
	if err != nil {
		respondInternalError(c)

		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *AccountHandler) GetActivity(c *gin.Context) {
	limit, err := dto.ParseLimit(c.Request.URL.Query())
	if err != nil {
		respondBadRequest(c, queryparam.Message(err))

		return
	}

	page, err := h.transactions.RecentActivity(c.Request.Context(), c.Param("address"), limit)
	if err != nil {
		respondDomainError(c, err, "That wallet has no activity.")

		return
	}

	response, err := dto.NewActivityResponse(page, h.maskTransactionID)
	if err != nil {
		respondInternalError(c)

		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *AccountHandler) maskTransactionID(id int64) (string, error) {
	return h.masker.Mask(idmask.KindTransaction, id)
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
