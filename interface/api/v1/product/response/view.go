package response

import (
	"github.com/danisbagus/golang-hexagon-mongo/core/model"
)

type ViewReponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID uint64 `json:"category_id"`
	Price      uint64 `json:"price"`
}

func NewViewReponse(product *model.Product, message string) interface{} {
	data := new(ViewReponse)
	data.ID = product.ID
	data.Name = product.Name
	data.CategoryID = product.CategoryID
	data.Price = product.Price

	response := map[string]interface{}{
		"message": message,
		"data":    data,
	}

	return response
}
