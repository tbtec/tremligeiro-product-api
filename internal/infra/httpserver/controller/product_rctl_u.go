package controller

import (
	"context"
	"time"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type ProductUpdateController struct {
	controller *ctl.UpdateProductController
}

type ProductUpdateRequest struct {
	ProductId   string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryId  int     `json:"categoryId"`
	Amount      float64 `json:"amount"`
}

type ProductUpdateResponse struct {
	ProductId   string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Amount      float64          `json:"amount"`
	Category    CategoryResponse `json:"category"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type CategoryUpdateResponse struct {
	CategoryID   int    `json:"id"`
	CategoryName string `json:"name"`
}

func NewProductUpdateByIdController(container *container.Container) httpserver.IController {
	return &ProductUpdateController{
		controller: ctl.NewUpdateProductController(container),
	}
}

func (controller *ProductUpdateController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	productRequest := ProductUpdateRequest{}

	product_id := request.ParseParamString("productId")

	errBody := request.ParseBody(ctx, &productRequest)
	if errBody != nil {
		return httpserver.HandleError(ctx, errBody)
	}

	command := productRequest.toCommand(product_id)

	product, err := controller.controller.Execute(ctx, command)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(product)
}

func (request *ProductUpdateRequest) toCommand(productId string) dto.UpdateProduct {
	return dto.UpdateProduct{
		ProductId:   productId,
		Name:        request.Name,
		Description: request.Description,
		CategoryId:  request.CategoryId,
		Amount:      request.Amount,
	}
}

func buildProductUpdateResponse(output usecase.CreateProductOutput) ProductUpdateResponse {
	return ProductUpdateResponse{
		ProductId:   output.ProductId,
		Name:        output.Name,
		Description: output.Description,
		Amount:      output.Amount,
		Category: CategoryResponse{
			CategoryID:   output.CategoryID,
			CategoryName: output.CategoryName,
		},
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	}
}
