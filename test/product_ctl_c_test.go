package test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

// Mock implementations
type mockProductRepo struct {
	CreateFunc func(ctx context.Context, p entity.Product) (entity.Product, error)
}

func (m *mockProductRepo) Create(ctx context.Context, p entity.Product) (entity.Product, error) {
	return m.CreateFunc(ctx, p)
}

type mockCategoryRepo struct {
	FindByIDFunc func(ctx context.Context, id string) (entity.Category, error)
}

func (m *mockCategoryRepo) FindByID(ctx context.Context, id string) (entity.Category, error) {
	return m.FindByIDFunc(ctx, id)
}

func TestCreateProductController_Execute_Success(t *testing.T) {
	ctx := context.Background()

	mockCategory := entity.Category{ID: "cat1", Name: "Category 1"}
	mockProduct := entity.Product{ID: "prod1", Name: "Product 1", CategoryID: "cat1"}

	productRepo := &mockProductRepo{
		CreateFunc: func(ctx context.Context, p entity.Product) (entity.Product, error) {
			return mockProduct, nil
		},
	}
	categoryRepo := &mockCategoryRepo{
		FindByIDFunc: func(ctx context.Context, id string) (entity.Category, error) {
			if id == "cat1" {
				return mockCategory, nil
			}
			return entity.Category{}, errors.New("not found")
		},
	}

	testContainer := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}

	controller := NewCreateProductController(testContainer)

	input := dto.CreateProduct{
		Name:       "Product 1",
		CategoryID: "cat1",
	}

	result, err := controller.Execute(ctx, input)
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", result.Name)
	assert.Equal(t, "cat1", result.CategoryID)
}

func TestCreateProductController_Execute_CategoryNotFound(t *testing.T) {
	ctx := context.Background()

	productRepo := &mockProductRepo{
		CreateFunc: func(ctx context.Context, p entity.Product) (entity.Product, error) {
			return entity.Product{}, nil
		},
	}
	categoryRepo := &mockCategoryRepo{
		FindByIDFunc: func(ctx context.Context, id string) (entity.Category, error) {
			return entity.Category{}, errors.New("not found")
		},
	}

	testContainer := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}

	controller := NewCreateProductController(testContainer)

	input := dto.CreateProduct{
		Name:       "Product 2",
		CategoryID: "invalid-cat",
	}

	_, err := controller.Execute(ctx, input)
	assert.Error(t, err)
}
