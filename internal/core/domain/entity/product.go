package entity

import (
	"time"

	"github.com/tbtec/tremligeiro/internal/types/ulid"
)

type Product struct {
	ID          string
	Name        string
	Description string
	CategoryId  int
	Amount      float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(name string, description string, categoryId int, amount float64) (*Product, error) {

	return &Product{
		ID:          ulid.NewUlid().String(),
		Name:        name,
		Description: description,
		CategoryId:  categoryId,
		Amount:      amount,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}, nil
}
