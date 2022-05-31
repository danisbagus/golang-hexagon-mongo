package healthcheck

import (
	"net/http"

	port "github.com/danisbagus/golang-hexagon-mongo/core/port/healthCheck"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service port.Service
}

func New(service port.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Ping(c echo.Context) error {
	responsePing := h.service.Ping()
	response := map[string]interface{}{"message": "Success ping", "data": responsePing}
	return c.JSON(http.StatusOK, response)
}
