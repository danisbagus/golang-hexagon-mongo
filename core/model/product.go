package model

import "time"

type Product struct {
	ID          string
	Name        string
	CategoryIDs []uint64
	Price       uint64
	CreatedAt   time.Time
}
