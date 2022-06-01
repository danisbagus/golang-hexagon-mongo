package model

type Product struct {
	ID          string
	Name        string
	CategoryIDs []uint64
	Price       uint64
}
