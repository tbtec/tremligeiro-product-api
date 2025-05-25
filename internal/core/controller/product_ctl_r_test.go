package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/test/repository"
)

func TestFindProductController_Execute_Success(t *testing.T) {

	mockProductsSlice := []model.Product{
		{ID: "prod1", Name: "Product 1", CategoryId: 10},
		{ID: "prod2", Name: "Product 2", CategoryId: 10},
	}
	mockProducts := &mockProductsSlice

	productRepo := &repository.MockProductRepo{
		FindByCategoryFunc: func(ctx context.Context, categoryId int) (*[]model.Product, error) {
			if categoryId == 10 {

				return mockProducts, nil
			}

			return nil, errors.New("not found")
		},
	}

	categoryRepo := &repository.MockCategoryRepo{
		FindByIdFunc: func(id int) *entity.Category {
			if id == 10 {
				return &entity.Category{ID: 10, Name: "Category 10"}
			}
			return nil
		},
	}

	container := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}

	controller := NewFindProductController(container)

	ctx := context.Background()
	result, err := controller.Execute(ctx, 10)

	assert.NoError(t, err)
	assert.Len(t, result.Content, 2)
	assert.Equal(t, "Product 1", result.Content[0].Name)
	assert.Equal(t, "Product 2", result.Content[1].Name)
}

func TestFindProductController_Execute_Error(t *testing.T) {

	productRepo := &repository.MockProductRepo{
		CreateFunc: func(ctx context.Context, product *model.Product) error {
			return nil
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
	controller := NewFindProductController(container)

	ctx := context.Background()
	_, err := controller.Execute(ctx, 10)

	assert.Error(t, err)
}
