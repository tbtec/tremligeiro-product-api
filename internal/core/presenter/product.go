package presenter

import (
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
)

// import "github.com/tbtec/tremligeiro/internal/core/usecase"

type ProductPresenter struct {
}

func NewProductPresenter() *ProductPresenter {
	return &ProductPresenter{}
}

func (presenter *ProductPresenter) BuildProductCreateResponse(product entity.Product, category entity.Category) dto.Product {
	return dto.Product{
		ProductId:   product.ID,
		Name:        product.Name,
		Description: product.Description,
		Amount:      product.Amount,
		Category: dto.Category{
			ID:   category.ID,
			Name: category.Name,
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func (presenter *ProductPresenter) BuildProductContentResponse(products []entity.Product, category entity.Category) dto.ProductContent {
	response := []dto.Product{}

	for _, product := range products {
		response = append(response, presenter.BuildProductCreateResponse(product, category))
	}

	return dto.ProductContent{Content: response}
}
