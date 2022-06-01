package request

type UpdateRequest struct {
	Name        string   `json:"name"`
	CategoryIDs []uint64 `json:"category_ids"`
	Price       uint64   `json:"price"`
}
