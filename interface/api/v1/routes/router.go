package routes

import (
	transactorRepository "github.com/danisbagus/golang-hexagon-mongo/infrastructure/repository/mongo/transactor"

	healtCheckRepository "github.com/danisbagus/golang-hexagon-mongo/infrastructure/repository/mongo/healtchCheck"
	productRepository "github.com/danisbagus/golang-hexagon-mongo/infrastructure/repository/mongo/product"

	healtCheckService "github.com/danisbagus/golang-hexagon-mongo/core/healthCheck"
	productService "github.com/danisbagus/golang-hexagon-mongo/core/product"

	healtCheckHandler "github.com/danisbagus/golang-hexagon-mongo/interface/api/v1/healthCheck"
	productHandler "github.com/danisbagus/golang-hexagon-mongo/interface/api/v1/product"

	"github.com/danisbagus/golang-hexagon-mongo/utils/config/database"
	"github.com/labstack/echo/v4"
)

func API(route *echo.Group) {

	mongoClient := database.MongoClient
	mongoDB := database.MongoDatabase

	transactorRepository := transactorRepository.New(mongoClient)

	healthCheckRepository := healtCheckRepository.New(mongoClient)
	productRepository := productRepository.New(mongoDB, mongoClient)

	healtCheckService := healtCheckService.New(healthCheckRepository)
	productService := productService.New(productRepository, transactorRepository)

	healtCheckHandler := healtCheckHandler.New(healtCheckService)
	productHandler := productHandler.New(productService)

	healthCheckRoute := route.Group("/health")
	healthCheckRoute.GET("/ping", healtCheckHandler.Ping)

	productkRoute := route.Group("/product")
	productkRoute.POST("", productHandler.Insert)
	productkRoute.GET("", productHandler.List)
	productkRoute.GET("/:id", productHandler.View)
	productkRoute.PUT("/:id", productHandler.Update)
	productkRoute.DELETE("/:id", productHandler.Delete)
}
