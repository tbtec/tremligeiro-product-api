package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type FindProductController struct {
	usc *usecase.UscFindProduct
}

func NewFindProductController(container *container.Container) *FindProductController {
	return &FindProductController{
		usc: usecase.NewUseCaseFindProduct(
			gateway.NewProductGateway(container.ProductRepository),
			gateway.NewCategoryGateway(container.CategoryRepository),
			presenter.NewProductPresenter(),
		),
	}
}

func (ctl *FindProductController) Execute(ctx context.Context, input int) (dto.ProductContent, error) {
	return ctl.usc.FindByCategory(ctx, input)
}
