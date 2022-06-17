package main

import (
	"time"

	v1Router "github.com/danisbagus/golang-hexagon-mongo/interface/api/v1/routes"

	mongodb "github.com/danisbagus/golang-hexagon-mongo/modules/mongodb"
	"github.com/danisbagus/golang-hexagon-mongo/utils/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// get app config
	appConfig := config.GetAPPConfig()

	// init modules
	mongodb.Init(appConfig.MongoDatabase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(ServiceRequestTime)

	v1 := e.Group("/v1")
	v1Router.API(v1)

	port := "5000"
	if err := e.Start(":" + port); err != nil {
		e.Logger.Info("Shutting down the server")
	}

}

func ServiceRequestTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("X-App-RequestTime", time.Now().Format(time.RFC3339))
		return next(c)
	}
}
