// go
package test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/model"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

// Mock ProductRepository with FindByCategory
type mockProductRepository struct {
	products []*model.Product
	err      error
}

func (m *mockProductRepository) FindByCategory(ctx context.Context, categoryID int) ([]*model.Product, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.products, nil
}

// Mock CategoryRepository
type mockCategoryRepository struct{}

func (m *mockCategoryRepository) FindByID(ctx context.Context, id int) (*model.Category, error) {
	return &model.Category{ID: id, Name: "Test Category"}, nil
}

func TestFindProductController_Execute_Success(t *testing.T) {
	mockProducts := []*model.Product{
		{ID: 1, Name: "Product 1", CategoryID: 10},
		{ID: 2, Name: "Product 2", CategoryID: 10},
	}
	c := &container.Container{
		ProductRepository:  &mockProductRepository{products: mockProducts},
		CategoryRepository: &mockCategoryRepository{},
	}
	controller := NewFindProductController(c)

	ctx := context.Background()
	result, err := controller.Execute(ctx, 10)

	assert.NoError(t, err)
	assert.Len(t, result.Products, 2)
	assert.Equal(t, "Product 1", result.Products[0].Name)
	assert.Equal(t, "Product 2", result.Products[1].Name)
}

func TestFindProductController_Execute_Empty(t *testing.T) {
	c := &container.Container{
		ProductRepository:  &mockProductRepository{products: []*model.Product{}},
		CategoryRepository: &mockCategoryRepository{},
	}
	controller := NewFindProductController(c)

	ctx := context.Background()
	result, err := controller.Execute(ctx, 99)

	assert.NoError(t, err)
	assert.Len(t, result.Products, 0)
}

func TestFindProductController_Execute_Error(t *testing.T) {
	c := &container.Container{
		ProductRepository:  &mockProductRepository{err: errors.New("db error")},
		CategoryRepository: &mockCategoryRepository{},
	}
	controller := NewFindProductController(c)

	ctx := context.Background()
	_, err := controller.Execute(ctx, 10)

	assert.Error(t, err)
}
