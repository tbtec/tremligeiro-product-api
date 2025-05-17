package controller

import (
	"context"
	"log/slog"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type ProductDeleteController struct {
	controller *ctl.DeleteProductController
}

func NewProductDeleteByIdRestController(container *container.Container) httpserver.IController {
	return &ProductDeleteController{
		controller: ctl.NewDeleteProductController(container),
	}
}

func (controller *ProductDeleteController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	product_id := request.ParseParamString("productId")

	output, err := controller.controller.Execute(ctx, product_id)

	if err != nil || output == "" {
		if err.Error() == "record not found" {
			slog.ErrorContext(ctx, (err.Error() + " id: " + product_id))
			return httpserver.NotFound(err)
		}
		slog.ErrorContext(ctx, err.Error())
		return httpserver.UnprocessableEntity(err)
	}

	return httpserver.NoContent()

}
