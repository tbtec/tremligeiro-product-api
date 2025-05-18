package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UscFindProduct struct {
	productGateway   *gateway.ProductGateway
	categoryGateway  *gateway.CategoryGateway
	productPresenter *presenter.ProductPresenter
}

func NewUseCaseFindProduct(productGateway *gateway.ProductGateway,
	categoryGateway *gateway.CategoryGateway,
	productPresenter *presenter.ProductPresenter) *UscFindProduct {
	return &UscFindProduct{
		productGateway:   productGateway,
		categoryGateway:  categoryGateway,
		productPresenter: productPresenter,
	}
}

func (usc *UscFindProduct) FindByCategory(ctx context.Context, categoryId int) (dto.ProductContent, error) {

	category := usc.categoryGateway.FindById(categoryId)
	if category == nil {
		return dto.ProductContent{}, ErrCategoryNotExists
	}

	products, error := usc.productGateway.FindByCategory(ctx, categoryId)
	if error != nil {
		return dto.ProductContent{}, error
	}

	return usc.productPresenter.BuildProductContentResponse(products, *category), nil
}
