package main

import (
	"database/sql"
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/api/routes"
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/database/repository"
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/service"
	_ "github.com/lib/pq"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	connStr := "postgres://nbp:root@localhost/nbp?sslmode=disable"
	db, _ := sql.Open("postgres", connStr)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	requestRepository := repository.NewRequestRepository(db)
	currencyService := service.NewCurrencyService(requestRepository)

	router := gin.Default()
	routers.RegisterRoutes(router, currencyService)

	router.Run("localhost:8080")
}
