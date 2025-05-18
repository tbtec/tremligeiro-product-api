package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type DeleteProductController struct {
	usc *usecase.UscDeleteProduct
}

func NewDeleteProductController(container *container.Container) *DeleteProductController {
	return &DeleteProductController{
		usc: usecase.NewUseCaseDeleteProduct(
			gateway.NewProductGateway(container.ProductRepository),
			presenter.NewProductPresenter(),
		),
	}
}

func (ctl *DeleteProductController) Execute(ctx context.Context, input string) (string, error) {
	return ctl.usc.DeleteById(ctx, input)
}
