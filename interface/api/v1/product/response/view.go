package response

import (
	"github.com/danisbagus/golang-hexagon-mongo/core/model"
)

type ViewReponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	CategoryIDs []uint64   `json:"category_ids"`
	Price       uint64     `json:"price"`
	Categories  []Category `json:"categories"`
}

type Category struct {
	CategoryID uint64
	Name       string
}

func NewViewReponse(product *model.Product, message string) interface{} {
	data := new(ViewReponse)
	data.ID = product.ID
	data.Name = product.Name
	data.CategoryIDs = product.CategoryIDs
	data.Price = product.Price

	response := map[string]interface{}{
		"message": message,
		"data":    data,
	}

	return response
}
