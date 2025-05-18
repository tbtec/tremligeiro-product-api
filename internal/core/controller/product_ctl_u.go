package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type UpdateProductController struct {
	usc *usecase.UscUpdateProduct
}

func NewUpdateProductController(container *container.Container) *UpdateProductController {
	return &UpdateProductController{
		usc: usecase.NewUseCaseUpdateProduct(
			gateway.NewProductGateway(container.ProductRepository),
			gateway.NewCategoryGateway(container.CategoryRepository),
			presenter.NewProductPresenter(),
		),
	}
}

func (ctl *UpdateProductController) Execute(ctx context.Context, command dto.UpdateProduct) (dto.Product, error) {
	return ctl.usc.UpdateById(ctx, command)
}
