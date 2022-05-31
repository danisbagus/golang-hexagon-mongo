package request

type UpdateRequest struct {
	Name       string `json:"name"`
	CategoryID uint64 `json:"category_id"`
	Price      uint64 `json:"price"`
}
