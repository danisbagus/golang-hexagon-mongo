package response

import (
	"github.com/danisbagus/golang-hexagon-mongo/core/model"
)

type ListReponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	CategoryIDs []uint64 `json:"category_ids"`
	Price       uint64   `json:"price"`
}

func NewListResponse(products []model.Product, message string) interface{} {
	listData := make([]ListReponse, 0)

	for _, product := range products {
		var data ListReponse
		data.ID = product.ID
		data.Name = product.Name
		data.CategoryIDs = product.CategoryIDs
		data.Price = product.Price

		listData = append(listData, data)
	}

	response := map[string]interface{}{
		"message": message,
		"data":    listData,
	}
	return response
}
