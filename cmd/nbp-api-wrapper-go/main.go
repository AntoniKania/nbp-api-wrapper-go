package main

import (
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/api/routes"
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	service := service.NewCurrencyService()

	routers.RegisterRoutes(router, service)

	router.Run("localhost:8080")
}
