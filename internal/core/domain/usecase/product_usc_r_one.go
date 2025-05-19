package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UscFindOneProduct struct {
	productGateway   *gateway.ProductGateway
	categoryGateway  *gateway.CategoryGateway
	productPresenter *presenter.ProductPresenter
}

func NewUseCaseFindOneProduct(productGateway *gateway.ProductGateway,
	categoryGateway *gateway.CategoryGateway,
	productPresenter *presenter.ProductPresenter) *UscFindOneProduct {
	return &UscFindOneProduct{
		productGateway:   productGateway,
		categoryGateway:  categoryGateway,
		productPresenter: productPresenter,
	}
}

func (usc *UscFindOneProduct) FindByProductId(ctx context.Context, productId string) (dto.Product, error) {

	product, error := usc.productGateway.FindOne(ctx, productId)

	if error != nil {
		return dto.Product{}, error
	}

	categoryId := product.CategoryId
	category := usc.categoryGateway.FindById(categoryId)
	if category == nil {
		return dto.Product{}, ErrCategoryNotExists
	}

	return usc.productPresenter.BuildOneProductContentResponse(*product, *category), nil
}
