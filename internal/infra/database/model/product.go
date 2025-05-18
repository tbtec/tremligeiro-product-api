package model

import "time"

type Product struct {
	ID          string    `gorm:"column:product_id;primaryKey"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CategoryId  int       `gorm:"column:category_id"`
	Amount      float64   `gorm:"column:amount"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
