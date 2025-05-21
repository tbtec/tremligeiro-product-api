package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type FindOneProductController struct {
	usc *usecase.UscFindOneProduct
}

func NewFindOneProductController(container *container.Container) *FindOneProductController {
	return &FindOneProductController{
		usc: usecase.NewUseCaseFindOneProduct(
			gateway.NewProductGateway(container.ProductRepository),
			gateway.NewCategoryGateway(container.CategoryRepository),
			presenter.NewProductPresenter(),
		),
	}
}

func (ctl *FindOneProductController) Execute(ctx context.Context, input string) (dto.Product, error) {
	return ctl.usc.FindByProductId(ctx, input)
}
