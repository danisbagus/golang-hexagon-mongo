package routes

import (
	transactorRepository "github.com/danisbagus/golang-hexagon-mongo/infrastructure/repository/mongo/transactor"

	healtCheckRepository "github.com/danisbagus/golang-hexagon-mongo/infrastructure/repository/mongo/healtchCheck"
	productRepository "github.com/danisbagus/golang-hexagon-mongo/infrastructure/repository/mongo/product"

	healtCheckService "github.com/danisbagus/golang-hexagon-mongo/core/healthCheck"
	productService "github.com/danisbagus/golang-hexagon-mongo/core/product"

	healtCheckHandler "github.com/danisbagus/golang-hexagon-mongo/interface/api/v1/healthCheck"
	productHandler "github.com/danisbagus/golang-hexagon-mongo/interface/api/v1/product"

	mongodb "github.com/danisbagus/golang-hexagon-mongo/modules/mongodb"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func API(route *echo.Group) {

	mongoClient, err := mongodb.GetClient()
	if err != nil {
		log.Fatal(err)
	}

	mongoDB, err := mongodb.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

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
