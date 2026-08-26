package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/pkg/idmask"
	"github.com/farhapartex/lending-platform/core-service/pkg/queryparam"
)

type LiquidationHandlerParams struct {
	Liquidations domain.LiquidationService
	Masker       *idmask.Masker
}

type LiquidationHandler struct {
	liquidations domain.LiquidationService
	masker       *idmask.Masker
}

func NewLiquidationHandler(params LiquidationHandlerParams) *LiquidationHandler {
	return &LiquidationHandler{
		liquidations: params.Liquidations,
		masker:       params.Masker,
	}
}

func (h *LiquidationHandler) ListHistory(c *gin.Context) {
	marketID, err := h.marketFilter(c)
	if err != nil {
		respondBadRequest(c, "That market identifier is not valid.")

		return
	}

	request, err := dto.ParseLiquidationListRequest(marketID, c.Request.URL.Query())
	if err != nil {
		respondBadRequest(c, queryparam.Message(err))

		return
	}

	page, err := h.liquidations.List(c.Request.Context(), request)
	if err != nil {
		respondDomainError(c, err, "There are no liquidations to show.")

		return
	}

	response, err := dto.NewLiquidationListResponse(page, h.maskLiquidationID)
	if err != nil {
		respondInternalError(c)

		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *LiquidationHandler) GetReceipt(c *gin.Context) {
	id, err := h.masker.Unmask(idmask.KindLiquidation, c.Param("liquidationId"))
	if err != nil {
		respondBadRequest(c, "That liquidation identifier is not valid.")

		return
	}

	liquidation, err := h.liquidations.ByID(c.Request.Context(), id)
	if err != nil {
		respondDomainError(c, err, "That liquidation does not exist.")

		return
	}

	publicID, err := h.masker.Mask(idmask.KindLiquidation, liquidation.ID)
	if err != nil {
		respondInternalError(c)

		return
	}

	c.JSON(http.StatusOK, dto.NewLiquidationResponse(liquidation, publicID))
}

func (h *LiquidationHandler) marketFilter(c *gin.Context) (*int64, error) {
	raw := queryparam.String(c.Request.URL.Query(), dto.ParamMarket)
	if raw == "" {
		return nil, nil
	}

	marketID, err := h.masker.Unmask(idmask.KindMarket, raw)
	if err != nil {
		return nil, err
	}

	return &marketID, nil
}

func (h *LiquidationHandler) maskLiquidationID(id int64) (string, error) {
	return h.masker.Mask(idmask.KindLiquidation, id)
}
