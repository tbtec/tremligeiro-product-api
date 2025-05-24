package test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/model"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

// Mock ProductRepository for integration-like testing
type mockProductRepository struct {
	product *model.Product
	err     error
}

func (m *mockProductRepository) FindByID(ctx context.Context, id string) (*model.Product, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.product, nil
}

// Mock CategoryRepository
type mockCategoryRepository struct{}

func (m *mockCategoryRepository) FindByID(ctx context.Context, id string) (*model.Category, error) {
	return &model.Category{ID: id, Name: "Test Category"}, nil
}

func TestFindOneProductController_Execute_Success(t *testing.T) {
	mockProduct := &model.Product{
		ID:         "prod-123",
		Name:       "Test Product",
		CategoryID: "cat-1",
	}
	container := &container.Container{
		ProductRepository:  &mockProductRepository{product: mockProduct},
		CategoryRepository: &mockCategoryRepository{},
	}
	controller := NewFindOneProductController(container)

	ctx := context.Background()
	result, err := controller.Execute(ctx, "prod-123")

	assert.NoError(t, err)
	assert.Equal(t, "prod-123", result.ID)
	assert.Equal(t, "Test Product", result.Name)
}

func TestFindOneProductController_Execute_NotFound(t *testing.T) {
	container := &container.Container{
		ProductRepository:  &mockProductRepository{err: errors.New("not found")},
		CategoryRepository: &mockCategoryRepository{},
	}
	controller := NewFindOneProductController(container)

	ctx := context.Background()
	_, err := controller.Execute(ctx, "prod-404")

	assert.Error(t, err)
}
