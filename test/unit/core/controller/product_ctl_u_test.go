package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/test/fixtures"
)

func TestUpdateProductController_Execute_Success(t *testing.T) {

	mockProduct := &model.Product{
		ID:         "prod-1",
		Name:       "Old Name",
		CategoryId: 1,
		Amount:     10.0,
	}

	// Mock repositories
	productRepo := &fixtures.MockProductRepo{

		UpdateByIdFunc: func(ctx context.Context, p *model.Product) error {
			if p.ID == "prod-1" {
				return nil
			}
			return errors.New("not found")
		},

		FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {
			if id == "prod-1" {
				return mockProduct, nil
			}
			return nil, errors.New("not found")
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

	container := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}

	controller := controller.NewUpdateProductController(container)

	updateCmd := dto.UpdateProduct{
		ProductId:  "prod-1",
		Name:       "New Name",
		CategoryId: 1,
		Amount:     20.0,
	}

	ctx := context.Background()
	result, err := controller.Execute(ctx, updateCmd)

	assert.NoError(t, err)
	assert.Equal(t, "prod-1", result.ProductId)
	assert.Equal(t, "New Name", result.Name)
	assert.Equal(t, 20.0, result.Amount)
}

func TestUpdateProductController_Execute_NotFound(t *testing.T) {

	// Mock repositories
	productRepo := &fixtures.MockProductRepo{

		UpdateByIdFunc: func(ctx context.Context, p *model.Product) error {

			return errors.New("not found")
		},

		FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {

			return nil, errors.New("not found")
		},
	}
	categoryRepo := &fixtures.MockCategoryRepo{

		FindByIdFunc: func(id int) *entity.Category {
			return nil
		},
	}

	container := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}
	controller := controller.NewUpdateProductController(container)

	updateCmd := dto.UpdateProduct{
		ProductId: "prod-404",
		Name:      "Doesn't Matter",
		Amount:    99.0,
	}
	ctx := context.Background()
	_, err := controller.Execute(ctx, updateCmd)

	assert.Error(t, err)
}
