package controller

import (
	"context"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type ProductFindOneController struct {
	controller *ctl.FindOneProductController
}

func NewProductFindOneRestController(container *container.Container) httpserver.IController {
	return &ProductFindOneController{
		controller: ctl.NewFindOneProductController(container),
	}
}

func (controller *ProductFindOneController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	product_id := request.ParseParamString("productId")

	product, err := controller.controller.Execute(ctx, product_id)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(product)
}
