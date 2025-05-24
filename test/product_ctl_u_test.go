// go
package test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/model"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

// Mock ProductRepository for update
type mockProductRepository struct {
	product *model.Product
	err     error
}

func (m *mockProductRepository) UpdateByID(ctx context.Context, id string, update *model.Product) (*model.Product, error) {
	if m.err != nil {
		return nil, m.err
	}
	if m.product == nil || m.product.ID != id {
		return nil, errors.New("not found")
	}
	// Simulate update
	m.product.Name = update.Name
	m.product.Price = update.Price
	return m.product, nil
}

// Mock CategoryRepository
type mockCategoryRepository struct{}

func (m *mockCategoryRepository) FindByID(ctx context.Context, id string) (*model.Category, error) {
	return &model.Category{ID: id, Name: "Test Category"}, nil
}

func TestUpdateProductController_Execute_Success(t *testing.T) {
	mockProd := &model.Product{
		ID:         "prod-1",
		Name:       "Old Name",
		CategoryID: "cat-1",
		Price:      10.0,
	}
	container := &container.Container{
		ProductRepository:  &mockProductRepository{product: mockProd},
		CategoryRepository: &mockCategoryRepository{},
	}
	controller := NewUpdateProductController(container)

	updateCmd := dto.UpdateProduct{
		ID:    "prod-1",
		Name:  "New Name",
		Price: 20.0,
	}
	ctx := context.Background()
	result, err := controller.Execute(ctx, updateCmd)

	assert.NoError(t, err)
	assert.Equal(t, "prod-1", result.ID)
	assert.Equal(t, "New Name", result.Name)
	assert.Equal(t, 20.0, result.Price)
}

func TestUpdateProductController_Execute_NotFound(t *testing.T) {
	container := &container.Container{
		ProductRepository:  &mockProductRepository{product: nil},
		CategoryRepository: &mockCategoryRepository{},
	}
	controller := NewUpdateProductController(container)

	updateCmd := dto.UpdateProduct{
		ID:    "prod-404",
		Name:  "Doesn't Matter",
		Price: 99.0,
	}
	ctx := context.Background()
	_, err := controller.Execute(ctx, updateCmd)

	assert.Error(t, err)
}

func TestUpdateProductController_Execute_Error(t *testing.T) {
	container := &container.Container{
		ProductRepository:  &mockProductRepository{err: errors.New("db error")},
		CategoryRepository: &mockCategoryRepository{},
	}
	controller := NewUpdateProductController(container)

	updateCmd := dto.UpdateProduct{
		ID:    "prod-1",
		Name:  "Any",
		Price: 1.0,
	}
	ctx := context.Background()
	_, err := controller.Execute(ctx, updateCmd)

	assert.Error(t, err)
}
