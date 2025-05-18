package controller

import (
	"context"
	"strconv"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type ProductFindController struct {
	controller *ctl.FindProductController
}

func NewProductFindByCategoryRestController(container *container.Container) httpserver.IController {
	return &ProductFindController{
		controller: ctl.NewFindProductController(container),
	}
}

func (controller *ProductFindController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	command, err := strconv.Atoi(request.Query["categoryId"])
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	product, err := controller.controller.Execute(ctx, command)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(product)
}
