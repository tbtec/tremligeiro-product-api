package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
)

type UscDeleteProduct struct {
	productGateway   *gateway.ProductGateway
	productPresenter *presenter.ProductPresenter
}

func NewUseCaseDeleteProduct(productGateway *gateway.ProductGateway,
	productPresenter *presenter.ProductPresenter) *UscDeleteProduct {
	return &UscDeleteProduct{
		productGateway:   productGateway,
		productPresenter: productPresenter,
	}
}

func (usc *UscDeleteProduct) DeleteById(ctx context.Context, id string) (string, error) {

	_, err := usc.productGateway.DeleteById(ctx, id)

	return id, err
}
