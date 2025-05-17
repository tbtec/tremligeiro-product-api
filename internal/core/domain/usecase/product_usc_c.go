package usecase

import (
	"context"
	"time"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/types/ulid"
)

type UscCreateProduct struct {
	productGateway   *gateway.ProductGateway
	categoryGateway  *gateway.CategoryGateway
	productPresenter *presenter.ProductPresenter
}

func NewUseCaseCreateProduct(productGateway *gateway.ProductGateway,
	categoryGateway *gateway.CategoryGateway,
	productPresenter *presenter.ProductPresenter) *UscCreateProduct {
	return &UscCreateProduct{
		productGateway:   productGateway,
		categoryGateway:  categoryGateway,
		productPresenter: productPresenter,
	}
}

func (usc *UscCreateProduct) Create(ctx context.Context, productDto dto.CreateProduct) (dto.Product, error) {

	category := usc.categoryGateway.FindById(productDto.CategoryId)
	if category == nil {
		return dto.Product{}, ErrCategoryNotExists
	}

	//p:= entity.NewProduct(DTO)
	product := entity.Product{
		ID:          ulid.NewUlid().String(),
		Name:        productDto.Name,
		Description: productDto.Description,
		CategoryId:  productDto.CategoryId,
		Amount:      productDto.Amount,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err := usc.productGateway.Create(ctx, &product)
	if err != nil {
		return dto.Product{}, err
	}

	return usc.productPresenter.BuildProductCreateResponse(product, *category), nil
}
