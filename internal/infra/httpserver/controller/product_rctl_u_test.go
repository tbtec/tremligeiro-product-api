package controller

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/test/repository"
)

func TestBuildProductUpdateResponse(t *testing.T) {
	now := time.Now()
	output := usecase.CreateProductOutput{
		ProductId:    "prod-123",
		Name:         "Test Product",
		Description:  "A product for testing",
		Amount:       99.99,
		CategoryID:   42,
		CategoryName: "Test Category",
		CreatedAt:    now,
		UpdatedAt:    now.Add(time.Hour),
	}

	resp := buildProductUpdateResponse(output)

	assert.Equal(t, output.ProductId, resp.ProductId)
	assert.Equal(t, output.Name, resp.Name)
	assert.Equal(t, output.Description, resp.Description)
	assert.Equal(t, output.Amount, resp.Amount)
	assert.Equal(t, output.CategoryID, resp.Category.CategoryID)
	assert.Equal(t, output.CategoryName, resp.Category.CategoryName)
	assert.Equal(t, output.CreatedAt, resp.CreatedAt)
	assert.Equal(t, output.UpdatedAt, resp.UpdatedAt)
}

func TestProductUpdateController_Handle_Success(t *testing.T) {

	mockProduct := &model.Product{
		ID:          "prod1",
		Name:        "Product 1",
		Description: "Description 1",
		Amount:      100.0,
		CategoryId:  1,
		CreatedAt:   time.Now(),
	}

	container := &container.Container{

		ProductRepository: &repository.MockProductRepo{

			FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {
				if id == "prod1" {
					return mockProduct, nil
				}
				return nil, errors.New("not found")
			},
		},

		CategoryRepository: &repository.MockCategoryRepo{
			FindByIdFunc: func(id int) *entity.Category {
				if id == 1 {
					return &entity.Category{ID: 1, Name: "Category 1"}
				}
				return nil
			},
		},
	}
	ctrl := NewProductUpdateByIdController(container)

	input := ProductUpdateRequest{
		Name:        "Produto Teste",
		Description: "Descrição",
		CategoryId:  1,
		Amount:      10.0,
	}
	inputBytes, _ := json.Marshal(input)
	req := httpserver.Request{
		Params: map[string]string{"productId": "prod1"},
		Body:   inputBytes,
	}

	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, 200, resp.Code)
}

func TestProductUpdateController_Handle_ExecuteError(t *testing.T) {

	container := &container.Container{

		ProductRepository: &repository.MockProductRepo{

			FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {
				return nil, errors.New("not found")
			},
		},
	}
	ctrl := NewProductUpdateByIdController(container)

	input := ProductUpdateRequest{
		Name:        "Produto Teste",
		Description: "Descrição",
		CategoryId:  1,
		Amount:      10.0,
	}
	inputBytes, _ := json.Marshal(input)
	req := httpserver.Request{
		Params: map[string]string{"productId": "123"},
		Body:   inputBytes,
	}

	resp := ctrl.Handle(context.Background(), req)

	assert.NotEqual(t, 200, resp.Code)
	// Ajuste conforme o tipo retornado por HandleError
	errMsg, ok := resp.Body.(httpserver.ErrorMessage)
	assert.True(t, ok)
	assert.Contains(t, errMsg.Error.Description, "Internal Server Error")
}
