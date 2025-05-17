package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type CreateProductController struct {
	container *container.Container
	usc       *usecase.UscCreateProduct
}

func NewCreateProductController(container *container.Container) *CreateProductController {
	return &CreateProductController{
		container: container,
		usc: usecase.NewUseCaseCreateProduct(
			gateway.NewProductGateway(container.ProductRepository),
			gateway.NewCategoryGateway(container.CategoryRepository),
			presenter.NewProductPresenter(),
		),
	}
}

func (ctl *CreateProductController) Execute(ctx context.Context, product dto.CreateProduct) (dto.Product, error) {
	return ctl.usc.Create(ctx, product)
}
