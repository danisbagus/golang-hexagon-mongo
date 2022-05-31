package routes

import (
	healtCheckService "github.com/danisbagus/golang-hexagon-mongo/core/healthCheck"

	healtCheckRepository "github.com/danisbagus/golang-hexagon-mongo/infrastructure/repository/mongo/healtchCheck"

	healtCheckHandler "github.com/danisbagus/golang-hexagon-mongo/interface/api/v1/healthCheck"

	"github.com/danisbagus/golang-hexagon-mongo/utils/config/database"
	"github.com/labstack/echo/v4"
)

func API(route *echo.Group) {

	mongoClient := database.MongoClient
	// mongoDB := database.MongoDatabase

	healthCheckRepository := healtCheckRepository.New(mongoClient)

	healtCheckService := healtCheckService.New(healthCheckRepository)

	healtCheckHandler := healtCheckHandler.New(healtCheckService)

	healthCheckRoute := route.Group("/health")

	healthCheckRoute.GET("/ping", healtCheckHandler.Ping)

}
