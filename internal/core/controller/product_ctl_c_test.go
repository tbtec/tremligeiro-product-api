package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/test/fixtures"
)

func TestCreateProductController_Execute_Success(t *testing.T) {
	ctx := context.Background()

	// Mock repositories
	productRepo := &fixtures.MockProductRepo{
		CreateFunc: func(ctx context.Context, p *model.Product) error {
			return nil
		},
	}
	categoryRepo := &fixtures.MockCategoryRepo{
		FindByIdFunc: func(id int) *entity.Category {
			if id == 1 {
				return &entity.Category{ID: 1, Name: "Category 1"}
			}
			return nil
		},
	}

	testContainer := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}

	controller := NewCreateProductController(testContainer)

	input := dto.CreateProduct{
		Name:        "Product 1",
		Description: "Description 1",
		CategoryId:  1,
		Amount:      100.0,
	}

	result, err := controller.Execute(ctx, input)
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", result.Name)
	assert.Equal(t, 1, result.Category.ID)
}

func TestCreateProductController_Execute_CategoryNotFound(t *testing.T) {
	ctx := context.Background()

	// Mock repositories
	productRepo := &fixtures.MockProductRepo{
		CreateFunc: func(ctx context.Context, p *model.Product) error {
			return nil
		},
	}
	categoryRepo := &fixtures.MockCategoryRepo{
		FindByIdFunc: func(id int) *entity.Category {
			return nil
		},
	}

	testContainer := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}

	controller := NewCreateProductController(testContainer)

	input := dto.CreateProduct{
		Name:        "Product 2",
		Description: "Description 2",
		CategoryId:  3,
		Amount:      100.0,
	}

	_, err := controller.Execute(ctx, input)
	assert.Error(t, err)
}
