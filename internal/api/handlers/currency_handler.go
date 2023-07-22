package handlers

import (
	"net/http"

	"github.com/AntoniKania/nbp-api-wrapper-go/internal/service"
	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	currencyService service.CurrencyService
}

func NewCurrencyHandler(currencyService service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{
		currencyService: currencyService,
	}
}

func (h *CurrencyHandler) GetAverageCurrencyRate(c *gin.Context) {
	startDate := c.Param("startDate")
	endDate := c.Param("endDate")
	currency := c.Param("currency")

	averageRate := h.currencyService.GetAverageCurrencyRate(c, currency, startDate, endDate)

	c.JSON(http.StatusOK, averageRate)
}
