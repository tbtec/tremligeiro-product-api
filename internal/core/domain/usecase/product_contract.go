package usecase

import (
	"time"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/types/ulid"
	"github.com/tbtec/tremligeiro/internal/types/xerrors"
)

var (
	ErrCategoryNotExists = xerrors.NewBusinessError("TL-PRODUCT-001", "Category not exists")
	ErrProductNotFound   = xerrors.NewBusinessError("TL-PRODUCT-002", "Product not found")
)

type CmdCreateProduct struct {
	Name        string
	Description string
	CategoryId  int
	Amount      float64
}

type CmdUpdateProduct struct {
	ProductId   string
	Name        string
	Description string
	CategoryId  int
	Amount      float64
	CreatedAt   time.Time
}

func (command *CmdCreateProduct) ToNewEntity() entity.Product {
	return entity.Product{
		ID:          ulid.NewUlid().String(),
		Name:        command.Name,
		Description: command.Description,
		CategoryId:  command.CategoryId,
		Amount:      command.Amount,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

func (command *CmdUpdateProduct) ToUpdateEntity() entity.Product {
	return entity.Product{
		ID:          command.ProductId,
		Name:        command.Name,
		Description: command.Description,
		CategoryId:  command.CategoryId,
		Amount:      command.Amount,
		CreatedAt:   command.CreatedAt,
		UpdatedAt:   time.Now().UTC(),
	}
}

type CreateProductOutput struct {
	ProductId    string
	Name         string
	Description  string
	Amount       float64
	CategoryID   int
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
