package routers

import (
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/api/handlers"
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, currenctService service.CurrencyService) {
	currencyHandler := handlers.NewCurrencyHandler(currenctService)

	router.GET("/nbp/:currency/:startDate/:endDate", currencyHandler.GetAverageCurrencyRate)
}
