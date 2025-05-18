package controller

import (
	"context"
	"time"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/internal/validator"
)

type ProductCreateRestController struct {
	controller *ctl.CreateProductController
}

type ProductCreateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	CategoryId  int     `json:"categoryId" validate:"required,oneof='1' '2' '3' '4'"`
	Amount      float64 `json:"amount" validate:"required"`
}

type ProductCreateResponse struct {
	ProductId   string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Amount      float64          `json:"amount"`
	Category    CategoryResponse `json:"category"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type CategoryResponse struct {
	CategoryID   int    `json:"id"`
	CategoryName string `json:"name"`
}

func NewProductCreateRestController(container *container.Container) httpserver.IController {
	return &ProductCreateRestController{
		controller: ctl.NewCreateProductController(container),
	}
}

func (ctl *ProductCreateRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	productRequest := dto.CreateProduct{}

	errBody := request.ParseBody(ctx, &productRequest)
	if errBody != nil {
		return httpserver.HandleError(ctx, errBody)
	}

	err := validator.Validate(productRequest)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	output, err := ctl.controller.Execute(ctx, productRequest)

	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}
