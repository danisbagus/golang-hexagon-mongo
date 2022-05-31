package product

import (
	"net/http"

	"github.com/danisbagus/golang-hexagon-mongo/core/model"
	port "github.com/danisbagus/golang-hexagon-mongo/core/port/product"
	"github.com/danisbagus/golang-hexagon-mongo/interface/api/v1/product/request"
	"github.com/danisbagus/golang-hexagon-mongo/interface/api/v1/product/response"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service port.Service
}

func New(service port.Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) Insert(c echo.Context) error {
	reqData := new(request.InsertRequest)
	if err := c.Bind(reqData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	form := new(model.Product)
	form.Name = reqData.Name
	form.CategoryID = reqData.CategoryID
	form.Price = reqData.Price

	err := h.service.Insert(form)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"message": "Successfully insert data"}
	return c.JSON(http.StatusOK, response)
}

func (h Handler) List(c echo.Context) error {
	products, err := h.service.List()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	resData := response.NewListResponse(products, "Successfully get data")
	return c.JSON(http.StatusOK, resData)
}

func (h Handler) View(c echo.Context) error {
	productID := c.Param("id")
	product, err := h.service.View(productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	resData := response.NewViewReponse(product, "Successfully get data")
	return c.JSON(http.StatusOK, resData)
}

func (h Handler) Update(c echo.Context) error {
	reqData := new(request.UpdateRequest)
	if err := c.Bind(reqData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	form := new(model.Product)
	form.ID = c.Param("id")
	form.Name = reqData.Name
	form.CategoryID = reqData.CategoryID
	form.Price = reqData.Price

	err := h.service.Update(form)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"message": "Successfully update data"}
	return c.JSON(http.StatusOK, response)
}

func (h Handler) Delete(c echo.Context) error {
	productID := c.Param("id")
	err := h.service.Delete(productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{"message": "Successfully delete data"}
	return c.JSON(http.StatusOK, response)
}
