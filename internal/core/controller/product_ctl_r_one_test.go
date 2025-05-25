package controller

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/test/repository"
)

func TestFindOneProductController_Execute_Success_WithMock(t *testing.T) {

	mockProduct := &model.Product{
		ID:          "prod1",
		Name:        "Product 1",
		Description: "Description 1",
		Amount:      100.0,
		CategoryId:  1,
		CreatedAt:   time.Now(),
	}

	productRepo := &repository.MockProductRepo{
		FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {
			if id == "prod1" {
				return mockProduct, nil
			}
			return nil, errors.New("not found")
		},
	}

	categoryRepo := &repository.MockCategoryRepo{
		FindByIdFunc: func(id int) *entity.Category {
			if id == 1 {
				return &entity.Category{ID: 1, Name: "Category 1"}
			}
			return nil
		},
	}

	container := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}

	controller := NewFindOneProductController(container)

	ctx := context.Background()
	result, err := controller.Execute(ctx, "prod1")

	assert.NoError(t, err)
	assert.Equal(t, "prod1", result.ProductId)
	assert.Equal(t, "Product 1", result.Name)
}

func TestFindOneProductController_Execute_NotFound(t *testing.T) {

	productRepo := &repository.MockProductRepo{
		FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {

			return nil, errors.New("not found")
		},
	}

	categoryRepo := &repository.MockCategoryRepo{
		FindByIdFunc: func(id int) *entity.Category {
			return nil
		},
	}

	container := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}

	controller := NewFindOneProductController(container)

	ctx := context.Background()
	_, err := controller.Execute(ctx, "prod-404")

	assert.Error(t, err)
}
