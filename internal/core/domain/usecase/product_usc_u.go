package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UscUpdateProduct struct {
	productGateway   *gateway.ProductGateway
	categoryGateway  *gateway.CategoryGateway
	productPresenter *presenter.ProductPresenter
}

func NewUseCaseUpdateProduct(productGateway *gateway.ProductGateway,
	categoryGateway *gateway.CategoryGateway,
	productPresenter *presenter.ProductPresenter) *UscUpdateProduct {
	return &UscUpdateProduct{
		productGateway:   productGateway,
		categoryGateway:  categoryGateway,
		productPresenter: productPresenter,
	}
}

func (usc *UscUpdateProduct) UpdateById(ctx context.Context, command dto.UpdateProduct) (dto.Product, error) {

	product, error := usc.productGateway.UpdateById(ctx, command)
	if error != nil {
		return dto.Product{}, error
	}

	category := usc.categoryGateway.FindById(product.CategoryId)
	if category == nil {
		return dto.Product{}, ErrCategoryNotExists
	}

	return usc.productPresenter.BuildProductCreateResponse(product, *category), nil
}
