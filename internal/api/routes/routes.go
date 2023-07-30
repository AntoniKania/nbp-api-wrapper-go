package routers

import (
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/api/handlers"
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, currencyService service.CurrencyService) {
	currencyHandler := handlers.NewCurrencyHandler(currencyService)

	router.GET("/nbp/:currency/:startDate/:endDate", currencyHandler.GetAverageCurrencyRate)
}
