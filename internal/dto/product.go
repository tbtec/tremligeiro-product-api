package dto

import "time"

type CreateProduct struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	CategoryId  int     `json:"categoryId" validate:"required,oneof='1' '2' '3' '4'"`
	Amount      float64 `json:"amount" validate:"required"`
}

type UpdateProduct struct {
	ProductId   string
	Name        string
	Description string
	CategoryId  int
	Amount      float64
	CreatedAt   time.Time
}

type Product struct {
	ProductId   string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    Category  `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProductContent struct {
	Content []Product `json:"content"`
}
