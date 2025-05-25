package controller

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

// Mock implementations
type mockProductRepo struct {
	CreateFunc         func(ctx context.Context, p *model.Product) error
	DeleteByIdFunc     func(ctx context.Context, id string) (*model.Product, error)
	FindByCategoryFunc func(ctx context.Context, categoryId int) (*[]model.Product, error)
	FindOneFunc        func(ctx context.Context, id string) (*model.Product, error)
	UpdateByIdFunc     func(ctx context.Context, p *model.Product) error
}

func (m *mockProductRepo) Create(ctx context.Context, p *model.Product) error {
	return m.CreateFunc(ctx, p)
}

func (m *mockProductRepo) DeleteById(ctx context.Context, id string) (*model.Product, error) {
	if m.DeleteByIdFunc != nil {
		return m.DeleteByIdFunc(ctx, id)
	}
	return nil, nil
}

func (m *mockProductRepo) FindByCategory(ctx context.Context, categoryId int) (*[]model.Product, error) {
	if m.FindByCategoryFunc != nil {
		return m.FindByCategoryFunc(ctx, categoryId)
	}
	return nil, nil
}

func (m *mockProductRepo) FindOne(ctx context.Context, id string) (*model.Product, error) {
	if m.FindOneFunc != nil {
		return m.FindOneFunc(ctx, id)
	}
	return nil, nil
}

func (m *mockProductRepo) UpdateById(ctx context.Context, p *model.Product) error {
	if m.UpdateByIdFunc != nil {
		return m.UpdateByIdFunc(ctx, p)
	}
	return nil
}

type mockCategoryRepo struct {
	FindByIdFunc func(id int) *entity.Category
}

func (m *mockCategoryRepo) FindById(id int) *entity.Category {
	return m.FindByIdFunc(id)
}

func TestCreateProductController_Execute_Success(t *testing.T) {
	ctx := context.Background()

	// Mock repositories
	productRepo := &mockProductRepo{
		CreateFunc: func(ctx context.Context, p *model.Product) error {
			return nil
		},
	}
	categoryRepo := &mockCategoryRepo{
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
	productRepo := &mockProductRepo{
		CreateFunc: func(ctx context.Context, p *model.Product) error {
			return nil
		},
	}
	categoryRepo := &mockCategoryRepo{
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

func TestDeleteProductController_Execute_Success(t *testing.T) {
	ctx := context.Background()

	productRepo := &mockProductRepo{
		DeleteByIdFunc: func(ctx context.Context, id string) (*model.Product, error) {
			if id == "prod1" {
				return &model.Product{ID: "prod1"}, nil
			}
			return nil, errors.New("not found")
		},
	}

	testContainer := &container.Container{
		ProductRepository: productRepo,
	}

	controller := NewDeleteProductController(testContainer)

	id := "prod1"
	result, err := controller.Execute(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, id, result)
}

func TestDeleteProductController_Execute_NotFound(t *testing.T) {
	ctx := context.Background()

	productRepo := &mockProductRepo{
		DeleteByIdFunc: func(ctx context.Context, id string) (*model.Product, error) {
			return nil, errors.New("not found")
		},
	}

	testContainer := &container.Container{
		ProductRepository: productRepo,
	}

	controller := NewDeleteProductController(testContainer)

	id := ""
	result, err := controller.Execute(ctx, id)
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestFindOneProductController_Execute_Success_WithMock(t *testing.T) {

	mockProduct := &model.Product{
		ID:          "prod1",
		Name:        "Product 1",
		Description: "Description 1",
		Amount:      100.0,
		CategoryId:  1,
		CreatedAt:   time.Now(),
	}

	productRepo := &mockProductRepo{
		FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {
			if id == "prod1" {
				return mockProduct, nil
			}
			return nil, errors.New("not found")
		},
	}

	categoryRepo := &mockCategoryRepo{
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

	// Mock repositories
	productRepo := &mockProductRepo{
		FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {

			return nil, errors.New("not found")
		},
	}

	categoryRepo := &mockCategoryRepo{
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

func TestFindProductController_Execute_Success(t *testing.T) {

	mockProductsSlice := []model.Product{
		{ID: "prod1", Name: "Product 1", CategoryId: 10},
		{ID: "prod2", Name: "Product 2", CategoryId: 10},
	}
	mockProducts := &mockProductsSlice

	productRepo := &mockProductRepo{
		FindByCategoryFunc: func(ctx context.Context, categoryId int) (*[]model.Product, error) {
			if categoryId == 10 {

				return mockProducts, nil
			}

			return nil, errors.New("not found")
		},
	}

	categoryRepo := &mockCategoryRepo{
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

	// Mock repositories
	productRepo := &mockProductRepo{
		CreateFunc: func(ctx context.Context, p *model.Product) error {
			return nil
		},
	}
	categoryRepo := &mockCategoryRepo{
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

func TestUpdateProductController_Execute_Success(t *testing.T) {

	mockProduct := &model.Product{
		ID:         "prod-1",
		Name:       "Old Name",
		CategoryId: 1,
		Amount:     10.0,
	}

	// Mock repositories
	productRepo := &mockProductRepo{

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

	categoryRepo := &mockCategoryRepo{
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

	controller := NewUpdateProductController(container)

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
	productRepo := &mockProductRepo{

		UpdateByIdFunc: func(ctx context.Context, p *model.Product) error {

			return errors.New("not found")
		},

		FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {

			return nil, errors.New("not found")
		},
	}
	categoryRepo := &mockCategoryRepo{

		FindByIdFunc: func(id int) *entity.Category {
			return nil
		},
	}

	container := &container.Container{
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}
	controller := NewUpdateProductController(container)

	updateCmd := dto.UpdateProduct{
		ProductId: "prod-404",
		Name:      "Doesn't Matter",
		Amount:    99.0,
	}
	ctx := context.Background()
	_, err := controller.Execute(ctx, updateCmd)

	assert.Error(t, err)
}
